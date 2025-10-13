import { Component, inject } from '@angular/core';
import { UserService } from '../../services/user-service';

@Component({
  selector: 'app-user',
  imports: [],
  templateUrl: './user.html',
  styleUrl: './user.css'
})
export class User {

  userService = inject(UserService)
  user = this.userService.user;
}
