import { Component, ElementRef, inject, signal, viewChild } from '@angular/core';
import { TaskService } from '../../services/task-service';
import { IFabOption, ITask } from '../../models/types';
import { RouterLink } from '@angular/router';
import { FabMenu } from '../shared/fab-menu/fab-menu';
import { Backdrop } from '../shared/backdrop/backdrop';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-tasks',
  imports: [RouterLink, FabMenu, Backdrop, ReactiveFormsModule],
  templateUrl: './tasks.html',
  styleUrl: './tasks.css'
})
export class Tasks {

  taskService = inject(TaskService)
  tasks = this.taskService.tasks

  isFormOpen = signal(false);
  createForm = new FormGroup({
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

  form = viewChild.required<ElementRef<HTMLFormElement>>('form');

  menuOptions: IFabOption[] = [
    {
      label: 'create task',
      icon: 'icons/add-document.png',
      action: () => {
        this.isFormOpen.set(true);
        this.form().nativeElement.style.display = "flex";
      }
    },
  ]

  async onCreateTask() {

    if (this.createForm.valid) {
      const newTask: ITask = {
        title: this.createForm.value.title!,
        content: this.createForm.value.content!,
        isRepeat: this.createForm.value.isRepeat!,
        interval: this.createForm.value.interval!,
        dueDate: new Date(this.createForm.value.dueDate!).toISOString(),
      };
      this.taskService.createTask(newTask).subscribe({
        next: () => {
          this.onCloseForm();
          this.taskService.refresh();
        },
        error: (err: HttpErrorResponse) => console.error('error:', err)
      });
    } else {
      console.log("invalid form")
    }

  }

  onCloseForm() {
    this.isFormOpen.set(false);
    this.createForm.reset();
    this.form().nativeElement.style.display = "none";
  }


}
