import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KartonViewMyComponent } from './karton-view-my.component';

describe('KartonViewMyComponent', () => {
  let component: KartonViewMyComponent;
  let fixture: ComponentFixture<KartonViewMyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ KartonViewMyComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(KartonViewMyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
