import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VakcineComponent } from './vakcine.component';

describe('VakcineComponent', () => {
  let component: VakcineComponent;
  let fixture: ComponentFixture<VakcineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VakcineComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VakcineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
