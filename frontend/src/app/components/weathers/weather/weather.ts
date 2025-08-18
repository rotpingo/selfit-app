import { Component, inject, signal } from '@angular/core';
import { FabMenu } from '../../shared/fab-menu/fab-menu';
import { Backdrop } from '../../shared/backdrop/backdrop';
import { IFabOption } from '../../../models/types';
import { WeatherService } from '../../../services/weather-service';
import { ActivatedRoute, Router } from '@angular/router';
import { toSignal } from '@angular/core/rxjs-interop';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-weather',
  imports: [FabMenu, Backdrop],
  templateUrl: './weather.html',
  styleUrl: './weather.css'
})
export class Weather {

  isFormOpen = signal(false);
  weatherService = inject(WeatherService)

  activeRoute = inject(ActivatedRoute);
  route = inject(Router);

  cityID = parseInt(this.activeRoute.snapshot.paramMap.get('id')!);
  cityData = toSignal(
    this.weatherService.getWeatherById(this.cityID),
    { initialValue: null }
  );

  menuOptions: IFabOption[] = [
    {
      label: 'delete task',
      icon: 'icons/delete-document.png',
      action: () => {
        this.weatherService.deleteCity(this.cityID).subscribe({
          next: () => {
            this.weatherService.refresh();
            this.route.navigate(["weather/"]);
          },
          error: (err: HttpErrorResponse) => console.error(err)
        });
      }
    },
    {
      label: 'edit task',
      icon: 'icons/file-edit.png',
      action: () => {
        console.log(this.cityData())
      }
    },
  ];

  onCloseForm() {
  }

}
