import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { DOCUMENT } from '@angular/common';
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-auth-logout-button',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './auth-logout-button.component.html',
  styleUrl: './auth-logout-button.component.css'
})
export class AuthLogoutButtonComponent {

  constructor(@Inject(DOCUMENT) public document: Document, public auth: AuthService) {}
}
