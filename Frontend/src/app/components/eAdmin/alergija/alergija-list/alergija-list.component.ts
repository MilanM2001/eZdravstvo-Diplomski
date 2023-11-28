import { Component, Input, OnInit } from '@angular/core';
import { Alergija } from 'src/app/models/alergija.model';

@Component({
  selector: 'app-alergija-list',
  templateUrl: './alergija-list.component.html',
  styleUrls: ['./alergija-list.component.css']
})
export class AlergijaListComponent implements OnInit {

  @Input() alergije: Alergija[] = []

  constructor() { }

  ngOnInit(): void {
  }

}
