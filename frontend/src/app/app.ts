import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Nav } from './components/nav/nav';
import { Sidebar } from './components/sidebar/sidebar';
import { Backdrop } from './components/shared/backdrop/backdrop';
import { AuthService } from './services/auth-service';
import { BootstrapService } from './services/bootstrap-service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Nav, Sidebar, Backdrop],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App implements OnInit {
  protected title = 'selfit';
  bootstrapped = signal(false);

  authService = inject(AuthService);
  bootstrapService = inject(BootstrapService);
  sideBarOpen = signal<boolean>(false);

  ngOnInit(): void {
    if (this.authService.checkToken()) {
      console.log("you are here");
      this.bootstrapService.loadAllData().subscribe({
        next: () => this.bootstrapped.set(true),
        error: () => this.authService.logout(),
      });
    } else {
      this.bootstrapped.set(true);
    }
  }

  onSideBarOpen() {
    this.sideBarOpen.set(true);
  }

  onSideBarClose() {
    this.sideBarOpen.set(false);
  }
}
