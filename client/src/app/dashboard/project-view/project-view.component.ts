import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { Project } from '../project';
import { ProjectPreviewComponent } from '../project-preview/project-preview.component';

@Component({
  selector: 'app-project-view',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    ProjectPreviewComponent
  ],
  templateUrl: './project-view.component.html',
  styleUrl: './project-view.component.css'
})
export class ProjectViewComponent {
  projects: Project[];
  toggleCard: boolean = false;
  projectCard: boolean = false;
  editCard: boolean = false;
  search: FormControl = new FormControl("");
  sort: FormControl = new FormControl("NEW");
  chosenProject: Project;
  isLoading: boolean = false;

  showProjectCard() {
    this.projectCard = true;
    this.toggleCard = false;
    this.editCard = false;
  }

  rProject(p: Project) {
    this.chosenProject = p;
  }
}
