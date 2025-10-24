import { Component, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
    selector: 'app-modal',
    standalone: true,
    imports: [CommonModule],
    templateUrl: './modal.component.html',
    styleUrls: ['./modal.component.scss']
})
export class ModalComponent {
    isVisible: boolean = false;
    title: string = '';
    message: string = '';
    confirmText: string = 'Confirmar';
    cancelText: string = 'Cancelar';

    private resolvePromise?: (value: boolean) => void;

    constructor(private cdr: ChangeDetectorRef) {}

    open(options: {
        title?: string;
        message: string;
        confirmText?: string;
        cancelText?: string;
    }): Promise<boolean> {
        this.title = options.title || 'Confirmação';
        this.message = options.message;
        this.confirmText = options.confirmText || 'Confirmar';
        this.cancelText = options.cancelText || 'Cancelar';
        this.isVisible = true;

        this.cdr.detectChanges();

        return new Promise<boolean>((resolve) => {
            this.resolvePromise = resolve;
        });
    }

    confirm(): void {
        this.isVisible = false;
        if (this.resolvePromise) {
            this.resolvePromise(true);
        }
    }

    cancel(): void {
        this.isVisible = false;
        if (this.resolvePromise) {
            this.resolvePromise(false);
        }
    }

    close(): void {
        this.cancel();
    }
}
