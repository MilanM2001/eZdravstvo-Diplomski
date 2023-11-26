import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import {StoreService} from "../../services/store-service.service";

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {

  constructor(
    private router: Router,
    public storeService: StoreService

  ) { }

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
