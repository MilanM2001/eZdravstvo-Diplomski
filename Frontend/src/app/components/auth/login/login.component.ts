import { Component, OnInit } from '@angular/core';
import {
  AbstractControl,
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { Credentials } from '../../../models/credentials';
import { AuthService } from '../../../services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  constructor(
    private authService: AuthService,
    private formBuilder: FormBuilder,
    private router: Router
  ) {}

  loginFormGroup: FormGroup = new FormGroup({
    jmbg: new FormControl(''),
    password: new FormControl(''),
  });

  notFound = false;
  submitted = false;

  ngOnInit(): void {
    this.loginFormGroup = this.formBuilder.group({
      jmbg: [
        '',
        [
          Validators.required,
          Validators.minLength(1),
          Validators.maxLength(30),
        ],
      ],
      password: [
        '',
        [
          Validators.required,
          Validators.minLength(3),
          Validators.maxLength(30),
        ],
      ],
    });
  }

  get loginGroup(): { [key: string]: AbstractControl } {
    return this.loginFormGroup.controls;
  }

  onSubmit() {
    this.submitted = true;

    if (this.loginFormGroup.invalid) {
      return;
    }

    let credentials = new Credentials();

    credentials.jmbg = this.loginFormGroup.get('jmbg')?.value;
    credentials.password = this.loginFormGroup.get('password')?.value;
    this.authService.Login(credentials).subscribe({
      next: (response) => {
        if (response != null) {
          if (response == 'JMBG not exist!') {
            localStorage.clear();
          } else if (response == "Password doesn't match!") {
            localStorage.clear();
          } else {
            localStorage.setItem('authToken', response);
            this.router.navigate(['/choose-service']).then();
          }
        }
      },
      error: (error) => {
        localStorage.clear();
        console.log(error.status);
        console.error(error);
        if ((error.status = 403)) {
          this.notFound = true;
        }
      },
    });
  }
}
