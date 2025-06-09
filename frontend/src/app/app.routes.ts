import { Routes } from '@angular/router';
import {RiskListScreen} from './screens/risk-list/risk-list.screen';
import {MobileMenuLayout} from './layout/mobile-menu/mobile-menu.layout';
import {LoginScreen} from './screens/login/login.screen';
import {RiskFormScreen} from './screens/risk-form/risk-form.screen';
import {MobilePublicLayout} from "./layout/mobile-public/mobile-public.layout";

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
      { path: 'risks', component: RiskListScreen },
      { path: 'risks/new', component: RiskFormScreen }
    ]
  },
  { path: '**', redirectTo: '/login' }
];
