import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { ClientPreviewComponent } from '../client-preview/client-preview.component';
import { ClientService } from '../../services/client.service';
import { Client } from '../client';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { debounceTime, distinctUntilChanged, switchMap, tap } from 'rxjs';
import { ClientCardComponent } from '../client-card/client-card.component';
import { CreateClientComponent } from '../create-client/create-client.component';
import { EditClientComponent } from '../edit-client/edit-client.component';


@Component({
  selector: 'app-client-view',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    ClientPreviewComponent,
    ClientCardComponent,
    CreateClientComponent,
    EditClientComponent
  ],
  templateUrl: './client-view.component.html',
  styleUrl: './client-view.component.css'
})
export class ClientViewComponent {
  clients: Client[];
  isLoading: boolean = false;
  toggleCard: boolean = false;
  clientCard: boolean = false;
  editCard: boolean = false;
  Search: FormControl = new FormControl("");
  Sort: FormControl = new FormControl("NEW")
  chosenClient: Client;

  constructor(private cs: ClientService) { }

  loadClient() {
    this.cs.getClient(this.chosenClient.ID!).subscribe((data) => {
      this.chosenClient = data;
    })
  }

  loadClients() {
    if (this.Search.value !== "") {
      this.cs.searchClient({ Search: this.Search.value, Sort: this.Sort.value }).subscribe(
        (data) => this.clients = data
      )
    } else {
      this.cs.getClients(this.Sort.value).subscribe((data) => this.clients = data)
    }
  }

  starred() {
    if (this.Search.value !== "") {
      this.cs.searchClient({ Search: this.Search.value, Sort: this.Sort.value }).subscribe(
        (data) => this.clients = data
      )
    } else {
      this.cs.getClients(this.Sort.value).subscribe((data) => this.clients = data)
    }
  }

  showClientCard() {
    this.clientCard = true;
    this.toggleCard = false;
    this.editCard = false;
  }

  showEditCard() {
    this.editCard = true;
    this.clientCard = false;
    this.toggleCard = false;
  }

  receiveClient(c: Client) {
    this.chosenClient = c;
  }

  receiveToggle($event: boolean) {
    this.toggleCard = $event;
    this.editCard = false;
  }

  rToggleEdit() {
    this.editCard = false;
    this.clientCard = true;
  }

  searchClients(s: string) {
    if (s) {
      return this.cs.searchClient({ Search: s, Sort: this.Sort.value }).pipe(tap(_ => this.isLoading = false))
    } else {
      return this.cs.getClients(this.Sort.value).pipe(tap(_ => this.isLoading = false))
    }
  }

  sortClients(s: string) {
    if (this.Search.value) {
      return this.cs.searchClient({ Search: this.Search.value, Sort: s }).pipe(tap(_ => this.isLoading = false))
    } else {
      return this.cs.getClients(s).pipe(tap(_ => this.isLoading = false))
    }
  }

  ngOnInit() {
    this.cs.getClients(this.Sort.value).subscribe((data) => {
      this.clients = data
    })

    this.Search.valueChanges.pipe(
      debounceTime(300),
      distinctUntilChanged(),
      tap(_ => {
        this.isLoading = true
      }),
      switchMap((searchTerm) => this.searchClients(searchTerm))
    ).subscribe({
      next: (data) => {
        this.clients = data
      },
    })

    this.Sort.valueChanges.pipe(
      debounceTime(300),
      distinctUntilChanged(),
      tap(_ => {
        this.isLoading = true
      }),
      switchMap((searchTerm) => this.sortClients(searchTerm))
    ).subscribe({
      next: (data) => {
        this.clients = data
      },
    })
  }
}
