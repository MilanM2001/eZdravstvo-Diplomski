import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KartonItemComponent } from './karton-item.component';

describe('KartonItemComponent', () => {
  let component: KartonItemComponent;
  let fixture: ComponentFixture<KartonItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ KartonItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(KartonItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
