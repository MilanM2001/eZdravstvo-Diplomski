import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PreglediLekarComponent } from './pregledi-lekar.component';

describe('PreglediLekarComponent', () => {
  let component: PreglediLekarComponent;
  let fixture: ComponentFixture<PreglediLekarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PreglediLekarComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PreglediLekarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
