import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PregledViewComponent } from './pregled-view.component';

describe('PregledViewComponent', () => {
  let component: PregledViewComponent;
  let fixture: ComponentFixture<PregledViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PregledViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PregledViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
