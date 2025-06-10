import {Component, signal, WritableSignal} from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {HttpClient} from '@angular/common/http';
import {RegisterService} from "../../services/register.service";
import {FileService} from "../../services/file.service";
import {ToastrService} from "ngx-toastr";

@Component({
    selector: 'app-risk-form',
    templateUrl: './register-form.screen.html',
    imports: [
        ReactiveFormsModule
    ],
    styleUrls: ['./register-form.screen.scss']
})
export class RegisterFormScreen {
    riskForm: FormGroup;
    selectedAudio: File | null = null;
    selectedPhoto: File | null = null;
    isLoading: WritableSignal<boolean> = signal(false);
    errorMessage: string | null = null;

    constructor(private fb: FormBuilder,
                private http: HttpClient,
                private registerService: RegisterService,
                private fileService: FileService,
                private toastr: ToastrService) {
        this.riskForm = this.fb.group({
            location: ['', Validators.required],
            riskLevel: ['', Validators.required],
            riskScale: [5, [Validators.required, Validators.min(1), Validators.max(10)]],
            description: ['', Validators.required]
        });
    }

    onAudioSelected(event: any): void {
        const file = event.target.files[0];
        if (file) {
            this.selectedAudio = file;
        }
    }

    removeAudio(): void {
        this.selectedAudio = null;
    }

    onPhotoSelected(event: any): void {
        const file = event.target.files[0];
        if (file) {
            this.selectedPhoto = file;
        }
    }

    removePhoto(): void {
        this.selectedPhoto = null;
    }

    onSubmit(): void {
        this.isLoading.set(true);
        this.errorMessage = null;

        const uploadTasks = [];

        // Upload photo if provided
        if (this.selectedPhoto) {
            uploadTasks.push(
                this.fileService.uploadImage(this.selectedPhoto).toPromise()
            );
        } else {
            uploadTasks.push(Promise.resolve({uid: null}));
        }

        // Upload audio if provided
        if (this.selectedAudio) {
            uploadTasks.push(
                this.fileService.uploadAudio(this.selectedAudio).toPromise()
            );
        } else {
            uploadTasks.push(Promise.resolve({uid: null}));
        }

        Promise.all(uploadTasks)
            .then(([imageResponse, audioResponse]) => {
                const register = {
                    title: this.riskForm.get('description')?.value,
                    body: this.riskForm.get('description')?.value,
                    riskScale: this.riskForm.get('riskScale')?.value,
                    local: this.riskForm.get('location')?.value,
                    status: 'Pending',
                    imageUid: imageResponse?.uid || undefined,
                    audioUid: audioResponse?.uid || undefined
                };

                this.registerService.createRegister(register).subscribe(
                    () => {
                        console.log('Registro criado com sucesso');
                        this.isLoading.set(false);
                        this.riskForm.reset();
                        this.selectedAudio = null;
                        this.selectedPhoto = null;
                    },
                    error => {
                        console.error('Erro ao criar o registro:', error);
                        this.errorMessage = 'Erro ao criar o registro. Tente novamente.';
                        this.isLoading.set(false);
                    }
                );
            })
            .catch(error => {
                console.error('Erro ao enviar os arquivos:', error);
                this.errorMessage = 'Erro ao enviar os arquivos. Tente novamente.';
                this.isLoading.set(false);
            });
    }
}