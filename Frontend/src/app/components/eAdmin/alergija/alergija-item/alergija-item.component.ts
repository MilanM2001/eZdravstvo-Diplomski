import { Component, Input, OnInit } from '@angular/core';
import { Alergija } from 'src/app/models/alergija.model';

@Component({
  selector: 'app-alergija-item',
  templateUrl: './alergija-item.component.html',
  styleUrls: ['./alergija-item.component.css']
})
export class AlergijaItemComponent implements OnInit {

  @Input() alergija: Alergija = new Alergija()

  constructor() { }

  ngOnInit(): void {
  }

}
