import { Component, OnInit } from '@angular/core';
import { Alergija } from 'src/app/models/alergija.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-alergije-admin',
  templateUrl: './alergije-admin.component.html',
  styleUrls: ['./alergije-admin.component.css']
})
export class AlergijeAdminComponent implements OnInit {

  constructor(private healthcareService: HealthcareService) { }

  alergije: Array<Alergija> = []

  ngOnInit(): void {
    this.healthcareService.GetSveAlergije().subscribe({
      next: (data) => {
        this.alergije = data
      },
      error: (error) => {
        console.error(error)
      }
    })
  }

}
