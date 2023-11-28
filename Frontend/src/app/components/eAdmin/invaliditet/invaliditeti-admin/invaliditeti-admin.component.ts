import { Component, OnInit } from '@angular/core';
import { Alergija } from 'src/app/models/alergija.model';
import { Invaliditet } from 'src/app/models/invaliditet.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-invaliditeti-admin',
  templateUrl: './invaliditeti-admin.component.html',
  styleUrls: ['./invaliditeti-admin.component.css']
})
export class InvaliditetiAdminComponent implements OnInit {

  constructor(private healthcareService: HealthcareService) { }

  invaliditeti: Array<Invaliditet> = []

  ngOnInit(): void {
    this.healthcareService.GetSveInvaliditete().subscribe({
      next: (data) => {
        this.invaliditeti = data
      },
      error: (error) => {
        console.error(error)
      }
    })
  }

}
