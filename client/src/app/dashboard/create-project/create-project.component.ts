import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { InputErrorComponent } from '../../components/error/input-error/input-error.component';
import { ApiError } from '../api-error';
import { Client } from '../client';
import { ClientService } from '../../services/client.service';
import { ProjectService } from '../../services/project.service';

@Component({
  selector: 'app-create-project',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    InputErrorComponent
  ],
  templateUrl: './create-project.component.html',
  styleUrl: './create-project.component.css'
})
export class CreateProjectComponent {
  @Input() toggleCard: boolean;
  @Output() toggleEvent = new EventEmitter<boolean>();
  success: string = "";
  isLoading: boolean = false;
  errors: ApiError[];
  clients: Client[];
  chosenClients: Client[] = [];
  filteredClients: Client[] = [];
  init: boolean = true;

  constructor(private cs: ClientService, private ps: ProjectService) {}

  projectForm = new FormGroup({
    Name: new FormControl(''),
    Description: new FormControl(''),
    Notes: new FormControl(''),
  });

  onSubmit() {
    const data = {
      Name: this.projectForm.get("Name")?.value,
      Description: this.projectForm.get("Description")?.value,
      Notes: this.projectForm.get("Notes")?.value,
      Clients: this.chosenClients
    }

    this.ps.createProject(data).subscribe({
      next: (data) => {
        console.log(data)
      }
    })
  }

  searchClients(e: Event) {
    const v = (e.target as HTMLInputElement).value

    this.filteredClients = this.clients.filter((client) => {
      if(client.FirstName.toUpperCase().includes(v.toUpperCase())) {
        return client
      }
      return
    })
  }

  addClient(c: Client) {
    this.chosenClients.push(c)
    const filtered = this.chosenClients.filter((i, p) =>{
      return this.chosenClients.indexOf(i) === p;
    });

    this.chosenClients = filtered;
  }

  removeClient(i: number){
    this.chosenClients.splice(i, 1);
  }

  toggleHandler() {
    this.toggleCard = !this.toggleCard
    this.toggleEvent.emit(this.toggleCard)
  }

  ngOnInit() {
    this.cs.getClients("NEW").subscribe({
      next: (data) => {
        this.clients = data;
      }
    })

    this.ps.getProjects("NEW").subscribe({
      next: (data) => {
        console.log(data)
      }
    })
  }
}
