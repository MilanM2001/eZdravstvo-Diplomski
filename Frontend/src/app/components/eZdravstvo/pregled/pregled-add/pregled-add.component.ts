import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators, AbstractControl } from '@angular/forms';
import { Router } from '@angular/router';
import { AddAppointment } from 'src/app/dto/addAppointment';
import { Pregled } from 'src/app/models/pregled.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-pregled-add',
  templateUrl: './pregled-add.component.html',
  styleUrls: ['./pregled-add.component.css']
})
export class PregledAddComponent implements OnInit {

  pregledFormGroup: FormGroup = new FormGroup({
    pocetakPregleda: new FormControl(''),
    zavrsetakPregleda: new FormControl(''),
    tipVakcine: new FormControl(''),
    tipPregleda: new FormControl('')
  });

  constructor(private healthcareService: HealthcareService,
              private router: Router,
              private formBuilder: FormBuilder) 
              { }

  submitted = false;
  alreadyExists = false;

  ngOnInit(): void {
    this.pregledFormGroup = this.formBuilder.group({
      pocetakPregleda: ['', [Validators.required]],
      zavrsetakPregleda: ['', [Validators.required]],
      tipVakcine: ['', [Validators.required]],
      tipPregleda: ['', [Validators.required]]
    });
  }

  get pregledGroup(): { [key: string]: AbstractControl } {
    return this.pregledFormGroup.controls;
  }

  removeError() {
    this.alreadyExists = false;
  }

  onSubmit() {
    this.submitted = true;

    if (this.pregledFormGroup.invalid) {
      return;
    }

    let addPregled: Pregled = new Pregled();

    var PocetakPregleda: Date = new Date(this.pregledFormGroup.get('pocetakPregleda')?.value)
    var ZavrsetakPregleda: Date = new Date(this.pregledFormGroup.get('zavrsetakPregleda')?.value)

    addPregled.pocetakPregleda = Number(PocetakPregleda.getTime()) / 1000;
    addPregled.zavrsetakPregleda = Number(ZavrsetakPregleda.getTime()) / 1000;
    addPregled.tipPregleda = this.pregledFormGroup.get('tipPregleda')?.value
    addPregled.tipVakcine = this.pregledFormGroup.get('tipVakcine')?.value

    console.log(addPregled)
    // this.healthcareService.AddAppointment(addAppointment)
    //   .subscribe({
    //     next: (data) => {
    //       this.router.navigate(['/Appointments-Doctor']);
    //     },
    //     error: (error) => {
    //       console.log(error);
    //       this.alreadyExists = true;
    //     }
    //   })

  }

  tipoviPregleda = new Array("Obican", "Vakcinacija")

}
