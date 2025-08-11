import { HttpClient } from '@angular/common/http';
import { inject, Injectable, Signal, signal } from '@angular/core';
import { ITracker } from '../models/types';
import { Observable } from 'rxjs';
import { toSignal } from '@angular/core/rxjs-interop';

@Injectable({
  providedIn: 'root'
})
export class TrackerService {

  private apiUrl = 'http://localhost:6969/api/tracker';
  http = inject(HttpClient)

  private readonly _trackers = signal<ITracker[]>([]);
  readonly trackers = this._trackers.asReadonly()

  constructor() {
    this.loadTrackers()
  }

  createTracker(tracker: ITracker): Observable<void> {
    return this.http.post<void>(this.apiUrl, tracker);
  }

  loadTrackers(): void {
    this.http.get<ITracker[]>(this.apiUrl).subscribe({
      next: (trackers) => this._trackers.set(trackers),
      error: (err) => console.error('Failed to load trackers', err)
    })
  }

  getTrackers(): Signal<ITracker[]> {
    const request$: Observable<ITracker[]> = this.http.get<ITracker[]>(this.apiUrl);
    return toSignal(request$, { initialValue: [] });
  }
}
