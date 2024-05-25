import { Component } from '@angular/core';
import { SideNavComponent } from './side-nav/side-nav.component';
import { ViewerComponent } from './viewer/viewer.component';
import { NavItem } from './nav-item';
import { HeaderComponent } from '../header/header.component';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [
    SideNavComponent,
    ViewerComponent,
    HeaderComponent
  ],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css'
})
export class DashboardComponent {

  selected: number = 2;

  navItems: NavItem[] = [
    { id: 1, name: "Account"},
    { id: 2, name: "Clients"},
    { id: 3, name: "Projects"},
    { id: 4, name: "Invoices"},
    { id: 5, name: "Invoice Generator"},
  ]

  receiveSelected($event: number) {
    this.selected = $event
  }
}
