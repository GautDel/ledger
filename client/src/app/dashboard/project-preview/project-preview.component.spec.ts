import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProjectPreviewComponent } from './project-preview.component';

describe('ProjectPreviewComponent', () => {
  let component: ProjectPreviewComponent;
  let fixture: ComponentFixture<ProjectPreviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ProjectPreviewComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ProjectPreviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
