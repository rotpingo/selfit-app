import { Component, ElementRef, inject, signal, viewChild } from '@angular/core';
import { FabMenu } from '../../shared/fab-menu/fab-menu';
import { IFabOption, INote } from '../../../models/types';
import { NoteService } from '../../../services/note-service';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';
import { DatePipe } from '@angular/common';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

@Component({
  selector: 'app-note',
  imports: [FabMenu, DatePipe, ReactiveFormsModule],
  templateUrl: './note.html',
  styleUrl: './note.css'
})
export class Note {

  noteService = inject(NoteService);
  activeRoute = inject(ActivatedRoute);
  route = inject(Router);

  noteID = parseInt(this.activeRoute.snapshot.paramMap.get('id')!);
  note = this.noteService.getNoteByID(this.noteID);

  isFormOpen = signal(false);
  form = viewChild.required<ElementRef<HTMLFormElement>>('form');

  editForm = new FormGroup({
    title: new FormControl('', {
      validators: [Validators.maxLength(20), Validators.minLength(3), Validators.required],
    }),
    content: new FormControl('', {
      validators: [Validators.minLength(1), Validators.required]
    }),
  });


  menuOptions: IFabOption[] = [
    {
      label: 'delete note',
      icon: 'icons/delete-document.png',
      action: () => {
        this.noteService.deleteNote(this.noteID).subscribe({
          next: () => {
            this.noteService.refresh();
            this.route.navigate(["notes/"]);
          },
          error: (err: HttpErrorResponse) => alert(err.message),
        });
      }
    },
    {
      label: 'edit note',
      icon: 'icons/file-edit.png',
      action: () => {
        this.editForm.patchValue({
          title: this.note()?.title,
          content: this.note()?.content
        });
        this.isFormOpen.set(true);
        this.form().nativeElement.style.display = "flex";
      }
    },
  ];

  async onEditNote() {
    if (this.editForm.valid) {
      const newNote: INote = {
        id: this.note()?.id,
        title: this.editForm.value.title!,
        content: this.editForm.value.content!,
        createdAt: this.note()?.createdAt
      };

      this.noteService.editNote(newNote).subscribe({
        next: () => {
          this.noteService.refresh();
          this.onCloseForm();
        },
        error: (err) => console.error('error:', err)
      });
    } else {
      console.log("invalid form")
    }

  };

  onCloseForm() {
    this.isFormOpen.set(false);
    this.editForm.reset();
    this.form().nativeElement.style.display = "none";
  }

}
