import { Component, inject } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { UserService } from '../../../services/user-service';
import { ISign } from '../../../models/types';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-login',
  imports: [ReactiveFormsModule],
  templateUrl: './login.html',
  styleUrl: './login.css'
})
export class Login {

  userService = inject(UserService);

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
        next: (response) => {
          //TODO: implement routing to the app
          console.log(response)
        },
        error: (err: HttpErrorResponse) => console.error(err)
      });
    }
  }
}

