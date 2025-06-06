import { Component } from '@angular/core';
import {RouterOutlet} from '@angular/router';

@Component({
  selector: 'app-public',
  templateUrl: './public.component.html',
  imports: [
    RouterOutlet
  ],
  styleUrls: ['./public.component.scss']
})
export class PublicComponent {
  constructor() { }
}
