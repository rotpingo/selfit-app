import { HttpClient } from '@angular/common/http';
import { computed, inject, Injectable, signal, Signal } from '@angular/core';
import { EMPTY, Observable, tap, map, catchError } from 'rxjs';
import { INote } from '../models/types';

@Injectable({
  providedIn: 'root'
})
export class NoteService {

  private apiUrl = 'http://localhost:6969/api/notes';
  http = inject(HttpClient)

  private readonly _notes = signal<INote[]>([]);
  readonly notes = this._notes.asReadonly()

  // loadNotes(): void {
  //   this.http.get<INote[]>(this.apiUrl).subscribe({
  //     next: (notes) => this._notes.set(notes),
  //     error: (err) => console.error('Failed to load notes', err)
  //   })

  loadNotes(): Observable<void> {
    return this.http.get<INote[]>(this.apiUrl).pipe(
      tap(notes => this._notes.set(notes)), // update your signal
      catchError(err => {
        console.error('Failed to load notes', err);
        return EMPTY;
      }),
      map(() => void 0) // ensure the observable returns void
    );
  }

  // getNotes(): Signal<INote[]> {
  //   const request$: Observable<INote[]> = this.http.get<INote[]>(this.apiUrl);
  //   return toSignal(request$, { initialValue: [] });
  // }

  // Get the Note from the loadedNotes pool, no need for HTTP REQUEST
  getNoteByID(noteID: number): Signal<INote | undefined> {
    return computed(() => this._notes().find((note) => note.id === noteID));
  }


  createNote(note: INote): Observable<void> {
    return this.http.post<void>(this.apiUrl, note);
  }

  editNote(note: INote): Observable<void> {
    const url = `${this.apiUrl}/${note.id}`;
    return this.http.put<void>(url, note)
  }

  deleteNote(noteID: number): Observable<void> {
    const url = `${this.apiUrl}/${noteID}`;
    return this.http.delete<void>(url)
  }

  refresh(): void {
    this.loadNotes()
  }

}
