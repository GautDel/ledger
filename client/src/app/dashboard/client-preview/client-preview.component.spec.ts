import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClientPreviewComponent } from './client-preview.component';

describe('ClientPreviewComponent', () => {
  let component: ClientPreviewComponent;
  let fixture: ComponentFixture<ClientPreviewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ClientPreviewComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ClientPreviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
