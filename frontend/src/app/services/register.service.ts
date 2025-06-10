import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import {Register} from "../models/register.model";

@Injectable({
  providedIn: 'root'
})
export class RegisterService {
  private apiUrl = '/api/registers';

  constructor(private http: HttpClient) { }

  getRegisters(): Observable<Register[]> {
    return this.http.get<Register[]>(this.apiUrl);
  }

  getRegister(id: number): Observable<Register> {
    return this.http.get<Register>(`${this.apiUrl}/${id}`);
  }

  createRegister(register: Partial<Register>): Observable<Register> {
    return this.http.post<Register>(this.apiUrl, register);
  }

}