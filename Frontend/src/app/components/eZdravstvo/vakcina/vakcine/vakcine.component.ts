import { Component, OnInit } from '@angular/core';
import { Vakcina } from 'src/app/models/vakcina.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-vakcine',
  templateUrl: './vakcine.component.html',
  styleUrls: ['./vakcine.component.css']
})
export class VakcineComponent implements OnInit {

  constructor(private healthcareService: HealthcareService) { }
  vakcine: Vakcina[] = []

  ngOnInit(): void {
    this.healthcareService.GetSveVakcine()
      .subscribe({
        next: (data) => {
          this.vakcine = data
        },
        error: (error) => {
          console.log(error)
        }
      })
  }

}
