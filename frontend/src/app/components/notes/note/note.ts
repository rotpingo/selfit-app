import { Component, inject } from '@angular/core';
import { FabMenu } from '../../shared/fab-menu/fab-menu';
import { IFabOption } from '../../../models/types';
import { NoteService } from '../../../services/note-service';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-note',
  imports: [FabMenu],
  templateUrl: './note.html',
  styleUrl: './note.css'
})
export class Note {

  noteSerivce = inject(NoteService);
  activeRoute = inject(ActivatedRoute);
  route = inject(Router);

  noteID = parseInt(this.activeRoute.snapshot.paramMap.get('id')!);
  note = this.noteSerivce.getNoteByID(this.noteID);

  menuOptions: IFabOption[] = [
    {
      label: 'delete note',
      icon: '',
      action: () => {
        this.noteSerivce.deleteNote(this.noteID).subscribe({
          next: () => {
            this.noteSerivce.refresh();
            this.route.navigate(["notes/"]);
          },
          error: (err: HttpErrorResponse) => alert(err.message),
        });
      }
    },
    {
      label: 'edit note',
      icon: '',
      action: () => {
      }
    },
  ]
}
