import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IpinfoComponent } from './ipinfo.component';

describe('IpinfoComponent', () => {
  let component: IpinfoComponent;
  let fixture: ComponentFixture<IpinfoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IpinfoComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IpinfoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
