import { Component, Input } from '@angular/core';
import { ApiError } from '../../../dashboard/api-error';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-input-error',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './input-error.component.html',
  styleUrl: './input-error.component.css'
})
export class InputErrorComponent {
  @Input() errors: ApiError[];
  @Input() field: string;
}
