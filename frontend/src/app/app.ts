import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Nav } from './components/nav/nav';
import { Sidebar } from './components/sidebar/sidebar';
import { Backdrop } from './components/shared/backdrop/backdrop';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Nav, Sidebar, Backdrop],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App {
  protected title = 'selfit';

  sideBarOpen = signal<boolean>(false);


  onSideBarOpen() {
    this.sideBarOpen.set(true);
  }

  onSideBarClose() {
    this.sideBarOpen.set(false);
  }
}
