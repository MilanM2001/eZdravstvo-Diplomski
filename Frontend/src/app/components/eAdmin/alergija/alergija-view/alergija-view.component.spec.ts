import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AlergijaViewComponent } from './alergija-view.component';

describe('AlergijaViewComponent', () => {
  let component: AlergijaViewComponent;
  let fixture: ComponentFixture<AlergijaViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AlergijaViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AlergijaViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
