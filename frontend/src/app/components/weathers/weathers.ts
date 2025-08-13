import { Component, ElementRef, inject, signal, viewChild } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { IFabOption, IWeatherRequest } from '../../models/types';
import { FabMenu } from '../shared/fab-menu/fab-menu';
import { Backdrop } from '../shared/backdrop/backdrop';
import { WeatherService } from '../../services/weather-service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-weathers',
  imports: [ReactiveFormsModule, FabMenu, Backdrop],
  templateUrl: './weathers.html',
  styleUrl: './weathers.css'
})
export class Weathers {

  isFormOpen = signal(false);
  weatherService = inject(WeatherService)
  cities = this.weatherService.weathers

  createForm = new FormGroup({
    city: new FormControl('', {
      validators: [Validators.required],
    }),
  });

  form = viewChild.required<ElementRef<HTMLFormElement>>('form');

  menuOptions: IFabOption[] = [
    {
      label: 'create note',
      icon: 'icons/add-document.png',
      action: () => {
        this.isFormOpen.set(true);
        this.form().nativeElement.style.display = "flex";
      }
    },
  ];

  onAddCity() {
    if (this.createForm.valid) {
      const city: IWeatherRequest = {
        name: this.createForm.value.city!
      }
      this.weatherService.addCity(city).subscribe({
        next: () => {
          this.onCloseForm()
        },
        error: (err: HttpErrorResponse) => console.error(err)
      })
    }
  }

  onCloseForm() {
    this.isFormOpen.set(false);
    this.createForm.reset();
    this.form().nativeElement.style.display = "none";
  }

}
