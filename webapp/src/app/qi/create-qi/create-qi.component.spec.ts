import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateQiComponent } from './create-qi.component';

describe('CreateQiComponent', () => {
  let component: CreateQiComponent;
  let fixture: ComponentFixture<CreateQiComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [CreateQiComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(CreateQiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
