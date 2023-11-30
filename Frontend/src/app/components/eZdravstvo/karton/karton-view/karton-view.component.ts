import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { User } from 'src/app/models/user.model';
import { RegistrarService } from 'src/app/services/registrar.service';

@Component({
  selector: 'app-karton-view',
  templateUrl: './karton-view.component.html',
  styleUrls: ['./karton-view.component.css']
})
export class KartonViewComponent implements OnInit {

  constructor(private registrarService: RegistrarService,
    private route: ActivatedRoute) { }

  jmbg = String(this.route.snapshot.paramMap.get('jmbg'));
  user: User = new User()

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

}
