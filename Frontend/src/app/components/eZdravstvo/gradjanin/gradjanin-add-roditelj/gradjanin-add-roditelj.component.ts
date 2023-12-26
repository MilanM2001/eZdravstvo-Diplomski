import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';
import { RegistrarService } from 'src/app/services/registrar.service';

@Component({
  selector: 'app-gradjanin-add-roditelj',
  templateUrl: './gradjanin-add-roditelj.component.html',
  styleUrls: ['./gradjanin-add-roditelj.component.css']
})
export class GradjaninAddRoditeljComponent implements OnInit {

  constructor(private route: ActivatedRoute,
    private router: Router,
    private healthcareService: HealthcareService,
    private registrarService: RegistrarService,
    private formBuilder: FormBuilder) { }

  user_id = String(this.route.snapshot.paramMap.get('id'));
  user: User = new User()

  submitted = false
  fatherNotExist = false

  formGroup: FormGroup = new FormGroup({
    ime: new FormControl(''),
    prezime: new FormControl(''),
    jmbgOca: new FormControl('')
  })

  ngOnInit(): void {
    this.registrarService.GetUserID(this.user_id)
      .subscribe({
        next: (data) => {
          this.user = data
        },
        error: (error) => {
          console.error(error)
        }
      })

    this.formGroup = this.formBuilder.group({
      ime: ['', [Validators.required, Validators.minLength(1), Validators.maxLength(20)]],
      prezime: ['', [Validators.required, Validators.minLength(1), Validators.maxLength(20)]],
      jmbgOca: ['', [Validators.required, Validators.minLength(1), Validators.maxLength(13)]],
    })

  }

  get group(): { [key: string]: AbstractControl } {
    return this.formGroup.controls;
  }

  getRandomNumber(): string {
    const randomNumber = Math.floor(Math.random() * 500)

    const formattedNumber = randomNumber.toString().padStart(3, '0');

    return formattedNumber;
  }

  onSubmit() {
    this.submitted = true

    if (this.formGroup.invalid) {
      return
    }

    let jmbg: string

    this.user.ime = this.formGroup.get("ime")?.value
    this.user.prezime = this.formGroup.get("prezime")?.value
    this.user.jmbgOca = this.formGroup.get("jmbgOca")?.value

    const dateOfBirth = new Date(this.user.datumRodjenja * 1000);
    const formattedDate =
      ('0' + dateOfBirth.getDate()).slice(-2) +
      ('0' + (dateOfBirth.getMonth() + 1)).slice(-2) +
      dateOfBirth.getFullYear().toString().slice(-3);

    let mestoRodjenja = this.mestaRodjenja[this.user.mestoRodjenja]
    let randomNum = this.getRandomNumber()

    jmbg = formattedDate + mestoRodjenja + randomNum
    this.user.jmbg = jmbg
    console.log(this.user)

  }

  mestaRodjenja: { [key: string]: string } = {
    "Beograd": "71",
    "Kragujevac": "72",
    "Jagodina": "72",
    "Niš": "73",
    "Pirot": "73",
    "Toplica": "73",
    "Leskovac": "74",
    "Vranje": "74",
    "Zaječar": "75",
    "Bor": "75",
    "Smederevo": "76",
    "Požarevac": "76",
    "Mačva": "77",
    "Kolubara": "77",
    "Čačak": "78",
    "Kraljevo": "78",
    "Kruševac": "78",
    "Užice": "79",
    "Novi Sad": "80",
    "Sombor": "81",
    "Subotica": "82",
    "Vrbas": "83",
    "Kikinda": "84",
    "Zrenjanin": "85",
    "Pančevo": "86",
    "Vršac": "87",
    "Ruma": "88",
    "Sremska Mitrovica": "89"
  };

}
