import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import { Router } from '@angular/router';
import { RegistersService } from '../../services/registers.service';
import {NgForOf, NgIf} from '@angular/common';
import {ToastrService} from "ngx-toastr";

@Component({
  selector: 'app-risk-form',
  templateUrl: './risk-form.screen.html',
  imports: [
    ReactiveFormsModule,
    NgIf,
    NgForOf
  ],
  styleUrls: ['./risk-form.screen.scss']
})
export class RiskFormScreen implements OnInit {
  riskForm: FormGroup;
  isLoading: boolean = false;
  errorMessage: string = '';
  selectedFiles: File[] = [];

  constructor(
    private fb: FormBuilder,
    private registersService: RegistersService,
    private router: Router,
    private toastr: ToastrService
  ) {
    this.riskForm = this.fb.group({
      location: ['', Validators.required],
      riskLevel: ['', Validators.required],
      riskScale: [5, [Validators.required, Validators.min(1), Validators.max(10)]],
      description: ['', Validators.required]
    });
  }

  ngOnInit(): void {
  }

  onFileSelected(event: any): void {
    const files = event.target.files;
    if (files) {
      for (let i = 0; i < files.length; i++) {
        const file = files[i];
        if (file.type.startsWith('audio/') && !this.selectedFiles.some(f => f.type.startsWith('audio/'))) {
          this.selectedFiles.push(file);
        } else if (file.type.startsWith('video/') && !this.selectedFiles.some(f => f.type.startsWith('video/'))) {
          this.selectedFiles.push(file);
        } else {
          this.toastr.warning('You can only upload one audio and one video file.', 'Warning');
        }
      }
    }
  }

  removeFile(index: number): void {
    this.selectedFiles.splice(index, 1);
  }

  onSubmit(): void {
    if (this.riskForm.invalid) {
      return;
    }

    this.isLoading = true;
    this.errorMessage = '';

    const riskData = {
      ...this.riskForm.value,
      title: `Risk: ${this.riskForm.value.riskLevel}`
    };

    this.registersService.createRegister(riskData).subscribe({
      next: () => {
        this.toastr.success('Registro de risco criado com sucesso!', 'Sucesso');
        this.router.navigate(['/risks']);
      },
      error: (error) => {
        this.errorMessage = 'Erro ao criar registro de risco. Por favor, tente novamente.';
        this.isLoading = false;
      },
      complete: () => {
        this.isLoading = false;
      }
    });
  }
}
