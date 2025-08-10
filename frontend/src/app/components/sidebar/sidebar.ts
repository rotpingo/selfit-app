import { Component, inject, input, output } from '@angular/core';
import { RouterLink } from '@angular/router';
import { AuthService } from '../../services/auth-service';

@Component({
  selector: 'app-sidebar',
  imports: [RouterLink],
  templateUrl: './sidebar.html',
  styleUrl: './sidebar.css'
})
export class Sidebar {
  authService = inject(AuthService)
  isOpen = input<boolean>(false);
  navigate = output<void>();

  onNavigate() {
    this.navigate.emit();
  }

  onLogout() {
    console.log("in")
    this.authService.removeToken();
  }
}
