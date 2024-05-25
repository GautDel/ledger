import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-confirmation-dialog',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './confirmation-dialog.component.html',
  styleUrl: './confirmation-dialog.component.css'
})
export class ConfirmationDialogComponent {
  @Input() show: boolean;
  @Output() toggleConfirmationEvent = new EventEmitter<void>()
  @Output() confirmDelEvent = new EventEmitter<void>()

  toggleConfirmation() {
    this.toggleConfirmationEvent.emit()
  }

  confirmDelHandler() {
    this.confirmDelEvent.emit()
  }
}

