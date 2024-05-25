import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Project } from '../project';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-project-preview',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './project-preview.component.html',
  styleUrl: './project-preview.component.css'
})
export class ProjectPreviewComponent {

  @Input() project: Project;
  @Output() setProjectEvent = new EventEmitter<Project>();
  @Output() showCard = new EventEmitter<void>();

  initProjects: number = 3;
  show: number = this.initProjects;

  onClick() {
    this.showCard.emit();
  }
}

