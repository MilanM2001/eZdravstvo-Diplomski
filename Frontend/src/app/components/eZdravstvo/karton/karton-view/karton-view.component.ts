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
  options = ['Pregledi', 'Vakcinacije', 'Alergije', 'Invaliditeti'];

  pregledi: Pregled[] = []
  vakcinacije: Pregled[] = []
  alergije: Alergija[] = []
  invaliditeti: Invaliditet[] = []

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
  }

  search(search_option: string) {
    // if (search_option == 'Pregledi') {
    //   this.healthcareService.GetMojiPreglediGradjanin(this.jmbg).subscribe({
    //     next: (data) => {
    //       this.pregledi = data;
    //     },
    //     error: (error) => {
    //       console.log(error);
    //     },
    //   });
    // }
    // if (search_option == 'Zauzeti') {
    //   this.healthcareService.GetMojiZauzetiPreglediLekar().subscribe({
    //     next: (data) => {
    //       this.pregledi = data;
    //     },
    //     error: (error) => {
    //       console.log(error);
    //     },
    //   });
    // }
    // if (search_option == 'Svi') {
    //   this.healthcareService.GetMojiPreglediLekar().subscribe({
    //     next: (data) => {
    //       this.pregledi = data;
    //     },
    //     error: (error) => {
    //       console.log(error);
    //     },
    //   });
    // }
  }

}
