import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ToastrService } from "ngx-toastr";

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './app.html',
  styleUrl: './app.scss',

})
export class App {
  protected title = 'senai-projeto-aplicado-1';
}
