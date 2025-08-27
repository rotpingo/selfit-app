import { Component, computed, ElementRef, inject, signal, viewChild } from '@angular/core';
import { TrackerService } from '../../../services/tracker-service';
import { ActivatedRoute, Router } from '@angular/router';
import { DatePipe } from '@angular/common';
import { FabMenu } from '../../shared/fab-menu/fab-menu';
import { IFabOption, ITracker } from '../../../models/types';
import { HttpErrorResponse } from '@angular/common/http';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';

@Component({
  selector: 'app-tracker',
  imports: [DatePipe, FabMenu, ReactiveFormsModule],
  templateUrl: './tracker.html',
  styleUrl: './tracker.css'
})
export class Tracker {

  trackerService = inject(TrackerService);
  activeRoute = inject(ActivatedRoute);
  route = inject(Router);

  trackerID = parseInt(this.activeRoute.snapshot.paramMap.get('id')!);
  tracker = this.trackerService.getTrackerByID(this.trackerID);

  isFormOpen = signal(false);
  form = viewChild.required<ElementRef<HTMLFormElement>>('form');
  trackerForm = viewChild.required<ElementRef<HTMLFormElement>>('trackerForm');

  daysSince = computed(() => {
    const startDate = new Date(this.tracker()?.startDate!);
    const now = new Date();
    const diff = now.getTime() - startDate.getTime();
    return Math.floor(diff / (1000 * 60 * 60 * 24));
  })

  editForm = new FormGroup({
    title: new FormControl('', {
      validators: [Validators.maxLength(20), Validators.minLength(3), Validators.required],
    }),
    notes: new FormControl('', {
      validators: [Validators.minLength(1), Validators.required]
    }),
  });


  menuOptions: IFabOption[] = [
    {
      label: 'delete tracker',
      icon: 'icons/delete-document.png',
      action: () => {
        this.trackerService.deleteTracker(this.trackerID).subscribe({
          next: () => {
            this.route.navigate(["trackers/"]);
            this.trackerService.refresh()
          },
          error: (err: HttpErrorResponse) => console.error(err),
        });
      }
    },
    {
      label: 'edit note',
      icon: 'icons/file-edit.png',
      action: () => {
        this.editForm.patchValue({
          title: this.tracker()?.title,
          notes: this.tracker()?.notes,
        });
        console.log(this.tracker())
        this.isFormOpen.set(true);
        this.form().nativeElement.style.display = "flex";
      }
    },
  ];

  async onEditTracker() {
    if (this.editForm.valid) {
      const newTracker: ITracker = {
        id: this.tracker()?.id,
        title: this.editForm.value.title!,
        notes: this.editForm.value.notes!,
      };
      this.trackerService.editTracker(newTracker).subscribe({
        next: () => {
          this.trackerService.refresh();
          this.onCloseForm();
        },
        error: (err: HttpErrorResponse) => console.error(err)
      });
    }
  }

  formatDate(date: string | Date | undefined | null): string | null {
    if (!date) return null;
    const d = new Date(date);
    return d.toISOString().split('T')[0];
  }

  onCloseForm() {
    this.isFormOpen.set(false);
    this.editForm.reset();
    this.form().nativeElement.style.display = "none";
  }

}
