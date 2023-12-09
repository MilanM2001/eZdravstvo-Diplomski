import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KartonEditInvaliditetiComponent } from './karton-edit-invaliditeti.component';

describe('KartonEditInvaliditetiComponent', () => {
  let component: KartonEditInvaliditetiComponent;
  let fixture: ComponentFixture<KartonEditInvaliditetiComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ KartonEditInvaliditetiComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(KartonEditInvaliditetiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
