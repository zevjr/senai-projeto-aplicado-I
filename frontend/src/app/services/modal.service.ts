import { Injectable, ComponentRef, ApplicationRef, createComponent, EnvironmentInjector } from '@angular/core';
import { ModalComponent } from '../components/modal/modal.component';

@Injectable({
    providedIn: 'root'
})
export class ModalService {
    private modalComponentRef?: ComponentRef<ModalComponent>;

    constructor(
        private readonly appRef: ApplicationRef,
        private readonly injector: EnvironmentInjector
    ) {}

    open(options: {
        title?: string;
        message: string;
        confirmText?: string;
        cancelText?: string;
    }): Promise<boolean> {
        console.log('ModalService.open called with options:', options);

        if (this.modalComponentRef) {
            this.close();
        }

        this.modalComponentRef = createComponent(ModalComponent, {
            environmentInjector: this.injector
        });

        this.appRef.attachView(this.modalComponentRef.hostView);

        const domElem = (this.modalComponentRef.hostView as any).rootNodes[0] as HTMLElement;
        document.body.appendChild(domElem);

        const promise = this.modalComponentRef.instance.open(options);

        promise.then((result) => {
            console.log('Modal closed with result:', result);
            setTimeout(() => {
                this.close();
            }, 100);
            return result;
        }).catch((error) => {
            console.error('Modal error:', error);
            this.close();
        });

        return promise;
    }

    private close(): void {
        console.log('Closing modal');
        if (this.modalComponentRef) {
            this.appRef.detachView(this.modalComponentRef.hostView);
            this.modalComponentRef.destroy();
            this.modalComponentRef = undefined;
        }
    }

    confirm(message: string, title?: string): Promise<boolean> {
        return this.open({
            title: title || 'Confirmação',
            message,
            confirmText: 'Confirmar',
            cancelText: 'Cancelar'
        });
    }

    alert(message: string, title?: string): Promise<boolean> {
        return this.open({
            title: title || 'Aviso',
            message,
            confirmText: 'OK',
            cancelText: 'Fechar'
        });
    }

    delete(message?: string): Promise<boolean> {
        return this.open({
            title: 'Confirmar Exclusão',
            message: message || 'Tem certeza que deseja excluir este item? Esta ação não pode ser desfeita.',
            confirmText: 'Excluir',
            cancelText: 'Cancelar'
        });
    }
}
