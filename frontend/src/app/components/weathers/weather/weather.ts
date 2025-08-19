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

  imgUrl = signal('');

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
        this.getImg(this.cityData()?.weather[0].id!)
      }
    },
  ];

  getImg(weatherCode: number) {
    switch (true) {
      case weatherCode >= 200 && weatherCode <= 232:
        return this.imgUrl.set("icons/weather-status/thunder-rain.png");
      case weatherCode >= 300 && weatherCode <= 321:
        return this.imgUrl.set("icons/weather-status/rain.png");
      case weatherCode >= 500 && weatherCode <= 531:
        return this.imgUrl.set("icons/weather-status/rain.png");
      case weatherCode >= 600 && weatherCode <= 601:
        return this.imgUrl.set("icons/weather-status/snow.png");
      case weatherCode == 602 || weatherCode == 620 || weatherCode == 621 || weatherCode == 622:
        return this.imgUrl.set("icons/weather-status/snow.png");
      case weatherCode >= 611 && weatherCode < 620:
        return this.imgUrl.set("icons/weather-status/snow-rain.png");
      case weatherCode >= 701 && weatherCode <= 781:
        return this.imgUrl.set("icons/weather-status/clouds.png");
      case weatherCode == 800:
        return this.imgUrl.set("icons/weather-status/sun.png");
      case weatherCode == 801:
        return this.imgUrl.set("icons/weather-status/sun-clouds.png");
      case weatherCode >= 802 && weatherCode <= 804:
        return this.imgUrl.set("icons/weather-status/clouds.png");
      default:
        return console.log("wrong weather status");
    }
  }

  onCloseForm() {
  }

}
