import { HttpClient } from '@angular/common/http';
import { computed, inject, Injectable, Signal, signal } from '@angular/core';
import { ITracker } from '../models/types';
import { Observable, map, tap, catchError, EMPTY } from 'rxjs';
import { toSignal } from '@angular/core/rxjs-interop';

@Injectable({
  providedIn: 'root'
})
export class TrackerService {

  private apiUrl = 'http://localhost:6969/api/trackers';
  http = inject(HttpClient)

  private readonly _trackers = signal<ITracker[]>([]);
  readonly trackers = this._trackers.asReadonly()

  // loadTrackers(): void {
  //   this.http.get<ITracker[]>(this.apiUrl).subscribe({
  //     next: (trackers) => this._trackers.set(trackers),
  //     error: (err) => console.error('Failed to load trackers', err)
  //   })
  // }

  loadTrackers(): Observable<void> {
    return this.http.get<ITracker[]>(this.apiUrl).pipe(
      tap(trackers => this._trackers.set(trackers)), // update your signal
      catchError(err => {
        console.error('Failed to load trackers', err);
        return EMPTY;
      }),
      map(() => void 0) // ensure the observable returns void
    );
  }


  // Get the Tracker from the loadedNotes pool, no need for HTTP REQUEST
  getTrackerByID(trackerID: number): Signal<ITracker | undefined> {
    return computed(() => this._trackers().find((tracker) => tracker.id === trackerID));
  }

  getTrackers(): Signal<ITracker[]> {
    const request$: Observable<ITracker[]> = this.http.get<ITracker[]>(this.apiUrl);
    return toSignal(request$, { initialValue: [] });
  }

  createTracker(tracker: ITracker): Observable<void> {
    return this.http.post<void>(this.apiUrl, tracker);
  }

  editTracker(tracker: ITracker): Observable<void> {
    const url = `${this.apiUrl}/${tracker.id}`;
    return this.http.put<void>(url, tracker)
  }

  deleteTracker(trackerID: number): Observable<void> {
    const url = `${this.apiUrl}/${trackerID}`;
    return this.http.delete<void>(url)
  }

  refresh(): void {
    this.loadTrackers()
  }


}
