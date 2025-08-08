import { Component, ElementRef, inject, signal, viewChild } from '@angular/core';
import { NoteService } from '../../services/note-service';
import { FabMenu } from '../shared/fab-menu/fab-menu';
import { IFabOption, INote } from '../../models/types';
import { Backdrop } from '../shared/backdrop/backdrop';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-notes',
  imports: [FabMenu, Backdrop, ReactiveFormsModule, RouterLink],
  templateUrl: './notes.html',
  styleUrl: './notes.css'
})
export class Notes {

  isFormOpen = signal(false);
  createForm = new FormGroup({
    title: new FormControl('', {
      validators: [Validators.maxLength(20), Validators.minLength(3), Validators.required],
    }),
    content: new FormControl('', {
      validators: [Validators.minLength(1), Validators.required]
    }),
  });

  form = viewChild.required<ElementRef<HTMLFormElement>>('form');
  noteService = inject(NoteService)
  notes = this.noteService.notes

  menuOptions: IFabOption[] = [
    {
      label: 'create note',
      icon: 'icons/add-document.png',
      action: () => {
        console.log(this.notes())
        this.isFormOpen.set(true);
        this.form().nativeElement.style.display = "flex";
      }
    },
  ]

  async onCreateNote() {

    if (this.createForm.valid) {
      const newNote: INote = {
        id: 0,
        title: this.createForm.value.title!,
        content: this.createForm.value.content!,
      };
      this.noteService.createNote(newNote).subscribe({
        next: () => {
          this.onCloseForm();
          this.noteService.refresh();
        },
        error: (err) => console.error('error:', err)
      });
    } else {
      console.log("invalid form")
    }

  }

  onCloseForm() {
    this.isFormOpen.set(false);
    this.createForm.reset();
    this.form().nativeElement.style.display = "none";
  }

}
