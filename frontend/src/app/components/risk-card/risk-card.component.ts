import { Component, Input, Output, EventEmitter } from '@angular/core';
import { Register } from "../../models/register.model";
import { ModalService } from "../../services/modal.service";

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

  constructor(private readonly modalService: ModalService) {}

  async onDeleteClick(): Promise<void> {

    const confirmed = await this.modalService.open({
      title: 'Confirmar Exclusão',
      message: `Tem certeza que deseja excluir o risco "${this.register.title}"? Esta ação não pode ser desfeita.`,
      confirmText: 'Sim, excluir',
      cancelText: 'Cancelar'
    });

    if (confirmed) {
      this.delete.emit(this.register);
    }
  }

  onEditClick(): void {
    this.edit.emit(this.register);
  }
}
