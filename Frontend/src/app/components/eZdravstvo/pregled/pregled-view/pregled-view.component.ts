import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Pregled } from 'src/app/models/pregled.model';
import { User } from 'src/app/models/user.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-pregled-view',
  templateUrl: './pregled-view.component.html',
  styleUrls: ['./pregled-view.component.css'],
})
export class PregledViewComponent implements OnInit {
  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private healthcareService: HealthcareService
  ) {}

  pregled: Pregled = new Pregled();
  pregled_id = String(this.route.snapshot.paramMap.get('id'));
  user: User = new User();

  ngOnInit(): void {
    this.healthcareService.GetPregledID(this.pregled_id).subscribe({
      next: (data) => {
        this.pregled = data;
      },
    });

    this.healthcareService.GetMe().subscribe({
      next: (data) => {
        this.user = data;
      },
      error: (error) => {
        console.log(error);
      },
    });
  }

  isMyPregled(): boolean {
    if (this.pregled.lekar.jmbg == this.user.jmbg) {
      return true;
    } else {
      return false;
    }
  }

  update() {
    this.healthcareService.ZakaziPregled(this.pregled_id).subscribe({
      next: () => {
        this.router.navigate(['/Pregledi-Gradjanin']);
      },
      error: (error) => {
        console.log(error);
      },
    });
  }

  delete() {
    this.healthcareService.DeletePregledID(this.pregled_id).subscribe({
      next: () => {
        this.router.navigate(['/Pregledi-Lekar']);
      },
      error: (error) => {
        console.log(error);
      },
    });
  }

  isTaken(): boolean {
    if (this.pregled.gradjanin != null) {
      return true;
    } else {
      return false;
    }
  }

  isVakcinacija(): boolean {
    if (this.pregled.tipPregleda == "Vakcinacija") {
      return true;
    } else {
      return false;
    }
  }
}
