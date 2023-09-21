import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators, AbstractControl } from '@angular/forms';
import { Router } from '@angular/router';
import { AddPregled } from 'src/app/dto/addPregled';
import { Vakcina } from 'src/app/models/vakcina.model';
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
    vakcina: new FormControl(''),
    tipPregleda: new FormControl('')
  });

  constructor(private healthcareService: HealthcareService,
              private router: Router,
              private formBuilder: FormBuilder) 
              { }

  submitted = false;
  alreadyExists = false;
  vakcine: Vakcina[] = []
  vakcina: Vakcina = new Vakcina()
  vakcinaID: string = ""

  ngOnInit(): void {
    this.pregledFormGroup = this.formBuilder.group({
      pocetakPregleda: ['', [Validators.required]],
      zavrsetakPregleda: ['', [Validators.required]],
      vakcina: ['',],
      tipPregleda: ['', [Validators.required]]
    });

    this.healthcareService.GetSveVakcine()
      .subscribe({
        next: (data) => {
          this.vakcine = data
        },
        error: (error) => {
          console.log(error)
        }
      })
  }

  get pregledGroup(): { [key: string]: AbstractControl } {
    return this.pregledFormGroup.controls;
  }

  removeError() {
    this.alreadyExists = false;
  }

  isVakcinacija(): boolean {
    if (this.pregledFormGroup.get('tipPregleda')?.value == "Vakcinacija") {
      return true
    } 
    return false
  }

  izabrao(): boolean {
    if (this.pregledFormGroup.get('tipPregleda')?.value == "Vakcinacija" || this.pregledFormGroup.get('tipPregleda')?.value == "Običan") {
      return true
    }
    return false
  }

  onSubmit() {
    this.submitted = true;

    if (this.pregledFormGroup.invalid) {
      return;
    }

    let addPregled: AddPregled = new AddPregled();

    var PocetakPregleda: Date = new Date(this.pregledFormGroup.get('pocetakPregleda')?.value)
    var ZavrsetakPregleda: Date = new Date(this.pregledFormGroup.get('zavrsetakPregleda')?.value)

    addPregled.pocetakPregleda = Number(PocetakPregleda.getTime()) / 1000;
    addPregled.zavrsetakPregleda = Number(ZavrsetakPregleda.getTime()) / 1000;

    if (this.pregledFormGroup.get('tipPregleda')?.value == "Običan") {
      addPregled.tipPregleda = "Obican"
    } 
    if (this.pregledFormGroup.get('tipPregleda')?.value == "Vakcinacija") {
      addPregled.tipPregleda = "Vakcinacija"
    }

    if (this.pregledFormGroup.get('tipPregleda')?.value == "Vakcinacija") {
      this.vakcinaID = this.pregledFormGroup.get("vakcina")?.value
    }

    if (this.vakcinaID != "") {
      this.healthcareService.GetVakcinaID(this.vakcinaID)
        .subscribe({
          next: (data) => {
            this.vakcina = data
            addPregled.vakcina = this.vakcina
          },
          error: (error) => {
            console.log(error)
          }
      })
    }

    console.log(addPregled)
    // console.log(this.vakcina)
    this.healthcareService.PostPregled(addPregled)
      .subscribe({
        next: (data) => {
          this.router.navigate(['/Pregledi-Lekar']);
        },
        error: (error) => {
          console.log(error);
          // this.alreadyExists = true;
        }
      })

  }

  tipoviPregleda = new Array("Običan", "Vakcinacija")

}
