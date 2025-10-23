import { Component, Input, Output, EventEmitter } from '@angular/core';
import { Register } from "../../models/register.model";

@Component({
  selector: 'app-risk-card',
  templateUrl: './risk-card.component.html',
  styleUrls: ['./risk-card.component.scss'],
  standalone: true
})
export class RiskCardComponent {
  @Input() register!: Register;
  @Output() delete = new EventEmitter<Register>();
  @Output() edit = new EventEmitter<Register>();

  onDeleteClick(): void {
    this.delete.emit(this.register);
  }

  onEditClick(): void {
    this.edit.emit(this.register);
  }
}
