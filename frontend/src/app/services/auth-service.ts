import { HttpClient, HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from "@angular/common/http";
import { inject, Injectable, signal } from "@angular/core";
import { Router } from "@angular/router";
import { catchError, Observable, throwError } from "rxjs";
import { IAuthResponse, ISign } from "../models/types";

@Injectable({
  providedIn: 'root'
})
export class AuthService implements HttpInterceptor {

  private apiUrl = 'http://localhost:6969/api/auth';

  isLoggedIn = signal(false);

  http = inject(HttpClient);
  router = inject(Router);

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const token = this.getToken();

    if (token) {
      request = request.clone({
        setHeaders: { Authorization: token }
      });
    }

    return next.handle(request).pipe(
      catchError((error: any) => {
        if (error instanceof HttpErrorResponse) {
          if (error.status === 401) {
            this.removeToken();
            this.router.navigate(['/auth/login']);
            return throwError(() => new Error("Not authorized."));
          }
        }

        return throwError(() => error);
      })
    )
  }

  registerUser(user: ISign): Observable<void> {
    const url = `${this.apiUrl}/register`;
    return this.http.post<void>(url, user);
  }

  loginUser(user: ISign): Observable<IAuthResponse> {
    const url = `${this.apiUrl}/login`;
    return this.http.post<IAuthResponse>(url, user);
  }

  getToken(): string | null {
    return localStorage.getItem('token');
  }

  setToken(token: string) {
    localStorage.setItem('token', token);
  }

  removeToken() {
    localStorage.removeItem('token');
  }

  logout() {
    this.removeToken();
    this.router.navigate(['/auth/login']);
  }

}
