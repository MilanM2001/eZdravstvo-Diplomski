import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { StoreService } from "../../services/store-service.service";
import { HealthcareService } from 'src/app/services/healthcare.service';
import { User } from 'src/app/models/user.model';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {

  constructor(
    private router: Router,
    public storeService: StoreService,
    public healthcareService: HealthcareService
  ) { }

  user: User = new User()

  ngOnInit(): void {
    this.healthcareService.GetMe()
      .subscribe({
        next: (data) => {
          this.user = data
        },
        error: (error) => {
          console.error(error)
        }
      })
  }

  isMother(): boolean {
    if (this.user.pol == "Zenski") {
      return true
    }
    return false
  }

  isLoggedIn(): boolean {
    if (localStorage.getItem("authToken") != null) {
      return true;
    }
    else {
      return false;
    }
  }

  logout() {
    localStorage.clear();
    this.router.navigate(['/Login']).then();
  }

}
