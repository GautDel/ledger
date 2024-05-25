import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Client } from '../client';

@Component({
  selector: 'app-client-preview',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './client-preview.component.html',
  styleUrl: './client-preview.component.css'
})
export class ClientPreviewComponent {
  @Input() client: Client;
  @Output() setClientEvent = new EventEmitter<Client>();
  @Output() showCard = new EventEmitter<void>();

  initProjects: number = 3;
  show: number = this.initProjects;

  toggleProjects() {
    this.show = this.client.Projects.length;
  }

  setClient() {
    this.setClientEvent.emit(this.client);
  }

  onClick() {
    this.showCard.emit();
  }
}
