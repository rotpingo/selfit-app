import { Component, inject } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { UserService } from '../../../services/user-service';
import { ISign } from '../../../models/types';
import { HttpErrorResponse, HttpResponse } from '@angular/common/http';

@Component({
  selector: 'app-register',
  imports: [ReactiveFormsModule],
  templateUrl: './register.html',
  styleUrl: './register.css'
})
export class Register {

  userService = inject(UserService);

  createForm = new FormGroup({
    email: new FormControl('', {
      validators: [Validators.email, Validators.required],
    }),
    password: new FormControl('', {
      validators: [Validators.required]
    }),
  });

  onCreateUser() {
    if (this.createForm.valid) {
      const newUser: ISign = {
        email: this.createForm.value.email!,
        password: this.createForm.value.password!
      }
      this.userService.registerUser(newUser).subscribe({
        next: (response) => {
          //TODO: implement routing to the app
          console.log(response)
        },
        error: (err: HttpErrorResponse) => console.error(err)
      });
    }
  }
}
