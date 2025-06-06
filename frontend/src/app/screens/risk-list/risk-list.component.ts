import { Component, OnInit } from '@angular/core';
import { Risk } from '../../models/risk.model';
import { RiskService } from '../../services/risk.service';
import {NgForOf, NgIf} from '@angular/common';
import {SearchBarComponent} from '../../components/search-bar/search-bar.component';
import {RiskCardComponent} from '../../components/risk-card/risk-card.component';

@Component({
  selector: 'app-risk-list',
  templateUrl: './risk-list.component.html',
  imports: [
    NgForOf,
    NgIf,
    SearchBarComponent,
    RiskCardComponent
  ],
  styleUrls: ['./risk-list.component.scss']
})
export class RiskListComponent implements OnInit {
  risks: Risk[] = [];
  filteredRisks: Risk[] = [];
  isLoading: boolean = true;
  errorMessage: string = '';

  constructor(private riskService: RiskService) { }

  ngOnInit(): void {
    this.loadRisks();
  }

  loadRisks(): void {
    this.isLoading = false;
    setTimeout(() => {
        this.riskService.getRisks().subscribe({
          next: (data) => {
            this.risks = data;
            this.filteredRisks = data;
            this.isLoading = false;
          },
          error: (error) => {
            this.errorMessage = 'Erro ao carregar riscos. Por favor, tente novamente.';
            this.isLoading = false;
          }
        });
      }, 5000);

  }

  filterRisks(searchTerm: string): void {
    if (!searchTerm) {
      this.filteredRisks = this.risks;
      return;
    }

    searchTerm = searchTerm.toLowerCase();
    this.filteredRisks = this.risks.filter(risk =>
      risk.title.toLowerCase().includes(searchTerm) ||
      risk.description.toLowerCase().includes(searchTerm) ||
      (risk.vehicle && risk.vehicle.toLowerCase().includes(searchTerm)) ||
      (risk.area && risk.area.toLowerCase().includes(searchTerm)) ||
      (risk.person && risk.person.toLowerCase().includes(searchTerm))
    );
  }
}
