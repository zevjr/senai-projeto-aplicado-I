import { Component, OnInit, signal, WritableSignal } from '@angular/core';
import { DatePipe, NgClass, NgFor, NgIf } from '@angular/common';
import { RegisterService } from '../../services/register.service';
import { Register } from '../../models/register.model';

interface DashboardStatusConfig {
  key: string;
  label: string;
  icon: string;
  badgeClass: string;
  barClass: string;
  accentClass: string;
}

interface DashboardStatusSummary extends DashboardStatusConfig {
  count: number;
  percent: number;
}

interface DashboardGroup {
  key: string;
  label: string;
  registers: Register[];
}

@Component({
  selector: 'app-admin-dashboard',
  standalone: true,
  templateUrl: './admin-dashboard.screen.html',
  styleUrls: ['./admin-dashboard.screen.scss'],
  imports: [
    NgIf,
    NgFor,
    NgClass,
    DatePipe
  ]
})
export class AdminDashboardScreen implements OnInit {
  registers: Register[] = [];
  groupedRegisters: DashboardGroup[] = [];
  statusSummary: DashboardStatusSummary[] = [];
  totalRegisters = 0;

  isLoading: WritableSignal<boolean> = signal(false);
  errorMessage: WritableSignal<string> = signal('');

  private readonly statusConfig: DashboardStatusConfig[] = [
    {
      key: 'aberto',
      label: 'Aberto',
      icon: 'bi-exclamation-octagon',
      badgeClass: 'bg-danger-subtle text-danger-emphasis',
      barClass: 'bg-danger',
      accentClass: 'text-danger'
    },
    {
      key: 'em analise',
      label: 'Em analise',
      icon: 'bi-search',
      badgeClass: 'bg-warning-subtle text-warning-emphasis',
      barClass: 'bg-warning',
      accentClass: 'text-warning'
    },
    {
      key: 'fechado',
      label: 'Fechado',
      icon: 'bi-check-circle',
      badgeClass: 'bg-success-subtle text-success-emphasis',
      barClass: 'bg-success',
      accentClass: 'text-success'
    }
  ];

  constructor(private readonly registerService: RegisterService) {}

  ngOnInit(): void {
    this.loadRegisters();
  }

  loadRegisters(): void {
    this.isLoading.set(true);
    this.errorMessage.set('');

    this.registerService.getRegisters().subscribe({
      next: (data) => {
        this.registers = data;
        this.totalRegisters = data.length;
        this.groupedRegisters = this.buildGroups(data);
        this.statusSummary = this.buildSummary(data);
        this.isLoading.set(false);
      },
      error: () => {
        this.errorMessage.set('Erro ao carregar registros. Tente novamente em instantes.');
        this.isLoading.set(false);
      }
    });
  }

  statusBadgeClass(statusKey: string): string {
    return this.statusConfig.find((cfg) => cfg.key === statusKey)?.badgeClass ?? 'bg-secondary';
  }

  statusLabel(statusKey: string): string {
    return this.statusConfig.find((cfg) => cfg.key === statusKey)?.label ?? 'Outros';
  }

  statusBadgeClassForValue(status?: string): string {
    return this.statusBadgeClass(this.normalizeStatus(status));
  }

  statusLabelForValue(status?: string): string {
    return this.statusLabel(this.normalizeStatus(status));
  }

  resolveImageUid(register: Register): string | undefined {
    return register.imageUid ?? register.image_uid ?? undefined;
  }

  resolveAudioUid(register: Register): string | undefined {
    return register.audioUid ?? register.audio_uid ?? undefined;
  }

  resolveCreatedAt(register: Register): string | undefined {
    return register.createdAt ?? register.created_at ?? undefined;
  }

  trackByStatus = (_index: number, item: DashboardGroup): string => item.key;

  private buildGroups(registers: Register[]): DashboardGroup[] {
    const map = new Map<string, Register[]>();

    registers.forEach((register) => {
      const normalized = this.normalizeStatus(register.status);
      const statusKey = this.statusConfig.find((cfg) => cfg.key === normalized) ? normalized : 'outros';

      if (!map.has(statusKey)) {
        map.set(statusKey, []);
      }

      map.get(statusKey)?.push(register);
    });

    const result: DashboardGroup[] = this.statusConfig.map((cfg) => ({
      key: cfg.key,
      label: cfg.label,
      registers: map.get(cfg.key) ?? []
    }));

    if (map.has('outros')) {
      result.push({
        key: 'outros',
        label: 'Outros',
        registers: map.get('outros') ?? []
      });
    }

    return result;
  }

  private buildSummary(registers: Register[]): DashboardStatusSummary[] {
    const total = registers.length || 1;

    return this.statusConfig.map((cfg) => {
      const count = registers.filter(
        (register) => this.normalizeStatus(register.status) === cfg.key
      ).length;

      return {
        ...cfg,
        count,
        percent: total ? Math.round((count / total) * 100) : 0
      };
    });
  }

  private normalizeStatus(value?: string): string {
    if (!value) {
      return 'desconhecido';
    }

    return value
      .normalize('NFD')
      .replace(/[\u0300-\u036f]/g, '')
      .toLowerCase()
      .trim();
  }
}
