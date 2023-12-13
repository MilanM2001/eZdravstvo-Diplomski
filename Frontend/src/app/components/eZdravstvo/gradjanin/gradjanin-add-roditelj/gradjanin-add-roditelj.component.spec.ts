import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GradjaninAddRoditeljComponent } from './gradjanin-add-roditelj.component';

describe('GradjaninAddRoditeljComponent', () => {
  let component: GradjaninAddRoditeljComponent;
  let fixture: ComponentFixture<GradjaninAddRoditeljComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GradjaninAddRoditeljComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GradjaninAddRoditeljComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
