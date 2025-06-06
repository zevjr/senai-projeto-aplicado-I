import { Component, EventEmitter, Output } from '@angular/core';
import { Subject } from 'rxjs';
import { debounceTime, distinctUntilChanged } from 'rxjs/operators';

@Component({
  selector: 'app-search-bar',
  templateUrl: './search-bar.component.html',
  styleUrls: ['./search-bar.component.scss']
})
export class SearchBarComponent {
  @Output() searchChanged = new EventEmitter<string>();
  private searchTerms = new Subject<string>();

  constructor() {
    this.searchTerms.pipe(
      debounceTime(300),
      distinctUntilChanged()
    ).subscribe(term => {
      this.searchChanged.emit(term);
    });
  }

  search(term: string): void {
    this.searchTerms.next(term);
  }

}
