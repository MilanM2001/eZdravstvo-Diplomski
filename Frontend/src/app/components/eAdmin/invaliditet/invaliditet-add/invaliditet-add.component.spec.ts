import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvaliditetAddComponent } from './invaliditet-add.component';

describe('InvaliditetAddComponent', () => {
  let component: InvaliditetAddComponent;
  let fixture: ComponentFixture<InvaliditetAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InvaliditetAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InvaliditetAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
