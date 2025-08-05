import { HttpClient } from '@angular/common/http';
import { computed, inject, Injectable, Signal, signal } from '@angular/core';
import { ITask } from '../models/types';
import { Observable } from 'rxjs';
import { toSignal } from '@angular/core/rxjs-interop';

@Injectable({
  providedIn: 'root'
})
export class TaskService {

  private apiUrl = 'http://localhost:6969/api/tasks';
  http = inject(HttpClient)

  private readonly _tasks = signal<ITask[]>([]);
  readonly tasks = this._tasks.asReadonly()

  constructor() {
    this.loadTasks();
  }

  loadTasks(): void {
    this.http.get<ITask[]>(this.apiUrl).subscribe({
      next: (tasks) => {
        this._tasks.set(tasks);
      },
      error: (err) => console.error('Failed to load tasks', err)
    })
  }

  refresh(): void {
    this.loadTasks()
  }

  getTasks(): Signal<ITask[]> {
    const request$: Observable<ITask[]> = this.http.get<ITask[]>(this.apiUrl);
    return toSignal(request$, { initialValue: [] });
  }

  getTaskByID(taskID: number): Signal<ITask | undefined> {
    return computed(() => this._tasks().find((task) => task.id === taskID));
  }

  createTask(task: ITask): Observable<void> {
    return this.http.post<void>(this.apiUrl, task);
  }

  editTask(task: ITask): Observable<void> {
    const url = `${this.apiUrl}/${task.id}`;
    return this.http.put<void>(url, task)
  }

  deleteTask(taskID: number): Observable<void> {
    const url = `${this.apiUrl}/${taskID}`;
    return this.http.delete<void>(url)
  }
}
