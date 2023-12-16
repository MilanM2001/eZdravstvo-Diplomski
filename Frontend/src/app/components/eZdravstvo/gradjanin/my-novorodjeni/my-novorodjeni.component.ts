import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';
import { RegistrarService } from 'src/app/services/registrar.service';

@Component({
  selector: 'app-my-novorodjeni',
  templateUrl: './my-novorodjeni.component.html',
  styleUrls: ['./my-novorodjeni.component.css']
})
export class MyNovorodjeniComponent implements OnInit {

  constructor(private healthcareService: HealthcareService,
    private registrarService: RegistrarService) { }

  user: User = new User()
  jmbg: string = ""
  newborns: User[] = []

  ngOnInit(): void {
    this.healthcareService.GetMe()
      .subscribe({
        next: (data) => {
          this.user = data
          this.jmbg = data.jmbg
        },
        error: (error) => {
          console.error(error)
        },
        complete: () => {
          this.registrarService.GetNewbornsByMotherJMBG(this.jmbg)
            .subscribe({
              next: (data) => {
                this.newborns = data
              },
              error: (error) => {
                console.error(error)
              }
            })
        }
      })

  }

}
