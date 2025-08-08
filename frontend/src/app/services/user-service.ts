import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ISign } from '../models/types';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private apiUrl = 'http://localhost:6969/api/auth';
  http = inject(HttpClient)

  registerUser(user: ISign): Observable<void> {
    const url = `${this.apiUrl}/register`;
    return this.http.post<void>(url, user);
  }

  loginUser(user: ISign): Observable<void> {
    const url = `${this.apiUrl}/login`;
    return this.http.post<void>(url, user);
  }
}
