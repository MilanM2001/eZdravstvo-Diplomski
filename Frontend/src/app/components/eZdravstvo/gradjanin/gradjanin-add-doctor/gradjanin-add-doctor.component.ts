import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, AbstractControl } from '@angular/forms';
import { Router } from '@angular/router';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';
import { RegistrarService } from 'src/app/services/registrar.service';

@Component({
  selector: 'app-gradjanin-add-doctor',
  templateUrl: './gradjanin-add-doctor.component.html',
  styleUrls: ['./gradjanin-add-doctor.component.css']
})
export class GradjaninAddDoctorComponent implements OnInit {

  constructor(
    private healthcareService: HealthcareService,
    private registrarService: RegistrarService,
    private formBuilder: FormBuilder,
    private router: Router) { }

  submitted = false
  motherNotExist = false

  formGroup: FormGroup = new FormGroup({
    jmbgMajke: new FormControl(''),
    datumRodjenja: new FormControl(''),
    mestoRodjenja: new FormControl(''),
    pol: new FormControl(''),
  })

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      jmbgMajke: ['', [Validators.required, Validators.minLength(1), Validators.maxLength(13)]],
      datumRodjenja: ['', [Validators.required]],
      mestoRodjenja: ['', [Validators.required, Validators.minLength(1), Validators.maxLength(30)]],
      pol: ['', [Validators.required]]
    })
  }

  get group(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }

  onSubmit() {
    this.submitted = true;

    if (this.formGroup.invalid) {
      return;
    }

    let user = new User()

    var datum: Date = new Date(this.formGroup.get('datumRodjenja')?.value)

    user.jmbgMajke = this.formGroup.get("jmbgMajke")?.value
    user.datumRodjenja = Number(datum.getTime()) / 1000
    user.mestoRodjenja = this.formGroup.get("mestoRodjenja")?.value
    user.pol = this.formGroup.get("pol")?.value

    this.registrarService.DoctorCreateUser(user)
      .subscribe({
        next: () => {
          this.router.navigate(['/Pregledi-Lekar'])
        },
        error: (error) => {
          console.log(error)
          if (error.status == 409) {
            this.motherNotExist = true
          }
        }
      })
  }

  drzave = new Array("Srbija", "Austrija", "Hrvatska", "Bosna", "Makedonija", "Bugarska", "Rumunija", "Crna Gora")
  polovi = new Array("Muski", "Zenski")

}
