import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Alergija } from 'src/app/models/alergija.model';
import { Invaliditet } from 'src/app/models/invaliditet.model';
import { Pregled } from 'src/app/models/pregled.model';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';
import { RegistrarService } from 'src/app/services/registrar.service';

@Component({
  selector: 'app-karton-view',
  templateUrl: './karton-view.component.html',
  styleUrls: ['./karton-view.component.css']
})
export class KartonViewComponent implements OnInit {

  constructor(private registrarService: RegistrarService,
    private healthcareService: HealthcareService,
    private route: ActivatedRoute) { }

  jmbg = String(this.route.snapshot.paramMap.get('jmbg'));
  user: User = new User()
  options = ['Pregledi', 'Alergije', 'Invaliditeti'];

  pregledi: Pregled[] = []
  vakcinacije: Pregled[] = []
  alergije: Alergija[] = []
  invaliditeti: Invaliditet[] = []
  search_value: string = ""

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
          this.alergije = data.alergije
          this.invaliditeti = data.invaliditeti
        },
        error: (error) => {
          console.error(error)
        }
      })
  }

  search(search_option: string) {
    if (search_option == 'Pregledi') {
      this.search_value = "Pregledi"
      this.healthcareService.GetMojiPreglediGradjanin().subscribe({
        next: (data) => {
          this.pregledi = data;
          console.log(this.pregledi)
        },
        error: (error) => {
          console.log(error);
        },
      });
    }
    if (search_option == 'Alergije') {
      this.search_value = "Alergije"
    }
    if (search_option == 'Invaliditeti') {
      this.search_value = "Invaliditeti"
    }
  }

}
