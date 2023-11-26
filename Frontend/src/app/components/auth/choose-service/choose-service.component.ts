import { Component, OnInit } from '@angular/core';
import {Credentials} from "../../../models/credentials";
import {Router} from "@angular/router";
import {StoreService} from "../../../services/store-service.service";

@Component({
  selector: 'app-choose-service',
  templateUrl: './choose-service.component.html',
  styleUrls: ['./choose-service.component.css']
})
export class ChooseServiceComponent implements OnInit {

  constructor(
    private router: Router,
    public storeService: StoreService
  ) { }

  ngOnInit(): void {
  }

  selectedAdmin(role: string, service: string) {
    localStorage.setItem('customRole', role)
    localStorage.setItem('service', service)
    this.router.navigate(['Welcome']).then()
  }

  selectService(role: string, service: string){
    if (role == this.storeService.getRoleFromToken()){
      localStorage.setItem('service', service)
      this.router.navigate(['regular-or-admin']).then()
    }else {
      localStorage.setItem('service', service)
      localStorage.setItem('customRole', 'Regular')
      this.router.navigate(['Welcome']).then()
    }
  }
}
