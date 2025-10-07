import { Component, inject } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { UserService } from '../../../services/user-service';
import { IAuthResponse, ISign } from '../../../models/types';
import { HttpErrorResponse } from '@angular/common/http';
import { Router } from '@angular/router';
import { AuthService } from '../../../services/auth-service';
import { BootstrapService } from '../../../services/bootstrap-service';

@Component({
  selector: 'app-login',
  imports: [ReactiveFormsModule],
  templateUrl: './login.html',
  styleUrl: './login.css'
})
export class Login {

  userService = inject(UserService);
  authService = inject(AuthService);
  bootstrapService = inject(BootstrapService)
  router = inject(Router);

  loginForm = new FormGroup({
    email: new FormControl('', {
      validators: [Validators.email, Validators.required],
    }),
    password: new FormControl('', {
      validators: [Validators.required]
    }),
  });

  onLoginUser() {
    if (this.loginForm.valid) {
      const user: ISign = {
        email: this.loginForm.value.email!,
        password: this.loginForm.value.password!
      }
      this.userService.loginUser(user).subscribe({
        next: (response: IAuthResponse) => {
          this.authService.setToken(response.token);
          this.bootstrapService.loadAllData();
          this.router.navigate(['/home']);
        },
        error: (err: HttpErrorResponse) => console.error(err)
      });
    } else {
      alert("Credentials invalid")
    }
  }
}

