import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import { Router } from '@angular/router';
import { RiskService } from '../../services/risk.service';
import {NgForOf, NgIf} from '@angular/common';

@Component({
  selector: 'app-risk-form',
  templateUrl: './risk-form.component.html',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    NgIf,
    NgForOf
  ],
  styleUrls: ['./risk-form.component.scss']
})
export class RiskFormComponent implements OnInit {
  riskForm: FormGroup;
  isLoading: boolean = false;
  errorMessage: string = '';
  selectedFiles: File[] = [];

  constructor(
    private fb: FormBuilder,
    private riskService: RiskService,
    private router: Router
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
        this.selectedFiles.push(files[i]);
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

    this.riskService.createRisk(riskData).subscribe({
      next: () => {
        this.router.navigate(['/risks']);
      },
      error: (error) => {
        this.errorMessage = 'Erro ao criar risco. Por favor, tente novamente.';
        this.isLoading = false;
      },
      complete: () => {
        this.isLoading = false;
      }
    });
  }
}
