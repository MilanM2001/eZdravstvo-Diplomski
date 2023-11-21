import { Component, OnInit } from '@angular/core';
import {
  FormGroup,
  FormControl,
  FormBuilder,
  Validators,
  AbstractControl,
} from '@angular/forms';
import { Router } from '@angular/router';
import { AddPregled } from 'src/app/dto/addPregled';
import { Vakcina } from 'src/app/models/vakcina.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-pregled-add',
  templateUrl: './pregled-add.component.html',
  styleUrls: ['./pregled-add.component.css'],
})
export class PregledAddComponent implements OnInit {
  constructor(
    private healthcareService: HealthcareService,
    private router: Router,
    private formBuilder: FormBuilder
  ) {}

  pregledFormGroup: FormGroup = new FormGroup({
    pocetakPregleda: new FormControl(''),
    zavrsetakPregleda: new FormControl(''),
    tipPregleda: new FormControl(''),
  });

  vakcinaFormGroup: FormGroup = new FormGroup({
    vakcina: new FormControl(''),
  });

  submittedPregled = false;
  submittedVakcina = false;
  alreadyExists = false;
  vakcine: Vakcina[] = [];
  vakcinaID: string = '';
  tipoviPregleda = new Array('Običan', 'Vakcinacija');

  ngOnInit(): void {
    this.pregledFormGroup = this.formBuilder.group({
      pocetakPregleda: ['', [Validators.required]],
      zavrsetakPregleda: ['', [Validators.required]],
      tipPregleda: ['', [Validators.required]],
    });

    this.vakcinaFormGroup = this.formBuilder.group({
      vakcina: ['', [Validators.required]],
    });

    this.healthcareService.GetSveVakcine().subscribe({
      next: (data) => {
        this.vakcine = data;
      },
      error: (error) => {
        console.error(error);
      },
    });
  }

  get pregledGroup(): { [key: string]: AbstractControl } {
    return this.pregledFormGroup.controls;
  }

  get vakcinaGroup(): { [key: string]: AbstractControl } {
    return this.vakcinaFormGroup.controls;
  }

  removeError() {
    this.alreadyExists = false;
  }

  isVakcinacija(): boolean {
    if (this.pregledFormGroup.get('tipPregleda')?.value == 'Vakcinacija') {
      return true;
    }
    return false;
  }

  hasChosen(): boolean {
    if (
      this.pregledFormGroup.get('tipPregleda')?.value == 'Vakcinacija' ||
      this.pregledFormGroup.get('tipPregleda')?.value == 'Običan'
    ) {
      return true;
    }
    return false;
  }

  onSubmit() {
    this.submittedPregled = true;

    if (this.pregledFormGroup.invalid) {
      return;
    }

    let pregled: AddPregled = new AddPregled();

    var PocetakPregleda: Date = new Date(
      this.pregledFormGroup.get('pocetakPregleda')?.value
    );
    var ZavrsetakPregleda: Date = new Date(
      this.pregledFormGroup.get('zavrsetakPregleda')?.value
    );

    if (this.pregledFormGroup.get('tipPregleda')?.value == 'Običan') {
      pregled.tipPregleda = 'Obican';
    }
    if (this.pregledFormGroup.get('tipPregleda')?.value == 'Vakcinacija') {
      pregled.tipPregleda = 'Vakcinacija';
    }

    if (this.pregledFormGroup.get('tipPregleda')?.value == 'Vakcinacija') {
      if (this.vakcinaFormGroup.invalid) {
        return;
      }

      this.vakcinaID = this.vakcinaFormGroup.get('vakcina')?.value;
    } else {
      this.vakcinaID = ""
    }

    pregled.pocetakPregleda = Number(PocetakPregleda.getTime()) / 1000;
    pregled.zavrsetakPregleda = Number(ZavrsetakPregleda.getTime()) / 1000;
    pregled.vakcinaID = this.vakcinaID

    console.log(pregled);

    this.healthcareService.PostPregled(pregled).subscribe({
      next: (data) => {
        this.router.navigate(['/Pregledi-Lekar']);
      },
      error: (error) => {
        console.error(error);
        if (error.status == 406) {
          this.alreadyExists = true;
        }
      },
    });
  }
}
