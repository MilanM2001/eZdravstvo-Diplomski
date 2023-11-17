import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VakcinaViewComponent } from './vakcina-view.component';

describe('VakcinaViewComponent', () => {
  let component: VakcinaViewComponent;
  let fixture: ComponentFixture<VakcinaViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VakcinaViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VakcinaViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
