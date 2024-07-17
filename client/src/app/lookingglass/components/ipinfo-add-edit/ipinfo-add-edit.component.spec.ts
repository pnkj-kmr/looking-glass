import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IpinfoAddEditComponent } from './ipinfo-add-edit.component';

describe('IpinfoAddEditComponent', () => {
  let component: IpinfoAddEditComponent;
  let fixture: ComponentFixture<IpinfoAddEditComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IpinfoAddEditComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IpinfoAddEditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
