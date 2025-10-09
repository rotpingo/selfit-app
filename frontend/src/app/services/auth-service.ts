import { HttpClient, HttpErrorResponse, HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from "@angular/common/http";
import { inject, Injectable, signal } from "@angular/core";
import { Router } from "@angular/router";
import { catchError, EMPTY, Observable, throwError } from "rxjs";
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

    if (!request.url.includes('/login') && !request.url.includes('/register')) {
      if (!token || this.isTokenExpired(token)) {
        this.logout();
        return EMPTY;
      }


      request = request.clone({
        setHeaders: { Authorization: token }
      });
    }
    return next.handle(request).pipe(
      catchError((error: any) => {
        if (error instanceof HttpErrorResponse) {
          if (error.status === 401) {
            this.logout();
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

  logout() {
    this.removeToken();
    this.isLoggedIn.set(false);
    this.router.navigate(['/auth/login']);
  }

  getToken(): string | null {
    return localStorage.getItem('token');
  }

  setToken(token: string) {
    localStorage.setItem('token', token);
    this.isLoggedIn.set(true);
  }

  removeToken() {
    localStorage.removeItem('token');
  }

  private decodeToken(token: string) {
    try {
      const payload = token.split('.')[1];
      const decoded = atob(payload);
      return JSON.parse(decoded);
    } catch {
      return null;
    }
  }

  isTokenExpired(token: string): boolean {
    const decoded = this.decodeToken(token);
    console.log("decoded: ", decoded)
    if (!decoded || !decoded.exp) {
      return true;
    }

    const expiry = decoded.exp * 1000; // converted to ms
    return Date.now() > expiry
  }

  checkToken(): boolean {
    const token = this.getToken();
    console.log("you are here")
    if (!token || this.isTokenExpired(token)) {
      this.logout();
      return false;
    }
    this.isLoggedIn.set(true);
    return true;
  }

}
