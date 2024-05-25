import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Client } from '../client';
import { CommonModule } from '@angular/common';
import { InputErrorComponent } from '../../components/error/input-error/input-error.component';
import { ApiError } from '../api-error';
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';
import { ClientService } from '../../services/client.service';
import { tap } from 'rxjs';

@Component({
  selector: 'app-edit-client',
  standalone: true,
  imports: [
    CommonModule,
    InputErrorComponent,
    ReactiveFormsModule
  ],
  templateUrl: './edit-client.component.html',
  styleUrl: './edit-client.component.css'
})
export class EditClientComponent {
  @Input() show: boolean;
  @Input() client: Client;
  @Output() clientUpdatedEvent = new EventEmitter<void>()
  @Output() toggleEvent = new EventEmitter<void>()
  errors: ApiError[];
  isLoading: boolean = false;
  success: string;

  form = this.fb.group({
    FirstName: [""],
    LastName: [""],
    Email: [""],
    Phone: [""],
    Address: [""],
    Country: [""],
    Description: [""],
  })

  toggleHandler() {
    this.toggleEvent.emit()
  }

  onSubmit() {
    this.cs.updateClient(this.form.value, this.client.ID!).pipe(
      tap(_ => {
        this.isLoading = true;
      })
    ).subscribe({
      next: (data) => {
        setTimeout(() => {
        this.isLoading = false;
        this.errors = [];
        this.success = data.message;
        this.clientUpdatedEvent.emit();
        }, 500)
      },
      error: (err) => {
        this.isLoading = false;
        this.success = "";
        this.errors = err.error.error;
      }
    })
  }

  constructor(private fb: FormBuilder, private cs: ClientService) { }

  ngOnChanges() {
    this.form.setValue({
      FirstName: this.client.FirstName,
      LastName: this.client.LastName,
      Email: this.client.Email,
      Phone: this.client.Phone,
      Address: this.client.Address,
      Country: this.client.Country,
      Description: this.client.Description,
    })
  }
}
