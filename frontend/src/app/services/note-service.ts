import { HttpClient } from '@angular/common/http';
import { inject, Injectable, signal, Signal } from '@angular/core';
import { Observable } from 'rxjs';
import { INote } from '../models/types';
import { toSignal } from '@angular/core/rxjs-interop';

@Injectable({
  providedIn: 'root'
})
export class NoteService {

  private apiUrl = 'http://localhost:6969/api/notes';
  http = inject(HttpClient)

  private readonly _notes = signal<INote[]>([]);
  readonly notes = this._notes.asReadonly()

  constructor() {
    this.loadNotes();
  }

  loadNotes(): void {
    this.http.get<INote[]>(this.apiUrl).subscribe({
      next: (notes) => this._notes.set(notes),
      error: (err) => console.error('Failed to load notes', err)
    })
  }

  getNotes(): Signal<INote[]> {
    const request$: Observable<INote[]> = this.http.get<INote[]>(this.apiUrl);
    return toSignal(request$, { initialValue: [] });
  }

  createNote(note: INote): Observable<void> {
    console.log("service called with note:", note)
    return this.http.post<void>(this.apiUrl, note);
  }

  refresh(): void {
    this.loadNotes()
  }

}
