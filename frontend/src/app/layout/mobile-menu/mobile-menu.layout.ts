import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import {RouterLink, RouterOutlet} from '@angular/router';

@Component({
  selector: 'mobile-menu',
  templateUrl: './mobile-menu.layout.html',
  imports: [
    RouterOutlet,
    RouterLink
  ],
  styleUrls: ['./mobile-menu.layout.scss']
})
export class MobileMenuLayout {
  constructor(private authService: AuthService) { }

  logout(): void {
    this.authService.logout();
  }
}
