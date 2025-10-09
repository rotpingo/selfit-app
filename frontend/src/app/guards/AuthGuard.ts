import { inject, Injectable } from "@angular/core";
import { CanActivate, Router } from "@angular/router";
import { AuthService } from "../services/auth-service";

@Injectable({ providedIn: 'root' })
export class AuthGuard implements CanActivate {

  authService = inject(AuthService);
  router = inject(Router);

  canActivate(): boolean {
    const valid = this.authService.checkToken();

    if (!valid) {
      this.router.navigate(['/auth/login'])
      return false;
    }

    return true;
  }

}
