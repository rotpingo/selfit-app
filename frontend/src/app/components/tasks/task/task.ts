import { Component, ElementRef, inject, signal, viewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { TaskService } from '../../../services/task-service';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { IFabOption, ITask } from '../../../models/types';
import { HttpErrorResponse } from '@angular/common/http';
import { FabMenu } from '../../shared/fab-menu/fab-menu';
import { DatePipe } from '@angular/common';
import { Backdrop } from '../../shared/backdrop/backdrop';

@Component({
  selector: 'app-task',
  imports: [ReactiveFormsModule, FabMenu, DatePipe, Backdrop],
  templateUrl: './task.html',
  styleUrl: './task.css'
})
export class Task {

  taskService = inject(TaskService);
  activeRoute = inject(ActivatedRoute);
  route = inject(Router);

  taskID = parseInt(this.activeRoute.snapshot.paramMap.get('id')!);
  task = this.taskService.getTaskByID(this.taskID);

  isFormOpen = signal(false);
  form = viewChild.required<ElementRef<HTMLFormElement>>('form');
  taskForm = viewChild.required<ElementRef<HTMLFormElement>>('taskForm');

  endForm = new FormGroup({
    notes: new FormControl(this.task()?.notes)
  })

  editForm = new FormGroup({
    title: new FormControl('', {
      validators: [Validators.maxLength(20), Validators.minLength(3), Validators.required],
    }),
    content: new FormControl('', {
      validators: [Validators.minLength(1), Validators.required]
    }),
    isRepeat: new FormControl(false, {
      validators: [Validators.required]
    }),
    interval: new FormControl(0),
    dueDate: new FormControl('', {
      validators: [Validators.required]
    })
  });


  menuOptions: IFabOption[] = [
    {
      label: 'delete task',
      icon: 'icons/delete-document.png',
      action: () => {
        this.taskService.deleteTask(this.taskID).subscribe({
          next: () => {
            this.taskService.refresh();
            this.route.navigate(["tasks/"]);
          },
          error: (err: HttpErrorResponse) => alert(err.message),
        });
      }
    },
    {
      label: 'edit task',
      icon: 'icons/file-edit.png',
      action: () => {
        this.editForm.patchValue({
          title: this.task()?.title,
          content: this.task()?.content,
          isRepeat: this.task()?.isRepeat,
          interval: this.task()?.interval,
          dueDate: this.formatDate(this.task()?.dueDate)
        });
        this.isFormOpen.set(true);
        this.form().nativeElement.style.display = "flex";
      }
    },
  ];

  formatDate(date: string | Date | undefined | null): string | null {
    if (!date) return null;
    const d = new Date(date);
    return d.toISOString().split('T')[0];
  }

  async onEditTask() {
    if (this.editForm.valid) {
      const newTask: ITask = {
        id: this.task()?.id,
        title: this.editForm.value.title!,
        content: this.editForm.value.content!,
        isRepeat: this.editForm.value.isRepeat!,
        interval: this.editForm.value.interval!,
        dueDate: new Date(this.editForm.value.dueDate!).toISOString(),
        userId: 0
      };
      this.taskService.editTask(newTask).subscribe({
        next: () => {
          this.taskService.refresh();
          this.onCloseForm();
        },
        error: (err: HttpErrorResponse) => console.error('error:', err)
      });
    } else {
      console.log("invalid form")
    }
  };

  onAbortTask() {
    const payload = {
      notes: this.endForm.value.notes ?? ''
    }

    this.taskService.abortTask(this.taskID, payload).subscribe({
      next: () => {
        this.taskService.refresh();
        this.route.navigate(["tasks/"]);
      },
      error: (err: HttpErrorResponse) => console.error(err),
    });
  }

  onCompleteTask() {

    const payload = {
      notes: this.endForm.value.notes ?? ''
    }

    this.taskService.completeTask(this.taskID, payload).subscribe({
      next: () => {
        this.taskService.refresh();
        this.route.navigate(["tasks/"]);
      },
      error: (err: HttpErrorResponse) => console.error(err),
    });
  }

  onCloseForm() {
    this.isFormOpen.set(false);
    this.editForm.reset();
    this.form().nativeElement.style.display = "none";
  }
}
