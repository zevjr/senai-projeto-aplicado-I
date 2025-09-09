import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-mobile-public',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './mobile-public.layout.html',
  styleUrls: ['./mobile-public.layout.scss']
})
export class MobilePublicLayout { }
