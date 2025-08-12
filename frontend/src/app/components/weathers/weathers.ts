import { Component, ElementRef, inject, signal, viewChild } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { IFabOption, IGeocodeResult } from '../../models/types';
import { FabMenu } from '../shared/fab-menu/fab-menu';
import { Backdrop } from '../shared/backdrop/backdrop';
import { WeatherService } from '../../services/weather-service';

@Component({
  selector: 'app-weathers',
  imports: [ReactiveFormsModule, FabMenu, Backdrop],
  templateUrl: './weathers.html',
  styleUrl: './weathers.css'
})
export class Weathers {

  isFormOpen = signal(false);
  weatherService = inject(WeatherService)

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
    console.log("clicked")
  }

  onSearchCity(event: Event) {
    const input = event?.target as HTMLInputElement;
    const value = input.value;
    console.log(value)
    this.weatherService.searchCity(value);
  }

  onSelectCity(city: IGeocodeResult) {
    console.log('Selected City:', city);
  }

  onCloseForm() {
    this.isFormOpen.set(false);
    this.createForm.reset();
    this.form().nativeElement.style.display = "none";
  }

}
