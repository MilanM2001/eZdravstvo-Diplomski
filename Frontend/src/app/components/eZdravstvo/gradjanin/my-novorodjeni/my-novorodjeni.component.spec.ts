import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyNovorodjeniComponent } from './my-novorodjeni.component';

describe('MyNovorodjeniComponent', () => {
  let component: MyNovorodjeniComponent;
  let fixture: ComponentFixture<MyNovorodjeniComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MyNovorodjeniComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MyNovorodjeniComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
