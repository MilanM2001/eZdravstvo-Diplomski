import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VakcinaListComponent } from './vakcina-list.component';

describe('VakcinaListComponent', () => {
  let component: VakcinaListComponent;
  let fixture: ComponentFixture<VakcinaListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VakcinaListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VakcinaListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
