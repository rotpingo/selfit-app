import { HttpClient } from "@angular/common/http";
import { inject, Injectable, signal } from "@angular/core";
import { Observable } from "rxjs";
import { IWeatherRequest } from "../models/types";

@Injectable({
  providedIn: 'root'
})
export class WeatherService {


  private apiUrl = 'http://localhost:6969/api/weather';
  http = inject(HttpClient)

  private readonly _weathers = signal<IWeatherRequest[]>([]);
  readonly weathers = this._weathers.asReadonly()

  constructor() {
    this.loadWeatherCities();
    console.log(this.weathers())
  }

  addCity(req: IWeatherRequest): Observable<void> {
    return this.http.post<void>(this.apiUrl, req);
  }

  loadWeatherCities(): void {
    this.http.get<IWeatherRequest[]>(this.apiUrl).subscribe({
      next: (weathers) => this._weathers.set(weathers),
      error: (err) => console.error('Failed to load weather cities', err)
    })
  }

  getWeathers() {

  }

}
