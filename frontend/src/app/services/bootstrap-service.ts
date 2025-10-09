import { inject, Injectable } from "@angular/core";
import { NoteService } from "./note-service";
import { TaskService } from "./task-service";
import { TrackerService } from "./tracker-service";
import { WeatherService } from "./weather-service";
import { forkJoin, map, Observable } from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class BootstrapService {

  noteService = inject(NoteService);
  taskService = inject(TaskService);
  trackerService = inject(TrackerService);
  weatherService = inject(WeatherService);

  loadAllData(): Observable<void> {
    return forkJoin([
      this.noteService.loadNotes(),
      this.taskService.loadTasks(),
      this.trackerService.loadTrackers(),
      this.weatherService.loadWeatherCities(),
    ]).pipe(map(() => void 0));
  }
}
