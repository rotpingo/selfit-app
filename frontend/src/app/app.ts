import { Component, inject, OnInit, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Nav } from './components/nav/nav';
import { Sidebar } from './components/sidebar/sidebar';
import { Backdrop } from './components/shared/backdrop/backdrop';
import { AuthService } from './services/auth-service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Nav, Sidebar, Backdrop],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App implements OnInit {
  protected title = 'selfit';

  authService = inject(AuthService)
  sideBarOpen = signal<boolean>(false);

  ngOnInit(): void {
    this.authService.checkToken();
  }

  onSideBarOpen() {
    this.sideBarOpen.set(true);
  }

  onSideBarClose() {
    this.sideBarOpen.set(false);
  }
}
