import { Component, OnInit } from '@angular/core';
import { StoreService } from 'src/app/services/store-service.service';
@Component({
  selector: 'app-regular-or-admin',
  templateUrl: './regular-or-admin.component.html',
  styleUrls: ['./regular-or-admin.component.css']
})
export class RegularOrAdminComponent implements OnInit {

  constructor(
    public storeService: StoreService,
  ) { }

  ngOnInit(): void {
  }

  setRole(role: string){
    localStorage.setItem('customRole', role)
  }

}
