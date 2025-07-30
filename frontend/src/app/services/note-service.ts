import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';

@Injectable({
  providedIn: 'root'
})
export class NoteService {

  private apiUrl = 'http://localhost:6969/api/notes';
  http = inject(HttpClient)

  getNotes(): Observable<string> {
    return this.http.get<string>(this.apiUrl);
  }

}
