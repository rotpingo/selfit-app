import { Component, output } from '@angular/core';

@Component({
  selector: 'app-backdrop',
  imports: [],
  templateUrl: './backdrop.html',
  styleUrl: './backdrop.css'
})
export class Backdrop {
  close = output<void>();

  onClick() {
    this.close.emit();
  }
}
