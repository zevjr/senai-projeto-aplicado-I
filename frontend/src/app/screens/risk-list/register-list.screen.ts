import {Component, OnInit, signal, WritableSignal} from '@angular/core';
import { RegisterService } from '../../services/register.service';
import {NgIf} from '@angular/common';
import {RiskCardComponent} from '../../components/risk-card/risk-card.component';
import {Register} from "../../models/register.model";
import {Router} from '@angular/router';

@Component({
  selector: 'app-risk-list',
  templateUrl: './register-list.screen.html',
  imports: [
    NgIf,
    RiskCardComponent
  ],
  styleUrls: ['./register-list.screen.scss']
})
export class RegisterListScreen implements OnInit {
  registers: Register[] = [];
  filteredRisks: WritableSignal<Register[]> = signal([]);
  isLoading: WritableSignal<Boolean> = signal(false);
  errorMessage: string = '';

  constructor(
    private readonly registersService: RegisterService,
    private readonly router: Router
  ) { }

  ngOnInit(): void {
    this.loadRisks();
  }

  loadRisks(): void {
    this.isLoading.set(false);
    this.registersService.getRegisters().subscribe({
      next: (data) => {
        this.registers = data;
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

    onDeleteRisk(uid: string) {
        this.registersService.deleteRegister(uid).subscribe({
            next: () => {
                this.loadRisks();
            },
            error: (error) => {
                console.error(error);
                this.errorMessage = 'Erro ao deletar risco. Por favor, tente novamente.';
            }
        });
    }

    onEditRisk(uid: string) {
      this.router.navigate(['/risks/edit', uid]);
    }
}
