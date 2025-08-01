import { Routes } from '@angular/router';
import { Notes } from './components/notes/notes';
import { Note } from './components/notes/note/note';

export const routes: Routes = [
  {
    path: "notes",
    component: Notes
  },
  {
    path: "notes/:id",
    component: Note
  }
];
