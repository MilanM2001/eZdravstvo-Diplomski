import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AlergijaAddComponent } from './alergija-add.component';

describe('AlergijaAddComponent', () => {
  let component: AlergijaAddComponent;
  let fixture: ComponentFixture<AlergijaAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AlergijaAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AlergijaAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
