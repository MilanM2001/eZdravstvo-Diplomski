import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvaliditetiAdminComponent } from './invaliditeti-admin.component';

describe('InvaliditetiAdminComponent', () => {
  let component: InvaliditetiAdminComponent;
  let fixture: ComponentFixture<InvaliditetiAdminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InvaliditetiAdminComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InvaliditetiAdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
