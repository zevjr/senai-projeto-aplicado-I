import {ChangeDetectorRef, Component, signal} from '@angular/core';
import {FormBuilder, FormGroup, ReactiveFormsModule, Validators} from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import {CommonModule, NgClass, NgIf} from '@angular/common';

@Component({
  selector: 'app-login',
   standalone: true,
  //templateUrl: './login.screen.html',
  imports: [
     CommonModule,
    ReactiveFormsModule,
    RouterModule,
    NgIf
     
  ],
   templateUrl: './login.screen.html',
  styleUrls: ['./login.screen.scss']
})
export class LoginScreen {
  loginForm: FormGroup;
  errorMessage: string = '';
  isLoading = signal(false);

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    private router: Router
  ) {
    this.loginForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', Validators.required]
    });
  }

  onSubmit(): void {
    if (this.loginForm.invalid) {
      return;
    }


    this.errorMessage = '';

    const { email, password } = this.loginForm.value;
    this.isLoading.set(true);
    this.authService.login(email, password).subscribe({
      next: () => {
        this.router.navigate(['/risks']);
      },
      error: (error) => {
        console.log("chegou aqui")
        this.errorMessage = error.message;
        this.isLoading.set(false);
      }
    });
  }
}
