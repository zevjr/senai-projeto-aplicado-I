import { Routes } from '@angular/router';

import {MobileMenuLayout} from './layout/mobile-menu/mobile-menu.layout';
import {LoginScreen} from './screens/login/login.screen';
import {MobilePublicLayout} from "./layout/mobile-public/mobile-public.layout";
import {RegisterFormScreen} from "./screens/risk-form/register-form.screen";
import {RegisterListScreen} from "./screens/risk-list/register-list.screen";

export const routes: Routes = [
  {
    path: '',
    component: MobilePublicLayout,
    children: [
      { path: '', redirectTo: '/login', pathMatch: 'full' },
      { path: 'login', component: LoginScreen }
    ]
  },
  {
    path: '',
    component: MobileMenuLayout,
    // canActivate: [AuthGuard],
    children: [
      { path: 'risks', component: RegisterListScreen },
      { path: 'risks/new', component: RegisterFormScreen },
      { path: 'risks/edit/:uid', component: RegisterFormScreen }
    ]
  },
  { path: '**', redirectTo: '/login' }
];
