import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { ToggleButtonComponent } from '../toggle-button/toggle-button.component';
import { NavItem } from '../nav-item';
import { NavIconComponent } from './nav-icon/nav-icon.component';

@Component({
  selector: 'app-side-nav',
  standalone: true,
  imports: [
    CommonModule,
    ToggleButtonComponent,
    NavIconComponent
  ],
  templateUrl: './side-nav.component.html',
  styleUrl: './side-nav.component.css'
})
export class SideNavComponent {

  constructor() { }

  @Input() navItems: NavItem[];
  @Input() selected: number;
  @Output() selectedEvent = new EventEmitter<number>()

  open: boolean = false;
  locked: boolean = false;

  receiveToggle($event: boolean) {
    this.open = $event
  }

  itemSelect(id: number) {
    this.selected = id;
    this.selectedEvent.emit(this.selected)
  }

  lockNav() {
    this.locked = !this.locked;
  }
}
