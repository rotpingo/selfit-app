import { Component, inject, output } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-nav',
  imports: [],
  templateUrl: './nav.html',
  styleUrl: './nav.css'
})
export class Nav {

  menuClick = output<void>();
  router = inject(Router);

  onMenuClick() {
    this.menuClick.emit();
  }

  isNavVisible(): boolean {
    const hiddenRoutes = ['/auth/login', '/auth/register'];
    return !hiddenRoutes.includes(this.router.url)
  }
}
