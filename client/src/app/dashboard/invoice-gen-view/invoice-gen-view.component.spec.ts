import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvoiceGenViewComponent } from './invoice-gen-view.component';

describe('InvoiceGenViewComponent', () => {
  let component: InvoiceGenViewComponent;
  let fixture: ComponentFixture<InvoiceGenViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [InvoiceGenViewComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(InvoiceGenViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
