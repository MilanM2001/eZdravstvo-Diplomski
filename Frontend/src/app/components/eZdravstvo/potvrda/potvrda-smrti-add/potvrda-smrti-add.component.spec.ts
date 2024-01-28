import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PotvrdaSmrtiAddComponent } from './potvrda-smrti-add.component';

describe('PotvrdaSmrtiAddComponent', () => {
  let component: PotvrdaSmrtiAddComponent;
  let fixture: ComponentFixture<PotvrdaSmrtiAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PotvrdaSmrtiAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PotvrdaSmrtiAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
