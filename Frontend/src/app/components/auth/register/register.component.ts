import { Component, OnInit } from '@angular/core';
import {
  FormControl,
  FormGroup,
  FormBuilder,
  Validators,
  AbstractControl,
} from '@angular/forms';
import { Credentials } from '../../../models/credentials';
import { AuthService } from '../../../services/auth.service';
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
})
export class RegisterComponent implements OnInit {
  constructor(
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private router: Router,
    private _snackBar: MatSnackBar
  ) {}

  registerFormGroup: FormGroup = new FormGroup({
    jmbg: new FormControl(''),
    password: new FormControl(''),
    repeatPassword: new FormControl(''),
  });

  submitted = false;
  passwordNotMatch = false;
  JMBGNotExist = false;
  JMBGExists = false;

  ngOnInit(): void {
    this.registerFormGroup = this.formBuilder.group({
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
      repeatPassword: [
        '',
        [
          Validators.required,
          Validators.minLength(3),
          Validators.maxLength(30),
        ],
      ],
    });
  }

  get registerGroup(): { [key: string]: AbstractControl } {
    return this.registerFormGroup.controls;
  }

  onSubmit() {
    this.submitted = true;

    if (this.registerFormGroup.invalid) {
      return;
    }

    const credentials = new Credentials();

    credentials._id = 0;
    credentials.jmbg = this.registerFormGroup.get('jmbg')?.value;
    credentials.password = this.registerFormGroup.get('password')?.value;
    credentials.userType = 'Regular';
    if (
      this.registerFormGroup.get('password')?.value ==
      this.registerFormGroup.get('repeatPassword')?.value
    ) {
      this.authService.Registration(credentials).subscribe({
        next: (response) => {
          this.router.navigate(['Login']).then();
        },
        error: (error) => {
          if (error.status == 202) {
            this.JMBGExists = true
            this.JMBGNotExist = false
          } else if (error.status == 201) {
            this.JMBGExists = false
            this.JMBGNotExist = true;
          }
         
          this.passwordNotMatch = false;
        },
      });
    } else {
      this.passwordNotMatch = true;
    }
  }

  openSnackBar(message: string, action: string) {
    this._snackBar.open(message, action, {
      duration: 3500,
      verticalPosition: 'top',
    });
  }
}
