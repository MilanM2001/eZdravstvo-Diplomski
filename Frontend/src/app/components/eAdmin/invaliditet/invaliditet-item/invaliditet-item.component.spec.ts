import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvaliditetItemComponent } from './invaliditet-item.component';

describe('InvaliditetItemComponent', () => {
  let component: InvaliditetItemComponent;
  let fixture: ComponentFixture<InvaliditetItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InvaliditetItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InvaliditetItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
