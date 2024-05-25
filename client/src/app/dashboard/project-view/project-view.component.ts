import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { Project } from '../project';
import { ProjectPreviewComponent } from '../project-preview/project-preview.component';
import { ProjectService } from '../../services/project.service';
import { debounceTime, distinctUntilChanged, switchMap, tap } from 'rxjs';

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

  constructor(private ps: ProjectService) {}

  showProjectCard() {
    this.projectCard = true;
    this.toggleCard = false;
    this.editCard = false;
  }

  rProject(p: Project) {
    this.chosenProject = p;
  }

  searchProjects(s: string) {
    if (s) {
      return this.ps.searchProjects({ Search: s, Sort: this.sort.value }).pipe(tap(_ => this.isLoading = false))
    } else {
      return this.ps.getProjects(this.sort.value).pipe(tap(_ => this.isLoading = false))
    }
  }

  sortClients(s: string) {
    if (this.search.value) {
      return this.ps.searchProjects({ Search: this.search.value, Sort: s }).pipe(tap(_ => this.isLoading = false))
    } else {
      return this.ps.getProjects(s).pipe(tap(_ => this.isLoading = false))
    }
  }

  ngOnInit() {
    this.ps.getProjects(this.sort.value).subscribe((data) => {
      this.projects = data
    })

    this.search.valueChanges.pipe(
      debounceTime(300),
      distinctUntilChanged(),
      tap(_ => {
        this.isLoading = true
      }),
      switchMap((searchTerm) => this.searchProjects(searchTerm))
    ).subscribe({
      next: (data) => {
        this.projects = data
      },
    })

    this.sort.valueChanges.pipe(
      debounceTime(300),
      distinctUntilChanged(),
      tap(_ => {
        this.isLoading = true
      }),
      switchMap((searchTerm) => this.sortClients(searchTerm))
    ).subscribe({
      next: (data) => {
        this.projects = data
      },
    })
  }
}
