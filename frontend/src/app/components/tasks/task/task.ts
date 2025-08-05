import { Component, ElementRef, inject, signal, viewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { TaskService } from '../../../services/task-service';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { IFabOption, ITask } from '../../../models/types';
import { HttpErrorResponse } from '@angular/common/http';
import { FabMenu } from '../../shared/fab-menu/fab-menu';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'app-task',
  imports: [ReactiveFormsModule, FabMenu, DatePipe],
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
    dueDate: new FormControl(new Date(Date.now()), {
      validators: [Validators.required]
    })
  });


  menuOptions: IFabOption[] = [
    {
      label: 'delete task',
      icon: '',
      action: () => {
        this.taskService.deleteTask(this.taskID).subscribe({
          next: () => {
            this.taskService.refresh();
            this.route.navigate(["notes/"]);
          },
          error: (err: HttpErrorResponse) => alert(err.message),
        });
      }
    },
    {
      label: 'edit task',
      icon: '',
      action: () => {
        this.editForm.patchValue({
          title: this.task()?.title,
          content: this.task()?.content
        });
        this.isFormOpen.set(true);
        this.form().nativeElement.style.display = "flex";
      }
    },
  ];

  async onEditTask() {
    if (this.editForm.valid) {
      const newTask: ITask = {
        id: this.task()?.id,
        title: this.editForm.value.title!,
        content: this.editForm.value.content!,
        isRepeat: this.editForm.value.isRepeat!,
        interval: this.editForm.value.interval!,
        dueDate: this.editForm.value.dueDate!,
        createdAt: this.task()?.createdAt,
        userId: 0
      };

      console.log(newTask);

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

  onCloseForm() {
    this.isFormOpen.set(false);
    this.editForm.reset();
    this.form().nativeElement.style.display = "none";
  }

}
