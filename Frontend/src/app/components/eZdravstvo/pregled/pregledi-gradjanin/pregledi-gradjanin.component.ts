import { Component, OnInit } from '@angular/core';
import { Pregled } from 'src/app/models/pregled.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-pregledi-gradjanin',
  templateUrl: './pregledi-gradjanin.component.html',
  styleUrls: ['./pregledi-gradjanin.component.css']
})
export class PreglediGradjaninComponent implements OnInit {
  pregledi: Array<Pregled> = [];

  constructor(private healthcareService: HealthcareService) { }

  ngOnInit(): void {
    this.healthcareService.GetSviSlobodniPregledi().subscribe({
      next: (data) => {
        this.pregledi = data;
      },
      error: (error) => {
        console.log(error);
      },
    });
  }
}
