import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Risk } from '../models/risk.model';

@Injectable({
  providedIn: 'root'
})
export class RegistersService {
  private apiUrl = '/api/registers';

  constructor(private http: HttpClient) { }

  getRegisters(): Observable<Risk[]> {
    return this.http.get<Risk[]>(this.apiUrl);
  }

  getRegister(id: number): Observable<Risk> {
    return this.http.get<Risk>(`${this.apiUrl}/${id}`);
  }

  createRegister(risk: Partial<Risk>): Observable<Risk> {
    return this.http.post<Risk>(this.apiUrl, risk);
  }

}