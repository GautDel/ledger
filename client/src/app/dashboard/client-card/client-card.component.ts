import { CommonModule } from '@angular/common';
import { Component, ElementRef, EventEmitter, HostListener, Input, Output } from '@angular/core';
import { Client } from '../client';
import { ConfirmationDialogComponent } from '../../components/confirmation-dialog/confirmation-dialog.component';
import { ClientService } from '../../services/client.service';

@Component({
  selector: 'app-client-card',
  standalone: true,
  imports: [
    CommonModule,
    ConfirmationDialogComponent
  ],
  templateUrl: './client-card.component.html',
  styleUrl: './client-card.component.css'
})
export class ClientCardComponent {
  @Input() client: Client;
  @Output() delClientEvent = new EventEmitter<void>();
  @Output() editCardEvent = new EventEmitter<void>();
  @Output() starredEvent = new EventEmitter<void>();
  @Output() hideEvent = new EventEmitter<void>();
  showOptions: boolean = false;
  showConfirmation: boolean = false;
  starred: boolean;
  clientY: number = 0;
  mouseY: number = 0;
  isDragging = false;

  constructor(private el: ElementRef, private cs: ClientService) { }

  starClient(clientID: number) {
    this.starred = !this.starred;

    this.cs.updateClientStar({ Starred: this.starred }, clientID).subscribe({
      next: () => {
        this.starredEvent.emit()
      },
      error: (err) => {
        console.log(err)
      }
    })
  }


  onMouseMove(event: TouchEvent) {
    if (this.isDragging) {
      event.preventDefault()
      if (this.mouseY < this.clientY - this.mouseY) {
        this.mouseY = event.changedTouches[0].clientY;
        this.clientY = this.mouseY
      } else {
        this.mouseY = event.changedTouches[0].clientY - 85;
        this.clientY = this.mouseY
      }
    }
  }

  onMouseDown() {
    this.isDragging = true;
  }

  onMouseExit() {
    this.clientY = 0;
    this.isDragging = false;
  }

  onMouseUp() {
    if (this.clientY > 300) {
      this.hideEvent.emit()
    }

    this.isDragging = false;
    this.clientY = 0
  }

  resetOptions() {
    this.showOptions = false;
  }

  @HostListener("document:click", ["$event.target"])

  public onClick(target: any) {
    const clickedInside = this.el.nativeElement.contains(target)
    if (!clickedInside) {
      this.resetOptions()
    }
  }

  rCloseConfirmEvent() {
    this.showConfirmation = false;
  }

  toggleEditHandler() {
    this.editCardEvent.emit()
  }

  rConfirmDel() {
    this.showConfirmation = false;
    this.cs.deleteClient(this.client.ID).subscribe({
      next: (data) => {
        this.delClientEvent.emit()
        console.log(data)
      },
      error: (err) => { console.log(err) },
    })
  }

  ngOnInit() {
    this.starred = this.client.Starred
  }

  ngOnChanges() {
    this.starred = this.client.Starred
  }
}
