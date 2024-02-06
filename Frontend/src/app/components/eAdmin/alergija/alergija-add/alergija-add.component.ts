import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AddAlergija } from 'src/app/dto/addAlergija';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-alergija-add',
  templateUrl: './alergija-add.component.html',
  styleUrls: ['./alergija-add.component.css']
})
export class AlergijaAddComponent implements OnInit {

  constructor(
    private healthcareService: HealthcareService,
    private formBuilder: FormBuilder,
    private router: Router
  ) { }

  alergijaFormGroup: FormGroup = new FormGroup({
    naziv: new FormControl(''),
  })

  submitted = false
  alreadyExists = false

  ngOnInit(): void {
    this.alergijaFormGroup = this.formBuilder.group({
      naziv: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(30)]],
    })
  }

  get alergijaGroup(): { [key: string]: AbstractControl } {
    return this.alergijaFormGroup.controls
  }

  onSubmit() {
    this.submitted = true

    if (this.alergijaFormGroup.invalid) {
      return
    }

    let alergija: AddAlergija = new AddAlergija();

    alergija.naziv = this.alergijaFormGroup.get('naziv')?.value

    this.healthcareService.PostAlergija(alergija).subscribe({
      next: (data) => {
        this.router.navigate(['/Alergije-Admin'])
      },
      error: (error) => {
        if (error.status == 406) {
          this.alreadyExists = true
        }
        console.error(error)
      }
    })
  }

}
