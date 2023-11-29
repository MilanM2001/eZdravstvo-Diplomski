import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KartonListComponent } from './karton-list.component';

describe('KartonListComponent', () => {
  let component: KartonListComponent;
  let fixture: ComponentFixture<KartonListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ KartonListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(KartonListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
