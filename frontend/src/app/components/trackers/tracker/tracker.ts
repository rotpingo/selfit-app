import { Component, inject } from '@angular/core';
import { TrackerService } from '../../../services/tracker-service';
import { ActivatedRoute, Router } from '@angular/router';
import { DatePipe } from '@angular/common';
import { FabMenu } from '../../shared/fab-menu/fab-menu';
import { IFabOption } from '../../../models/types';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-tracker',
  imports: [DatePipe, FabMenu],
  templateUrl: './tracker.html',
  styleUrl: './tracker.css'
})
export class Tracker {

  trackerService = inject(TrackerService);
  activeRoute = inject(ActivatedRoute);
  route = inject(Router);

  trackerID = parseInt(this.activeRoute.snapshot.paramMap.get('id')!);
  tracker = this.trackerService.getTrackerByID(this.trackerID);

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
        console.log(this.tracker())
      }
    },
  ];

}
