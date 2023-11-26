import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AlergijaListComponent } from './alergija-list.component';

describe('AlergijaListComponent', () => {
  let component: AlergijaListComponent;
  let fixture: ComponentFixture<AlergijaListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AlergijaListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AlergijaListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
