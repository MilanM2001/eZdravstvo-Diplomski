import { Component, OnInit } from '@angular/core';
import { Pregled } from 'src/app/models/pregled.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-pregledi-lekar',
  templateUrl: './pregledi-lekar.component.html',
  styleUrls: ['./pregledi-lekar.component.css'],
})
export class PreglediLekarComponent implements OnInit {
  pregledi: Array<Pregled> = [];
  options = ['Slobodni', 'Zauzeti', 'Svi'];

  constructor(private healthcareService: HealthcareService) {}

  ngOnInit(): void {
    this.healthcareService.GetMojiPreglediLekar().subscribe({
      next: (data) => {
        this.pregledi = data;
      },
      error: (error) => {
        console.log(error);
      },
    });
  }

  search(search_option: string) {
    if (search_option == 'Slobodni') {
      this.healthcareService.GetMojiSlobodniPreglediLekar().subscribe({
        next: (data) => {
          this.pregledi = data;
        },
        error: (error) => {
          console.log(error);
        },
      });
    }
    if (search_option == 'Zauzeti') {
      this.healthcareService.GetMojiZauzetiPreglediLekar().subscribe({
        next: (data) => {
          this.pregledi = data;
        },
        error: (error) => {
          console.log(error);
        },
      });
    }
    if (search_option == 'Svi') {
      this.healthcareService.GetMojiPreglediLekar().subscribe({
        next: (data) => {
          this.pregledi = data;
        },
        error: (error) => {
          console.log(error);
        },
      });
    }
  }
}
