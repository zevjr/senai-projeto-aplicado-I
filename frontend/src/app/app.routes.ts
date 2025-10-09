import { Routes } from '@angular/router';

import { MobileMenuLayout } from './layout/mobile-menu/mobile-menu.layout';
import { MobilePublicLayout } from './layout/mobile-public/mobile-public.layout';
import { LoginScreen } from './screens/login/login.screen';
import { RegisterFormScreen } from './screens/risk-form/register-form.screen';
import { RegisterListScreen } from './screens/risk-list/register-list.screen';
import { RegisterComponent } from './register/register';

export const routes: Routes = [
  {
    path: '',
    component: MobilePublicLayout,
    children: [
      { path: '', redirectTo: 'login', pathMatch: 'full' },
      { path: 'login', component: LoginScreen },
      { path: 'register', component: RegisterComponent }
    ]
  },
  {
    path: '',
    component: MobileMenuLayout,
    // canActivate: [AuthGuard],
    children: [
      { path: 'risks', component: RegisterListScreen },
      { path: 'risks/new', component: RegisterFormScreen }
    ]
  },
  { path: '**', redirectTo: 'login' }
];
