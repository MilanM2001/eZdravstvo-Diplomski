import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AddVakcina } from 'src/app/dto/addVakcina';
import { Vakcina } from 'src/app/models/vakcina.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-vakcina-add',
  templateUrl: './vakcina-add.component.html',
  styleUrls: ['./vakcina-add.component.css']
})
export class VakcinaAddComponent implements OnInit {

  vakcinaFormGroup: FormGroup = new FormGroup({
    naziv: new FormControl(''),
    kompanija: new FormControl('')
  })

  constructor(private healthcareService: HealthcareService,
              private formBuilder: FormBuilder,
              private router: Router) 
  { }

  submitted = false
  alreadyExists = false

  ngOnInit(): void {
    this.vakcinaFormGroup = this.formBuilder.group({
      naziv: ['', [Validators.required]],
      kompanija: ['', [Validators.required]]
    })
  }

  get vakcinaGroup(): { [key: string]: AbstractControl } {
    return this.vakcinaFormGroup.controls
  }

  onSubmit() {
    this.submitted = true

    if (this.vakcinaFormGroup.invalid) {
      return
    }

    let vakcina: AddVakcina = new AddVakcina()

    vakcina.naziv = this.vakcinaFormGroup.get('naziv')?.value
    vakcina.kompanija = this.vakcinaFormGroup.get('kompanija')?.value

    console.log(vakcina)

    this.healthcareService.PostVakcina(vakcina)
      .subscribe({
        next: (data) => {
          console.log("Uspeh")
          this.router.navigate(['/Pregledi-Lekar']);
        },
        error: (error) => {
          console.log(error)
        }
      })
  }

}
