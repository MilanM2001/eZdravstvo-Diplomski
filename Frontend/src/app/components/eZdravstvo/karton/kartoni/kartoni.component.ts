import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators, AbstractControl } from '@angular/forms';
import { Karton } from 'src/app/models/karton.model';
import { HealthcareService } from 'src/app/services/healthcare.service';

@Component({
  selector: 'app-kartoni',
  templateUrl: './kartoni.component.html',
  styleUrls: ['./kartoni.component.css']
})
export class KartoniComponent implements OnInit {

  constructor(private healthcareService: HealthcareService,
    private formBuilder: FormBuilder) { }

  kartoni: Karton[] = []

  searchFormGroup: FormGroup = new FormGroup({
    search_input: new FormControl('')
  })

  submitted = false
  searchTimeout: any;

  ngOnInit(): void {
    this.searchFormGroup = this.formBuilder.group({
      search_input: ['', [Validators.required]]
    })

    this.healthcareService.GetSveKartone().subscribe({
      next: (data) => {
        this.kartoni = data
      },
      error: (error) => {
        console.error(error)
      }
    })
  }

  onSearchInputChange() {
    clearTimeout(this.searchTimeout);

    this.searchTimeout = setTimeout(() => {
      if (!this.searchFormGroup.get("search_input")?.hasError('required')) {
        this.search();
      }
    }, 1000);
  }

  search() {
    let search_input = this.searchFormGroup.get("search_input")?.value;
    console.log(search_input);
  }

  ngOnDestroy() {
    clearTimeout(this.searchTimeout);
  }

  get searchGroup(): { [key: string]: AbstractControl } {
    return this.searchFormGroup.controls
  }

}
