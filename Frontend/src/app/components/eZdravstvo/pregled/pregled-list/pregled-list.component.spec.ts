import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PregledListComponent } from './pregled-list.component';

describe('PregledListComponent', () => {
  let component: PregledListComponent;
  let fixture: ComponentFixture<PregledListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PregledListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PregledListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
