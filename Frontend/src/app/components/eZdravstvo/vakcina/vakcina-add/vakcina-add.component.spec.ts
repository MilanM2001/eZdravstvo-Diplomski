import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VakcinaAddComponent } from './vakcina-add.component';

describe('VakcinaAddComponent', () => {
  let component: VakcinaAddComponent;
  let fixture: ComponentFixture<VakcinaAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VakcinaAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VakcinaAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
