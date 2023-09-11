import { Component, Input, OnInit } from '@angular/core';
import { Pregled } from 'src/app/models/pregled.model';

@Component({
  selector: 'app-pregled-list',
  templateUrl: './pregled-list.component.html',
  styleUrls: ['./pregled-list.component.css']
})
export class PregledListComponent implements OnInit {

  @Input() pregledi: Pregled[] = [];

  constructor() { }

  ngOnInit(): void {
  }

}
