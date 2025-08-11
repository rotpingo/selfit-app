import { Routes } from '@angular/router';
import { Notes } from './components/notes/notes';
import { Note } from './components/notes/note/note';
import { Tasks } from './components/tasks/tasks';
import { Task } from './components/tasks/task/task';
import { Login } from './components/authentication/login/login';
import { Register } from './components/authentication/register/register';
import { Tracker } from './components/tracker/tracker';

export const routes: Routes = [
  {
    path: "auth/login",
    component: Login
  },
  {
    path: "auth/register",
    component: Register
  },
  {
    path: "notes",
    component: Notes
  },
  {
    path: "notes/:id",
    component: Note
  },
  {
    path: "tasks",
    component: Tasks
  },
  {
    path: "tasks/:id",
    component: Task
  },
  {
    path: "tracker",
    component: Tracker
  }
];
