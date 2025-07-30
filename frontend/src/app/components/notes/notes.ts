import { Component, inject } from '@angular/core';
import { NoteService } from '../../services/note-service';

@Component({
  selector: 'app-notes',
  imports: [],
  templateUrl: './notes.html',
  styleUrl: './notes.css'
})
export class Notes {

  noteService = inject(NoteService)

  onClick() {
    this.noteService.getNotes().subscribe({
      next: (response) => console.log(response),
      error: (error) => console.log(error)
    });
  }
}
