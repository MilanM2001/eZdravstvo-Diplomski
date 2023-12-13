import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GradjaninAddDoctorComponent } from './gradjanin-add-doctor.component';

describe('GradjaninAddDoctorComponent', () => {
  let component: GradjaninAddDoctorComponent;
  let fixture: ComponentFixture<GradjaninAddDoctorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GradjaninAddDoctorComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GradjaninAddDoctorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
