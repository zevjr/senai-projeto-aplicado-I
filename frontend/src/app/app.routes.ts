import { Routes } from '@angular/router';
import {RiskListComponent} from './screens/risk-list/risk-list.component';
import {AuthGuard} from './guards/auth.guard';
import {MobileMenuComponent} from './layout/mobile-menu/mobile-menu.component';
import {LoginComponent} from './screens/login/login.component';
import {PublicComponent} from './layout/public/public.component';
import {RiskFormComponent} from './screens/risk-form/risk-form.component';

export const routes: Routes = [
  {
    path: '',
    component: PublicComponent,
    children: [
      { path: '', redirectTo: '/login', pathMatch: 'full' },
      { path: 'login', component: LoginComponent }
    ]
  },
  {
    path: '',
    component: MobileMenuComponent,
    // canActivate: [AuthGuard],
    children: [
      { path: 'risks', component: RiskListComponent },
      { path: 'risks/new', component: RiskFormComponent }
    ]
  },
  { path: '**', redirectTo: '/login' }
];
