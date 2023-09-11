import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PregledAddComponent } from './pregled-add.component';

describe('PregledAddComponent', () => {
  let component: PregledAddComponent;
  let fixture: ComponentFixture<PregledAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PregledAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PregledAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
