import { Component, OnInit } from '@angular/core';
import {
  AbstractControl,
  FormBuilder,
  FormControl,
  FormGroup,
  Validators,
} from '@angular/forms';
import { Router } from '@angular/router';
import { AddVakcina } from 'src/app/dto/addVakcina';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-vakcina-add',
  templateUrl: './vakcina-add.component.html',
  styleUrls: ['./vakcina-add.component.css'],
})
export class VakcinaAddComponent implements OnInit {
  constructor(
    private healthcareService: HealthcareService,
    private formBuilder: FormBuilder,
    private router: Router
  ) { }

  vakcinaFormGroup: FormGroup = new FormGroup({
    naziv: new FormControl(''),
    kompanija: new FormControl(''),
  });

  submitted = false;
  alreadyExists = false;

  ngOnInit(): void {
    this.vakcinaFormGroup = this.formBuilder.group({
      naziv: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(20)]],
      kompanija: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(20)]],
    });
  }

  get vakcinaGroup(): { [key: string]: AbstractControl } {
    return this.vakcinaFormGroup.controls;
  }

  onSubmit() {
    this.submitted = true;

    if (this.vakcinaFormGroup.invalid) {
      return;
    }

    let vakcina: AddVakcina = new AddVakcina();

    vakcina.naziv = this.vakcinaFormGroup.get('naziv')?.value;
    vakcina.kompanija = this.vakcinaFormGroup.get('kompanija')?.value;

    console.log(vakcina);

    this.healthcareService.PostVakcina(vakcina).subscribe({
      next: (data) => {
        this.router.navigate(['/Vakcine']);
      },
      error: (error) => {
        console.log(error);
        if (error.status == 406) {
          this.alreadyExists = true;
        }
      },
    });
  }
}
