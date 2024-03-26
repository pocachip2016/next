import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PricePanelComponent } from './price-panel.component';

describe('PricePanelComponent', () => {
  let component: PricePanelComponent;
  let fixture: ComponentFixture<PricePanelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PricePanelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PricePanelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
