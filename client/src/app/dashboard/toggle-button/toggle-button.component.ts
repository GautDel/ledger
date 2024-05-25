import { Component, EventEmitter, Output, Input} from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-toggle-button',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './toggle-button.component.html',
  styleUrl: './toggle-button.component.css'
})
export class ToggleButtonComponent {

  constructor(){}

  @Input() open: boolean;
  @Output() toggleEvent = new EventEmitter<boolean>()

  toggleNav() {
    this.open = !this.open
    this.toggleEvent.emit(this.open)
  }
}
