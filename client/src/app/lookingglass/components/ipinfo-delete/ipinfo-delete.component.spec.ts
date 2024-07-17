import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IpinfoDeleteComponent } from './ipinfo-delete.component';

describe('IpinfoDeleteComponent', () => {
  let component: IpinfoDeleteComponent;
  let fixture: ComponentFixture<IpinfoDeleteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IpinfoDeleteComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IpinfoDeleteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
