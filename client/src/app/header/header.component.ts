import { Component, OnInit } from '@angular/core';
import { AuthButtonComponent } from '../auth-button/auth-button.component';
import { AuthService } from '@auth0/auth0-angular';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [AuthButtonComponent, CommonModule, RouterModule],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})

export class HeaderComponent implements OnInit {
  metadata = {};

  // Inject both AuthService and HttpClient
  constructor(public auth: AuthService) { }

  ngOnInit(): void {
  }
}
