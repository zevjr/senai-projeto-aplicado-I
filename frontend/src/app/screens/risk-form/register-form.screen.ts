import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import {HttpClient} from '@angular/common/http';
import {RegisterService} from "../../services/register.service";
import {FileService} from "../../services/file.service";
import {ToastrService} from "ngx-toastr";
import {ActivatedRoute, Router} from '@angular/router';

@Component({
    selector: 'app-risk-form',
    templateUrl: './register-form.screen.html',
    imports: [
        ReactiveFormsModule
    ],
    styleUrls: ['./register-form.screen.scss']
})
export class RegisterFormScreen implements OnInit {
    // Form
    riskForm: FormGroup;

    // File Uploads
    selectedAudio: File | null = null;
    selectedPhoto: File | null = null;

    // Estado de UI
    isLoading: boolean = false;
    errorMessage: string | null = null;

    // Modo de edição
    editMode: boolean = false;
    registerUid: string | null = null;

    constructor(
        private fb: FormBuilder,
        private http: HttpClient,
        private registerService: RegisterService,
        private fileService: FileService,
        private toastr: ToastrService,
        private route: ActivatedRoute,
        private router: Router
    ) {
        this.riskForm = this.fb.group({
            location: ['', Validators.required],
            riskScale: [5, [Validators.required, Validators.min(1), Validators.max(10)]],
            description: ['', Validators.required]
        });
    }

    ngOnInit(): void {
        // Verificar se estamos editando um registro existente
        this.route.paramMap.subscribe(params => {
            const uid = params.get('uid');

            if (uid) {
                this.registerUid = uid;
                this.editMode = true;
                this.loadRegister(uid);
            }
        });
    }

    // Carrega um registro existente para edição
    loadRegister(uid: string): void {
        this.isLoading = true;

        this.registerService.getRegister(uid as any).subscribe({
            next: (register) => {
                // Definir valores iniciais
                const initialValues = {
                    location: register.local,
                    riskScale: register.riskScale,
                    description: register.body
                };

                // Atualizar formulário
                this.riskForm.patchValue(initialValues);
                this.isLoading = false;
            },
            error: (error) => {
                console.error('Erro ao carregar registro:', error);
                this.errorMessage = 'Erro ao carregar o registro. Tente novamente.';
                this.isLoading = false;
            }
        });
    }

    // Manipulação de arquivos de áudio
    onAudioSelected(event: any): void {
        const file = event.target.files[0];
        if (file) {
            this.selectedAudio = file;
        }
    }

    removeAudio(): void {
        this.selectedAudio = null;
    }

    // Manipulação de arquivos de imagem
    onPhotoSelected(event: any): void {
        const file = event.target.files[0];
        if (file) {
            this.selectedPhoto = file;
        }
    }

    removePhoto(): void {
        this.selectedPhoto = null;
    }

    // Verifica se o formulário é válido
    isFormValid(): boolean {
        return this.riskForm.valid;
    }

    // Submissão do formulário
    onSubmit(): void {
        if (!this.riskForm.valid) {
            return;
        }

        this.isLoading = true;
        this.errorMessage = null;

        const uploadTasks = [];

        // Upload de foto (se houver)
        if (this.selectedPhoto) {
            uploadTasks.push(
                this.fileService.uploadImage(this.selectedPhoto).toPromise()
            );
        } else {
            uploadTasks.push(Promise.resolve({uid: null}));
        }

        // Upload de áudio (se houver)
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

                // Escolher operação baseado no modo
                const operation = this.editMode && this.registerUid
                    ? this.registerService.updateRegister(this.registerUid, register)
                    : this.registerService.createRegister(register);

                operation.subscribe({
                    next: () => {
                        const message = this.editMode
                            ? 'Registro atualizado com sucesso'
                            : 'Registro criado com sucesso';

                        this.toastr.success(message);
                        this.isLoading = false;

                        // Redirecionar para a lista
                        this.router.navigate(['/risks']);
                    },
                    error: (error) => {
                        const action = this.editMode ? 'atualizar' : 'criar';
                        this.errorMessage = `Erro ao ${action} o registro. Tente novamente.`;
                        this.toastr.error(this.errorMessage);
                        this.isLoading = false;
                    }
                });
            })
            .catch(error => {
                this.errorMessage = 'Erro ao enviar os arquivos. Tente novamente.';
                this.toastr.error(this.errorMessage);
                this.isLoading = false;
            });
    }
}