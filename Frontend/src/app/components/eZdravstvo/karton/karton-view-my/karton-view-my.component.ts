import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Alergija } from 'src/app/models/alergija.model';
import { Invaliditet } from 'src/app/models/invaliditet.model';
import { Pregled } from 'src/app/models/pregled.model';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';
import { RegistrarService } from 'src/app/services/registrar.service';

@Component({
  selector: 'app-karton-view-my',
  templateUrl: './karton-view-my.component.html',
  styleUrls: ['./karton-view-my.component.css']
})
export class KartonViewMyComponent implements OnInit {

  constructor(private registrarService: RegistrarService,
    private healthcareService: HealthcareService,
    private route: ActivatedRoute) { }

  user: User = new User()
  options = ['Pregledi', 'Alergije', 'Invaliditeti'];

  pregledi: Pregled[] = []
  vakcinacije: Pregled[] = []
  alergije: Alergija[] = []
  invaliditeti: Invaliditet[] = []
  search_value: string = ""

  ngOnInit(): void {
    this.healthcareService.GetMe()
      .subscribe({
        next: (data) => {
          this.user = data
        },
        error: (error) => {
          console.error(error)
        },
        complete: () => {
          this.healthcareService.GetKartonJMBG(this.user.jmbg)
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
      })
  }

  search(search_option: string) {
    if (search_option == 'Pregledi') {
      this.search_value = "Pregledi"
      this.healthcareService.GetMojiPreglediGradjanin().subscribe({
        next: (data) => {
          this.pregledi = data;
        },
        error: (error) => {
          console.error(error);
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

  isMother(): boolean {
    if (this.user.pol == "Zenski") {
      return true
    }
    return false
  }

}
