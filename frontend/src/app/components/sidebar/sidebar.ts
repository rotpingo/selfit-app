import { Component, input, output } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-sidebar',
  imports: [RouterLink],
  templateUrl: './sidebar.html',
  styleUrl: './sidebar.css'
})
export class Sidebar {
  isOpen = input<boolean>(false);
  navigate = output<void>();

  onNavigate() {
    this.navigate.emit();
  }
}
