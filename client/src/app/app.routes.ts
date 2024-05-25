import { Routes } from '@angular/router';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { AuthGuard } from '@auth0/auth0-angular';


export const routes: Routes = [
  {
    path: '',
    component: LandingPageComponent,
    title: 'Ledger - Dashboard'
  },
  {
    path: 'dashboard',
    canActivate: [AuthGuard],
    component: DashboardComponent,
    title: 'Ledger - Dashboard'
  }
];

