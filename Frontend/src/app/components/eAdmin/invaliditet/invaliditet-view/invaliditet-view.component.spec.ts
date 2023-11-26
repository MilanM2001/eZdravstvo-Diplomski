import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvaliditetViewComponent } from './invaliditet-view.component';

describe('InvaliditetViewComponent', () => {
  let component: InvaliditetViewComponent;
  let fixture: ComponentFixture<InvaliditetViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InvaliditetViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InvaliditetViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
