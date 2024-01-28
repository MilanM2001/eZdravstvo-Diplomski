import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AddPotvrdaSmrti } from 'src/app/dto/addPotvrdaSmrti';
import { PotvrdaSmrti } from 'src/app/models/potvrdaSmrti.model';
import { RegistrarService } from 'src/app/services/registrar.service';

@Component({
  selector: 'app-potvrda-smrti-add',
  templateUrl: './potvrda-smrti-add.component.html',
  styleUrls: ['./potvrda-smrti-add.component.css']
})
export class PotvrdaSmrtiAddComponent implements OnInit {

  constructor(
    private registrarService: RegistrarService,
    private formBuilder: FormBuilder,
    private router: Router,
    private route: ActivatedRoute
  ) { }

  submitted = false
  today: string = new Date().toISOString().split('T')[0]
  jmbg = String(this.route.snapshot.paramMap.get('jmbg'));

  formGroup: FormGroup = new FormGroup({
    datumSmrti: new FormControl(''),
    mestoSmrti: new FormControl('')
  })

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      datumSmrti: ['', [Validators.required]],
      mestoSmrti: ['', [Validators.required]]
    })
  }

  get group(): { [key: string]: AbstractControl } {
    return this.formGroup.controls
  }

  onSubmit() {
    this.submitted = true

    if (this.formGroup.invalid) {
      return
    }

    let potvrda = new AddPotvrdaSmrti()

    var datum: Date = new Date(this.formGroup.get('datumSmrti')?.value)

    potvrda.datumSmrti = Number(datum.getTime()) / 1000
    potvrda.mestoSmrti = this.formGroup.get('mestoSmrti')?.value
    potvrda.jmbg = this.jmbg

    console.log(potvrda)

    this.registrarService.PostPotvrdaSmrti(potvrda)
      .subscribe({
        next: (data) => {
          this.router.navigate(['/Pregledi-Lekar'])
        },
        error: (error) => {
          console.error(error)
        }
      })
  }

}
