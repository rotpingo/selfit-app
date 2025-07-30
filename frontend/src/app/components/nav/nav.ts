import { Component, output } from '@angular/core';

@Component({
  selector: 'app-nav',
  imports: [],
  templateUrl: './nav.html',
  styleUrl: './nav.css'
})
export class Nav {

  menuClick = output<void>();

  onMenuClick() {
    this.menuClick.emit();
  }
}
