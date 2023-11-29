import { Component, Input, OnInit } from '@angular/core';
import { Karton } from 'src/app/models/karton.model';

@Component({
  selector: 'app-karton-item',
  templateUrl: './karton-item.component.html',
  styleUrls: ['./karton-item.component.css']
})
export class KartonItemComponent implements OnInit {

  @Input() karton: Karton = new Karton()

  constructor() { }

  ngOnInit(): void {
  }

}
