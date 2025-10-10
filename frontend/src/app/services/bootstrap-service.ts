import { effect, inject, Injectable } from "@angular/core";
import { NoteService } from "./note-service";
import { TaskService } from "./task-service";
import { TrackerService } from "./tracker-service";
import { WeatherService } from "./weather-service";
import { forkJoin, map, Observable } from "rxjs";
import { UserService } from "./user-service";
import { AuthService } from "./auth-service";

@Injectable({
  providedIn: 'root'
})
export class BootstrapService {

  authService = inject(AuthService);
  noteService = inject(NoteService);
  taskService = inject(TaskService);
  trackerService = inject(TrackerService);
  weatherService = inject(WeatherService);
  userService = inject(UserService);

  constructor() {
    effect(() => {
      console.log(this.authService.isLoggedIn())
      if (this.authService.isLoggedIn()) {
        this.loadAllData().subscribe({
          next: () => console.log("Data loaded"),
          error: (err) => console.error(err)
        });
      } else {
        this.clearAllData();
      }
    })
  }

  loadAllData(): Observable<void> {
    return forkJoin([
      this.noteService.loadNotes(),
      this.taskService.loadTasks(),
      this.trackerService.loadTrackers(),
      this.weatherService.loadWeatherCities(),
      // this.userService.loadUser()
    ]).pipe(map(() => void 0));
  }

  clearAllData(): void {
    this.noteService.clearNotes();
    this.taskService.clearTasks();
    this.trackerService.clearTrackers();
    this.weatherService.clearCitiesWeather();
    // this.userService.clearUser();
  }

}
