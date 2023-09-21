import { Component, Input, OnInit } from '@angular/core';
import { Pregled } from 'src/app/models/pregled.model';

@Component({
  selector: 'app-pregled-item',
  templateUrl: './pregled-item.component.html',
  styleUrls: ['./pregled-item.component.css']
})
export class PregledItemComponent implements OnInit {

  @Input() pregled: Pregled = new Pregled();

  constructor() { }

  ngOnInit(): void {
  }

  isTaken(): boolean {
    if (this.pregled.gradjanin != null) {
      return true;
    } else {
      return false;
    }
  }

}
