import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AlergijaItemComponent } from './alergija-item.component';

describe('AlergijaItemComponent', () => {
  let component: AlergijaItemComponent;
  let fixture: ComponentFixture<AlergijaItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AlergijaItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AlergijaItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
