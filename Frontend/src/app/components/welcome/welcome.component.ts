import { Component, OnInit } from '@angular/core';
import { StoreService } from "../../services/store-service.service";

@Component({
  selector: 'app-welcome',
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit {

  constructor(
    private storeService: StoreService,
  ) { }

  ngOnInit(): void {
  }

  isLoggedIn(): boolean {
    let token = localStorage.getItem("authToken")
    if(token && token != "") {
      return true
    }

    return false
  }

}
