import { Routes } from '@angular/router';
import { Notes } from './components/notes/notes';
import { Note } from './components/notes/note/note';
import { Tasks } from './components/tasks/tasks';
import { Task } from './components/tasks/task/task';
import { Login } from './components/authentication/login/login';
import { Register } from './components/authentication/register/register';
import { Trackers } from './components/trackers/trackers';
import { Tracker } from './components/trackers/tracker/tracker';
import { Weathers } from './components/weathers/weathers';
import { Weather } from './components/weathers/weather/weather';
import { Home } from './components/home/home';

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
    path: '',
    redirectTo: "home",
    pathMatch: 'full'
  },
  {
    path: "home",
    component: Home,
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
    path: "trackers",
    component: Trackers
  },
  {
    path: "trackers/:id",
    component: Tracker
  },
  {
    path: "weather",
    component: Weathers
  }, {
    path: "weather/:id",
    component: Weather
  },
];
