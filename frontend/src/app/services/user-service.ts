import { HttpClient } from '@angular/common/http';
import { inject, Injectable, signal } from '@angular/core';
import { IUser } from '../models/types';
import { AuthService } from './auth-service';
import { Observable, tap, map, catchError, EMPTY } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private apiUrl = 'http://localhost:6969/api/user';
  http = inject(HttpClient)
  authService = inject(AuthService)

  private readonly _user = signal<IUser | null>(null);
  readonly user = this._user.asReadonly()

  loadUser(): Observable<void> {
    return this.http.get<IUser>(this.apiUrl).pipe(
      tap(user => this._user.set(user)), // update your signal
      catchError(err => {
        console.error('Failed to load user data', err);
        return EMPTY;
      }),
      map(() => void 0) // ensure the observable returns void
    );
  }

  clearUser(): void {
    this._user.set(null)
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
