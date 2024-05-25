import { CommonModule } from '@angular/common';
import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { DOCUMENT } from '@angular/common';
import { Router, RouterModule } from '@angular/router';

@Component({
  selector: 'app-auth-button',
  standalone: true,
  imports: [CommonModule, RouterModule],
  templateUrl: './auth-button.component.html',
  styleUrl: './auth-button.component.css'

})

export class AuthButtonComponent {
  constructor(
    @Inject(DOCUMENT) public document: Document,
    public auth: AuthService,
    public router: Router
  ) {}
}
