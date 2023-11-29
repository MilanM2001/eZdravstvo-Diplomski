import { Component, Input, OnInit } from '@angular/core';
import { Karton } from 'src/app/models/karton.model';

@Component({
  selector: 'app-karton-list',
  templateUrl: './karton-list.component.html',
  styleUrls: ['./karton-list.component.css']
})
export class KartonListComponent implements OnInit {

  @Input() kartoni: Karton[] = []

  constructor() { }

  ngOnInit(): void {
  }

}
