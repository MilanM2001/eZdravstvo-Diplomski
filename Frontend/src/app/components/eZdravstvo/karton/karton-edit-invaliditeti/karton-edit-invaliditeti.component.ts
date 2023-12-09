import { Component, OnInit } from '@angular/core';
import { MatSelectChange } from '@angular/material/select';
import { ActivatedRoute, Router } from '@angular/router';
import { Alergija } from 'src/app/models/alergija.model';
import { Invaliditet } from 'src/app/models/invaliditet.model';
import { Karton } from 'src/app/models/karton.model';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';
import { RegistrarService } from 'src/app/services/registrar.service';

@Component({
  selector: 'app-karton-edit-invaliditeti',
  templateUrl: './karton-edit-invaliditeti.component.html',
  styleUrls: ['./karton-edit-invaliditeti.component.css']
})
export class KartonEditInvaliditetiComponent implements OnInit {

  constructor(
    private registrarService: RegistrarService,
    private healthcareService: HealthcareService,
    private route: ActivatedRoute,
    private router: Router) { }

  jmbg = String(this.route.snapshot.paramMap.get('jmbg'));
  user: User = new User()
  karton: Karton = new Karton()
  kartonInvaliditeti: Invaliditet[] = []
  invaliditeti: Invaliditet[] = []
  selectedInvaliditeti: Invaliditet[] = [];

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
          this.kartonInvaliditeti = this.karton.invaliditeti
        }, error: (error) => {
          console.error(error)
        }
      })

    this.healthcareService.GetSveInvaliditete()
      .subscribe({
        next: (data) => {
          this.invaliditeti = data

          this.selectedInvaliditeti = this.invaliditeti.filter(invaliditet =>
            this.kartonInvaliditeti.some(kartonInvaliditet => kartonInvaliditet.naziv === invaliditet.naziv)
          );
        }, error: (error) => {
          console.error(error)
        }
      })
  }

  onInvaliditetiSelectionChange(event: MatSelectChange): void {
    this.selectedInvaliditeti = event.value;

    this.kartonInvaliditeti = this.selectedInvaliditeti.map(selectedItem => {
      const existingItem = this.kartonInvaliditeti.find(alergija => alergija.naziv === selectedItem.naziv);

      if (existingItem) {
        return existingItem;
      } else {
        return { ...selectedItem } as Invaliditet;
      }
    });
  }

  submit(): void {
    this.karton.invaliditeti = this.selectedInvaliditeti;
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
