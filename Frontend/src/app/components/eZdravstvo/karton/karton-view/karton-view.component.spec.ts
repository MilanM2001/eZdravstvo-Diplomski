import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KartonViewComponent } from './karton-view.component';

describe('KartonViewComponent', () => {
  let component: KartonViewComponent;
  let fixture: ComponentFixture<KartonViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ KartonViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(KartonViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
