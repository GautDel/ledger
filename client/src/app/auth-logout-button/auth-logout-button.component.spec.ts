import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AuthLogoutButtonComponent } from './auth-logout-button.component';

describe('AuthLogoutButtonComponent', () => {
  let component: AuthLogoutButtonComponent;
  let fixture: ComponentFixture<AuthLogoutButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AuthLogoutButtonComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(AuthLogoutButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
