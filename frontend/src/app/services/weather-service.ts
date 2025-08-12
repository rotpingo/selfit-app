import { HttpClient } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { IWeatherRequest } from "../models/types";
import { Observable } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class WeatherService {


  private apiUrl = 'http://localhost:6969/api/weather';
  http = inject(HttpClient)

  addCity(req: IWeatherRequest): Observable<void> {
    return this.http.post<void>(this.apiUrl, req);
  }

}
