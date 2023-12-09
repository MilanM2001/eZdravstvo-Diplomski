import { Component, OnInit } from '@angular/core';
import { MatSelectChange } from '@angular/material/select';
import { ActivatedRoute, Router } from '@angular/router';
import { Alergija } from 'src/app/models/alergija.model';
import { Karton } from 'src/app/models/karton.model';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';
import { RegistrarService } from 'src/app/services/registrar.service';

@Component({
  selector: 'app-karton-edit-alergije',
  templateUrl: './karton-edit-alergije.component.html',
  styleUrls: ['./karton-edit-alergije.component.css']
})
export class KartonEditAlergijeComponent implements OnInit {
  constructor(
    private registrarService: RegistrarService,
    private healthcareService: HealthcareService,
    private route: ActivatedRoute,
    private router: Router) { }

  jmbg = String(this.route.snapshot.paramMap.get('jmbg'));
  user: User = new User()
  karton: Karton = new Karton()
  kartonAlergije: Alergija[] = []
  alergije: Alergija[] = []
  selectedAlergije: Alergija[] = [];

  ngOnInit(): void {
    this.registrarService.GetUserJMBG(this.jmbg)
      .subscribe({
        next: (data) => {
          this.user = data
        },
        error: (error) => {
          console.error(error)
        }
      })

    this.healthcareService.GetKartonJMBG(this.jmbg)
      .subscribe({
        next: (data) => {
          this.karton = data
          this.kartonAlergije = this.karton.alergije
        }, error: (error) => {
          console.error(error)
        }
      })

    this.healthcareService.GetSveAlergije()
      .subscribe({
        next: (data) => {
          this.alergije = data

          this.selectedAlergije = this.alergije.filter(alergija =>
            this.kartonAlergije.some(kartonAlergija => kartonAlergija.naziv === alergija.naziv)
          );
        }, error: (error) => {
          console.error(error)
        }
      })
  }

  onAlergijeSelectionChange(event: MatSelectChange): void {
    this.selectedAlergije = event.value;

    this.kartonAlergije = this.selectedAlergije.map(selectedItem => {
      const existingItem = this.kartonAlergije.find(alergija => alergija.naziv === selectedItem.naziv);

      if (existingItem) {
        return existingItem;
      } else {
        return { ...selectedItem } as Alergija;
      }
    });
  }

  submit(): void {
    this.karton.alergije = this.selectedAlergije;
    this.healthcareService.PutKarton(this.karton, this.jmbg)
      .subscribe({
        next: (data) => {
          this.router.navigate(['/Karton-View/' + this.jmbg]);
        },
        error: (error) => {
          console.error(error)
        }
      })
  }

}
