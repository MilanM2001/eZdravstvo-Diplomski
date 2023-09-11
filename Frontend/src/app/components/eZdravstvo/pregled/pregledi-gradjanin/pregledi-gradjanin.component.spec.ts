import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PreglediGradjaninComponent } from './pregledi-gradjanin.component';

describe('PreglediGradjaninComponent', () => {
  let component: PreglediGradjaninComponent;
  let fixture: ComponentFixture<PreglediGradjaninComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PreglediGradjaninComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PreglediGradjaninComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
