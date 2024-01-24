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
    tipPregleda: new FormControl(''),
    datumPregleda: new FormControl(''),
    satiPocetak: new FormControl(''),
    satiKraj: new FormControl('')
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
  today: string = new Date().toISOString().split('T')[0];

  ngOnInit(): void {
    this.pregledFormGroup = this.formBuilder.group({
      tipPregleda: ['', [Validators.required]],
      datumPregleda: ['', [Validators.required]],
      satiPocetak: ['', [Validators.required]],
      satiKraj: ['', [Validators.required]],
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

  onSubmit() {
    this.submittedPregled = true;
    this.submittedVakcina = true;

    if (this.pregledFormGroup.invalid) {
      return;
    }

    let pregled: AddPregled = new AddPregled();

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

    var datumPregleda: Date = this.pregledFormGroup.get('datumPregleda')?.value
    var satiPocetak: number = this.pregledFormGroup.get('satiPocetak')?.value
    var satiKraj: number = this.pregledFormGroup.get('satiKraj')?.value

    var PocetakPregleda: Date = new Date(
      datumPregleda
    )

    var ZavrsetakPregleda: Date = new Date(
      datumPregleda
    )

    PocetakPregleda.setHours(satiPocetak)
    ZavrsetakPregleda.setHours(satiKraj)
    
    pregled.pocetakPregleda = PocetakPregleda.getTime() / 1000
    pregled.zavrsetakPregleda = ZavrsetakPregleda.getTime() / 1000
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

    // firstScrollNumbers: number[] = Array.from({ length: 25 }, (_, i) => i);
  firstScrollNumbers: number[] = [7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23]
  secondScrollNumbers: number[] = [];
  
  selectedFirstScrollNumber: number = 0
  selectedSecondScrollNumber: number = 0

  onFirstScrollSelect(event: any) {
    this.selectedFirstScrollNumber = event.value;
    this.secondScrollNumbers = this.generateNumbersAfterSelection(this.selectedFirstScrollNumber);
    this.selectedSecondScrollNumber = 0;
  }

  onSecondScrollSelect(event: any) {
    this.selectedSecondScrollNumber = event.value;
  }

  private generateNumbersAfterSelection(startValue: number): number[] {
    return Array.from({ length: 24 - startValue }, (_, i) => startValue + i + 1);
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

  hasChosenType(): boolean {
    if (
      this.pregledFormGroup.get('tipPregleda')?.value == 'Vakcinacija' ||
      this.pregledFormGroup.get('tipPregleda')?.value == 'Običan'
    ) {
      return true;
    }
    return false;
  }

  hasChosenDate(): boolean {
    if (this.pregledFormGroup.get("datumPregleda")?.value == "") {
      return false;
    } else {
      return true;
    }
  }
}
