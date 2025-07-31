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
  notes = this.noteService.notes

}
