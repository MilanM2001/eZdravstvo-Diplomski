import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PregledItemComponent } from './pregled-item.component';

describe('PregledItemComponent', () => {
  let component: PregledItemComponent;
  let fixture: ComponentFixture<PregledItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PregledItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PregledItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
