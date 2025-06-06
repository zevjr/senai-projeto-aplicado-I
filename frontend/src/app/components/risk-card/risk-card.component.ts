import { Component, Input } from '@angular/core';
import { Risk } from '../../models/risk.model';
import {NgIf} from '@angular/common';

@Component({
  selector: 'app-risk-card',
  templateUrl: './risk-card.component.html',
  imports: [
    NgIf
  ],
  styleUrls: ['./risk-card.component.scss']
})
export class RiskCardComponent {
  @Input() risk!: Risk;
}
