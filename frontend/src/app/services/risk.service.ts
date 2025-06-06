import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Risk } from '../models/risk.model';

@Injectable({
  providedIn: 'root'
})
export class RiskService {
  private apiUrl = '/api/registers';

  constructor(private http: HttpClient) { }

  getRisks(): Observable<Risk[]> {
    return this.http.get<Risk[]>(this.apiUrl);
  }

  getRisk(id: number): Observable<Risk> {
    return this.http.get<Risk>(`${this.apiUrl}/${id}`);
  }

  createRisk(risk: Partial<Risk>): Observable<Risk> {
    return this.http.post<Risk>(this.apiUrl, risk);
  }

  updateRisk(id: number, risk: Partial<Risk>): Observable<Risk> {
    return this.http.put<Risk>(`${this.apiUrl}/${id}`, risk);
  }

  deleteRisk(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`);
  }
}