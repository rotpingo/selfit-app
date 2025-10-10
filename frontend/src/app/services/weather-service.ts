import { HttpClient } from "@angular/common/http";
import { inject, Injectable, signal } from "@angular/core";
import { Observable, tap, map, EMPTY, catchError } from "rxjs";
import { IWeatherRequest, IWeatherResponse } from "../models/types";

@Injectable({
  providedIn: 'root'
})
export class WeatherService {


  private apiUrl = 'http://localhost:6969/api/weather';
  http = inject(HttpClient)

  private readonly _weathers = signal<IWeatherRequest[]>([]);
  readonly weathers = this._weathers.asReadonly()

  addCity(req: IWeatherRequest): Observable<void> {
    return this.http.post<void>(this.apiUrl, req);
  }

  // loadWeatherCities(): void {
  //   this.http.get<IWeatherRequest[]>(this.apiUrl).subscribe({
  //     next: (weathers) => this._weathers.set(weathers),
  //     error: (err) => console.error('Failed to load weather cities', err)
  //   })
  // }

  loadWeatherCities(): Observable<void> {
    return this.http.get<IWeatherRequest[]>(this.apiUrl).pipe(
      tap(weathers => this._weathers.set(weathers)), // update your signal
      catchError(err => {
        console.error('Failed to load weather cities', err);
        return EMPTY;
      }),
      map(() => void 0) // ensure the observable returns void
    );
  }

  clearCitiesWeather(): void {
    this._weathers.set([])
  }

  getWeatherById(cityId: number): Observable<IWeatherResponse> {
    const url = `${this.apiUrl}/${cityId}`;
    return this.http.get<IWeatherResponse>(url);
  }

  deleteCity(cityID: number): Observable<void> {
    const url = `${this.apiUrl}/${cityID}`;
    return this.http.delete<void>(url)
  }


  refresh(): void {
    this.loadWeatherCities()
  }

}
