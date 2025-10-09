import { HttpClient } from '@angular/common/http';
import { inject, Injectable, signal } from '@angular/core';
import { IUser } from '../models/types';
import { AuthService } from './auth-service';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private apiUrl = 'http://localhost:6969/api/user';
  http = inject(HttpClient)
  authService = inject(AuthService)

  private readonly _user = signal<IUser | null>(null);
  readonly user = this._user.asReadonly()

  loadUser(): void {
    this.http.get<IUser>(this.apiUrl).subscribe({
      next: (user) => this._user.set(user),
      error: (err) => {
        console.error("ERROR: ", err);
        this.authService.logout();
      }
    })
  }

  editUser(user: IUser): Observable<void> {
    const url = `${this.apiUrl}/${user.id}`;
    return this.http.put<void>(url, user)
  }

  deleteUser(userId: number): Observable<void> {
    const url = `${this.apiUrl}/${userId}`;
    return this.http.delete<void>(url)
  }

}
