import { Component, inject } from '@angular/core';
import { TrackerService } from '../../../services/tracker-service';
import { ActivatedRoute, Router } from '@angular/router';
import { DatePipe } from '@angular/common';
import { FabMenu } from '../../shared/fab-menu/fab-menu';
import { IFabOption } from '../../../models/types';

@Component({
  selector: 'app-tracker',
  imports: [DatePipe, FabMenu],
  templateUrl: './tracker.html',
  styleUrl: './tracker.css'
})
export class Tracker {

  noteService = inject(TrackerService);
  activeRoute = inject(ActivatedRoute);
  route = inject(Router);

  trackerID = parseInt(this.activeRoute.snapshot.paramMap.get('id')!);
  tracker = this.noteService.getTrackerByID(this.trackerID);

  menuOptions: IFabOption[] = [
    {
      label: 'delete tracker',
      icon: 'icons/delete-document.png',
      action: () => {
      }
    },
    {
      label: 'edit note',
      icon: 'icons/file-edit.png',
      action: () => {
      }
    },
  ];

}
