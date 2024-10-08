import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators, AbstractControl } from '@angular/forms';
import { Router } from '@angular/router';
import { AddInvaliditet } from 'src/app/dto/addInvaliditet';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-invaliditet-add',
  templateUrl: './invaliditet-add.component.html',
  styleUrls: ['./invaliditet-add.component.css']
})
export class InvaliditetAddComponent implements OnInit {

  constructor(
    private healthcareService: HealthcareService,
    private formBuilder: FormBuilder,
    private router: Router
  ) { }

  invaliditetFormGroup: FormGroup = new FormGroup({
    naziv: new FormControl(''),
  })

  submitted = false
  alreadyExists = false

  ngOnInit(): void {
    this.invaliditetFormGroup = this.formBuilder.group({
      naziv: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(30)]],
    })
  }

  get invaliditetGroup(): { [key: string]: AbstractControl } {
    return this.invaliditetFormGroup.controls
  }

  onSubmit() {
    this.submitted = true

    if (this.invaliditetFormGroup.invalid) {
      return
    }

    let invaliditet: AddInvaliditet = new AddInvaliditet();

    invaliditet.naziv = this.invaliditetFormGroup.get('naziv')?.value

    this.healthcareService.PostInvaliditet(invaliditet).subscribe({
      next: (data) => {
        this.router.navigate(['/Invaliditeti-Admin'])
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
