import { Component, input, signal } from '@angular/core';
import { IFabOption } from '../../../models/types';
import { Backdrop } from '../backdrop/backdrop';

@Component({
  selector: 'app-fab-menu',
  imports: [Backdrop],
  templateUrl: './fab-menu.html',
  styleUrl: './fab-menu.css'
})
export class FabMenu {

  options = input<IFabOption[]>([])
  isOpen = signal(false);


  toggleMenu() {
    this.isOpen.update(open => !open);
  }

  runAction(option: IFabOption) {
    option.action();
    this.toggleMenu();
  }

  closeMenu() {
    this.toggleMenu();
  }
}
