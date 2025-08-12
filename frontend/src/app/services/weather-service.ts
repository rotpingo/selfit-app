import { HttpClient } from "@angular/common/http";
import { inject, Injectable, signal } from "@angular/core";
import { IGeocodeResult } from "../models/types";
import { debounceTime, Subject, switchMap } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class WeatherService {

  private searchTerms = new Subject<string>();

  private apiUrl = 'http://localhost:6969/api/weather';
  http = inject(HttpClient)

  geocodeResults = signal<IGeocodeResult[]>([]);

  constructor() {
    this.searchTerms
      .pipe(
        debounceTime(300),
        switchMap(term => this.http.get<IGeocodeResult[]>(`${this.apiUrl}/?city=${term}`))
      )
      .subscribe(results => this.geocodeResults.set(results));
  }

  searchCity(term: string) {
    this.searchTerms.next(term);
  }

}
