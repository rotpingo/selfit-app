import { HttpClient } from '@angular/common/http';
import { inject, Injectable, Signal, signal } from '@angular/core';
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
        const loadedTasks = this.toCamelCaseTasks(tasks);
        this._tasks.set(loadedTasks)
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

  createTask(task: ITask): Observable<void> {
    return this.http.post<void>(this.apiUrl, task);
  }

  // transforms the json snake_case format data into camel Case
  toCamelCaseTask(task: any): ITask {
    return {
      id: task.id,
      parentId: task.parent_id,
      title: task.title,
      content: task.content,
      status: task.status,
      isRepeat: task.is_repeat,
      interval: task.interval,
      notes: task.notes,
      dueDate: task.due_date,
      execAt: task.exec_at,
      createdAt: task.created_at,
      updatedAt: task.updated_at,
      userId: task.user_id,
    };
  }

  // transforms the json snake_case format data into camel Case for the entire array
  toCamelCaseTasks(tasks: any[]): ITask[] {
    return tasks.map(this.toCamelCaseTask);
  }
}
