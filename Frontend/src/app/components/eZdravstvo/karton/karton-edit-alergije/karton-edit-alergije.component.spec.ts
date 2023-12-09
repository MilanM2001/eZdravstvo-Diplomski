import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KartonEditAlergijeComponent } from './karton-edit-alergije.component';

describe('KartonEditAlergijeComponent', () => {
  let component: KartonEditAlergijeComponent;
  let fixture: ComponentFixture<KartonEditAlergijeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ KartonEditAlergijeComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(KartonEditAlergijeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
