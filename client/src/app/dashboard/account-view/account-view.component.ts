import { Component } from '@angular/core';
import { User } from '../user';
import { CommonModule } from '@angular/common';
import { UserService } from '../../services/user.service';
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';
import { ApiError } from '../api-error';
import { InputErrorComponent } from '../../components/error/input-error/input-error.component';

@Component({
  selector: 'app-account-view',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, InputErrorComponent],
  templateUrl: './account-view.component.html',
  styleUrl: './account-view.component.css'
})

export class AccountViewComponent {
  user: User;
  success: string = '';
  errors: ApiError[];

  form = this.fb.group({
    FirstName: "",
    LastName: "",
    CompanyName: "",
    Address: "",
    Email: "",
    Phone: "",
    CompanyNum: "",
  })

  onSubmit(): void {
    this.success = ' '
    this.us.updateUser(this.form.value).subscribe({
      next: (value) => {

        setTimeout(() => {
          this.success = value.message;
        }, 500)
        this.errors = [];
      },
      error: (err) => {
        this.success = '';
        this.errors = err.error.error
      }
    })
  }

  constructor(
    private us: UserService,
    private fb: FormBuilder,
  ) { }

  ngOnInit() {
    this.us.getUser().subscribe((data) => {

      this.user = data

      this.form.setValue({
        FirstName: this.user.FirstName,
        LastName: this.user.LastName,
        CompanyName: this.user.CompanyName,
        Address: this.user.Address,
        Email: this.user.Email,
        Phone: this.user.Phone,
        CompanyNum: this.user.CompanyNum,
      })
    })
  }
}
