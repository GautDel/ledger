import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-nav-icon',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './nav-icon.component.html',
  styleUrl: './nav-icon.component.css'
})
export class NavIconComponent {
  @Input() navItem: number;
}
