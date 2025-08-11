import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { ITracker } from '../models/types';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TrackerService {

  private apiUrl = 'http://localhost:6969/api/tracker';
  http = inject(HttpClient)

  createTracker(tracker: ITracker): Observable<void> {
    return this.http.post<void>(this.apiUrl, tracker);
  }

}
