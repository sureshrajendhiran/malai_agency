import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MainQiComponent } from './main-qi.component';

describe('MainQiComponent', () => {
  let component: MainQiComponent;
  let fixture: ComponentFixture<MainQiComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [MainQiComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(MainQiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
