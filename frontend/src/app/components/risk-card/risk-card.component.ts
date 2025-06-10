import { Component, Input } from '@angular/core';
import {NgIf} from '@angular/common';
import {Register} from "../../models/register.model";

@Component({
  selector: 'app-risk-card',
  templateUrl: './risk-card.component.html',
  imports: [],
  styleUrls: ['./risk-card.component.scss']
})
export class RiskCardComponent {
  @Input() register!: Register;
}
