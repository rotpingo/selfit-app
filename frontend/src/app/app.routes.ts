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
import { User } from './components/user/user';
import { AuthGuard } from './guards/AuthGuard';

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
    path: "home",
    component: Home,
    canActivate: [AuthGuard]
  },
  {
    path: "user",
    component: User,
    canActivate: [AuthGuard]
  },
  {
    path: "notes",
    component: Notes,
    canActivate: [AuthGuard]
  },
  {
    path: "notes/:id",
    component: Note,
    canActivate: [AuthGuard]
  },
  {
    path: "tasks",
    component: Tasks,
    canActivate: [AuthGuard]
  },
  {
    path: "tasks/:id",
    component: Task,
    canActivate: [AuthGuard]
  },
  {
    path: "trackers",
    component: Trackers,
    canActivate: [AuthGuard]
  },
  {
    path: "trackers/:id",
    component: Tracker,
    canActivate: [AuthGuard]
  },
  {
    path: "weather",
    component: Weathers,
    canActivate: [AuthGuard]
  }, {
    path: "weather/:id",
    component: Weather,
    canActivate: [AuthGuard]
  },
];
