import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class TrackerService {

  private apiUrl = 'http://localhost:6969/api/tracker';
  http = inject(HttpClient)

}
