import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import {RouterLink, RouterOutlet} from '@angular/router';

@Component({
  selector: 'app-mobile-menu',
  templateUrl: './mobile-menu.component.html',
  imports: [
    RouterOutlet,
    RouterLink
  ],
  styleUrls: ['./mobile-menu.component.scss']
})
export class MobileMenuComponent {
  constructor(private authService: AuthService) { }

  logout(): void {
    this.authService.logout();
  }
}
