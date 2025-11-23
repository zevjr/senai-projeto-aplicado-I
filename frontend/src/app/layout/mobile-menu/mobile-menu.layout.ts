import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import {RouterLink, RouterOutlet} from '@angular/router';
import {NgIf} from '@angular/common';

@Component({
  selector: 'mobile-menu',
  templateUrl: './mobile-menu.layout.html',
  imports: [
    RouterOutlet,
    RouterLink,
    NgIf
  ],
  styleUrls: ['./mobile-menu.layout.scss']
})
export class MobileMenuLayout {
  constructor(private authService: AuthService) { }

  get isAdmin(): boolean {
    return this.authService.getCurrentUser()?.role?.toLowerCase() === 'admin';
  }

  logout(): void {
    this.authService.logout();
  }
}
