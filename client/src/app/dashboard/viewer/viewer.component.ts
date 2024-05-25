import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AccountViewComponent } from '../account-view/account-view.component';
import { ClientViewComponent } from '../client-view/client-view.component';
import { ProjectViewComponent } from '../project-view/project-view.component';
import { InvoiceViewComponent } from '../invoice-view/invoice-view.component';
import { InvoiceGenViewComponent } from '../invoice-gen-view/invoice-gen-view.component';

@Component({
  selector: 'app-viewer',
  standalone: true,
  imports: [
    CommonModule,
    AccountViewComponent,
    ClientViewComponent,
    ProjectViewComponent,
    InvoiceViewComponent,
    InvoiceGenViewComponent,
  ],
  templateUrl: './viewer.component.html',
  styleUrl: './viewer.component.css'
})
export class ViewerComponent {
  @Input() selected: number;

}
