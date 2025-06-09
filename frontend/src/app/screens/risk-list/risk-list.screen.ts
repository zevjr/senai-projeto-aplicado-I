import {Component, OnInit, signal, WritableSignal} from '@angular/core';
import { Risk } from '../../models/risk.model';
import { RegistersService } from '../../services/registers.service';
import {NgIf} from '@angular/common';
import {RiskCardComponent} from '../../components/risk-card/risk-card.component';

@Component({
  selector: 'app-risk-list',
  templateUrl: './risk-list.screen.html',
  imports: [
    NgIf,
    RiskCardComponent
  ],
  styleUrls: ['./risk-list.screen.scss']
})
export class RiskListScreen implements OnInit {
  risks: Risk[] = [];
  filteredRisks: WritableSignal<Risk[]> = signal([]);
  isLoading: WritableSignal<Boolean> = signal(false);
  errorMessage: string = '';

  constructor(private riskService: RegistersService) { }

  ngOnInit(): void {
    this.loadRisks();
  }

  loadRisks(): void {
    this.isLoading.set(false);
    this.riskService.getRegisters().subscribe({
      next: (data) => {
        this.risks = data;
        this.filteredRisks.set(data);
        this.isLoading.set(false);
      },
      error: (error) => {
        this.errorMessage = 'Erro ao carregar riscos. Por favor, tente novamente.';
        this.isLoading.set(false);
      }
    });

  }

  // filterRisks(searchTerm: string): void {
  //   if (!searchTerm) {
  //     this.filteredRisks.set(this.risks);
  //     return;
  //   }
  //
  //   searchTerm = searchTerm.toLowerCase();
  //   this.filteredRisks.set(this.risks.filter(risk =>
  //     risk.title.toLowerCase().includes(searchTerm) ||
  //     risk.description.toLowerCase().includes(searchTerm) ||
  //     (risk.vehicle && risk.vehicle.toLowerCase().includes(searchTerm)) ||
  //     (risk.area && risk.area.toLowerCase().includes(searchTerm)) ||
  //     (risk.person && risk.person.toLowerCase().includes(searchTerm))
  //   ));
  // }
}
