import { HttpClient } from '@angular/common/http';
import { computed, inject, Injectable, signal, Signal } from '@angular/core';
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
      next: (notes) => {
        const loadedNotes = this.toCamelCaseNotes(notes);
        this._notes.set(loadedNotes)
      },
      error: (err) => console.error('Failed to load notes', err)
    })
  }

  getNotes(): Signal<INote[]> {
    const request$: Observable<INote[]> = this.http.get<INote[]>(this.apiUrl);
    return toSignal(request$, { initialValue: [] });
  }

  // getNoteById(noteId: number): Signal<INote | undefined> {
  //   const url = `${this.apiUrl}/notes/${noteId}`;
  //   const request$: Observable<INote> = this.http.get<INote>(url);
  //   return toSignal(request$, { initialValue: undefined });
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

  // transforms the json snake_case format data into camel Case
  toCamelCaseNote(note: any): INote {
    return {
      id: note.id,
      title: note.title,
      content: note.content,
      createdAt: note.created_at,
      updatedAt: note.updated_at
    };
  }

  // transforms the json snake_case format data into camel Case for the entire array
  toCamelCaseNotes(notes: any[]): INote[] {
    return notes.map(this.toCamelCaseNote);
  }
}
