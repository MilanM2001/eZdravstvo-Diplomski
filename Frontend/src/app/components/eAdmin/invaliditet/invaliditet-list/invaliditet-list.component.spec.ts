import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvaliditetListComponent } from './invaliditet-list.component';

describe('InvaliditetListComponent', () => {
  let component: InvaliditetListComponent;
  let fixture: ComponentFixture<InvaliditetListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InvaliditetListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InvaliditetListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
