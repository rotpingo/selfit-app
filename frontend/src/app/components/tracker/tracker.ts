import { Component, ElementRef, inject, signal, viewChild } from '@angular/core';
import { FabMenu } from '../shared/fab-menu/fab-menu';
import { IFabOption } from '../../models/types';
import { Backdrop } from '../shared/backdrop/backdrop';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { TrackerService } from '../../services/tracker-service';
import { ITracker } from '../../models/types';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-tracker',
  imports: [FabMenu, Backdrop, ReactiveFormsModule],
  templateUrl: './tracker.html',
  styleUrl: './tracker.css'
})
export class Tracker {

  trackerService = inject(TrackerService)
  trackers = this.trackerService.trackers
  isFormOpen = signal(false);
  form = viewChild.required<ElementRef<HTMLFormElement>>('form');

  createForm = new FormGroup({
    title: new FormControl('', {
      validators: [Validators.maxLength(20), Validators.minLength(3), Validators.required],
    }),
    notes: new FormControl('', {
      validators: [Validators.minLength(1), Validators.required]
    }),
    startDate: new FormControl('', {
      validators: [Validators.required]
    }),
  });


  menuOptions: IFabOption[] = [
    {
      label: 'create tracker',
      icon: 'icons/add-document.png',
      action: () => {
        this.isFormOpen.set(true);
        this.form().nativeElement.style.display = "flex";
      }
    },
  ];

  onCreateTracker() {
    if (this.createForm.valid) {
      const newTracker: ITracker = {
        title: this.createForm.value.title!,
        notes: this.createForm.value.notes!,
        startDate: new Date(this.createForm.value.startDate!).toISOString(),
      }
      this.trackerService.createTracker(newTracker).subscribe({
        next: () => {
          this.onCloseForm();
        },
        error: (err: HttpErrorResponse) => console.error(err),
      });
    }
  }

  onCloseForm() {
    this.isFormOpen.set(false);
    this.createForm.reset();
    this.form().nativeElement.style.display = "none";
  }

}
