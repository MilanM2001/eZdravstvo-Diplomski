import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VakcinaItemComponent } from './vakcina-item.component';

describe('VakcinaItemComponent', () => {
  let component: VakcinaItemComponent;
  let fixture: ComponentFixture<VakcinaItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VakcinaItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VakcinaItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
