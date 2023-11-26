import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AlergijeAdminComponent } from './alergije-admin.component';

describe('AlergijeAdminComponent', () => {
  let component: AlergijeAdminComponent;
  let fixture: ComponentFixture<AlergijeAdminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AlergijeAdminComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AlergijeAdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
