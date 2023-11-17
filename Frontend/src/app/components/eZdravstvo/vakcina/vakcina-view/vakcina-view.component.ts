import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators, AbstractControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AddVakcina } from 'src/app/dto/addVakcina';
import { Vakcina } from 'src/app/models/vakcina.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-vakcina-view',
  templateUrl: './vakcina-view.component.html',
  styleUrls: ['./vakcina-view.component.css']
})
export class VakcinaViewComponent implements OnInit {

  vakcinaFormGroup: FormGroup = new FormGroup({
    naziv: new FormControl(''),
    kompanija: new FormControl('')
  })

  constructor(private route: ActivatedRoute,
    private router: Router,
    private formBuilder: FormBuilder,
    private healthcareService: HealthcareService
  ) { }

  submitted = false
  alreadyExists = false

  vakcina: Vakcina = new Vakcina();
  vakcina_id = String(this.route.snapshot.paramMap.get("id"))

  ngOnInit(): void {
    this.vakcinaFormGroup = this.formBuilder.group({
      naziv: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(20)]],
      kompanija: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(20)]]
    })

    this.healthcareService.GetVakcinaID(this.vakcina_id)
      .subscribe({
        next: (data) => {
          this.vakcina = data;

          this.vakcinaFormGroup.get('naziv')?.setValue(this.vakcina.naziv);
          this.vakcinaFormGroup.get('kompanija')?.setValue(this.vakcina.kompanija);
        },
        error: (error) => {
          console.error(error)
        }
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

    this.healthcareService.PutVakcina(vakcina, this.vakcina_id)
      .subscribe({
        next: (data) => {
          console.log("Uspeh")
          this.router.navigate(['/Vakcine']);
        },
        error: (error) => {
          console.log(error)
          if (error.status == 406) {
            this.alreadyExists = true
          }
        }
      })
  }

  delete(): void {
    this.healthcareService.DeleteVakcinaID(this.vakcina_id)
      .subscribe({
        next: () => {
          this.router.navigate(["/Vakcine"])
        },
        error: (error) => {
          console.error(error)
        }
      })
  }

}
