import {Component, signal, WritableSignal} from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import { HttpClient } from '@angular/common/http';

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

  constructor(private fb: FormBuilder, private http: HttpClient) {
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
    if (!this.selectedAudio || !this.selectedPhoto) {
      this.errorMessage = 'Por favor, envie um áudio e uma foto antes de enviar o formulário.';
      return;
    }

    this.isLoading.set(true);
    this.errorMessage = null;

    const formData = new FormData();
    formData.append('location', this.riskForm.get('location')?.value);
    formData.append('riskLevel', this.riskForm.get('riskLevel')?.value);
    formData.append('riskScale', this.riskForm.get('riskScale')?.value);
    formData.append('description', this.riskForm.get('description')?.value);
    formData.append('audio', this.selectedAudio);
    formData.append('photo', this.selectedPhoto);

    this.http.post('/api/risk-form', formData).subscribe(
        response => {
          console.log('Formulário enviado com sucesso:', response);
          this.isLoading.set(false);
          this.riskForm.reset();
          this.selectedAudio = null;
          this.selectedPhoto = null;
        },
        error => {
          console.error('Erro ao enviar o formulário:', error);
          this.errorMessage = 'Erro ao enviar o formulário. Tente novamente.';
          this.isLoading.set(false);
        }
    );
  }
}