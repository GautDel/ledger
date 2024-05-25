import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { ClientService } from '../../services/client.service';
import { ApiError } from '../api-error';
import { InputErrorComponent } from '../../components/error/input-error/input-error.component';
import { tap } from 'rxjs';

@Component({
  selector: 'app-create-client',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    InputErrorComponent
  ],
  templateUrl: './create-client.component.html',
  styleUrl: './create-client.component.css'
})
export class CreateClientComponent {
  @Input() toggleCard: boolean;
  @Output() toggleEvent = new EventEmitter<boolean>();
  @Output() clientCreatedEvent = new EventEmitter<void>();
  success: string = "";
  isLoading: boolean = false;
  errors: ApiError[];

  constructor(private cs: ClientService) { }

  clientForm = new FormGroup({
    FirstName: new FormControl(''),
    LastName: new FormControl(''),
    Email: new FormControl(''),
    Phone: new FormControl(''),
    Address: new FormControl(''),
    Country: new FormControl(''),
    Description: new FormControl(''),
  });

  toggleHandler() {
    this.toggleCard = !this.toggleCard
    this.toggleEvent.emit(this.toggleCard)
  }

  onSubmit() {
    this.cs.createClient(this.clientForm.value).pipe(
      tap(() => this.isLoading = true)
    ).subscribe({
      next: (data) => {
        this.errors = [];
        setTimeout(() => {
          this.isLoading = false
          this.clientForm.reset();
          this.success = data.message;
          this.clientCreatedEvent.emit()
        }, 500)
      },
      error: (err) => {
        this.isLoading = false
        this.errors = err.error.error;
        this.success = "";
      }
    })
  }
}
