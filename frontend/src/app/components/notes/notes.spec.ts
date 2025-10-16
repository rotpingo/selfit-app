import { ComponentFixture, TestBed } from "@angular/core/testing";
import { Notes } from "./notes"
import { ReactiveFormsModule } from "@angular/forms";
import { NoteService } from "../../services/note-service";
import { provideZonelessChangeDetection } from "@angular/core";

class MockNoteService {
  notes = [];
}

describe('Notes', () => {
  let component: Notes;
  let fixture: ComponentFixture<Notes>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, Notes],
      providers: [
        { provide: NoteService, useClass: MockNoteService },
        provideZonelessChangeDetection()
      ],
    });

    fixture = TestBed.createComponent(Notes);
    component = fixture.componentInstance;

    component.form = jasmine.createSpy().and.returnValue({
      nativeElement: { style: {} }
    }) as any;
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  })

})
