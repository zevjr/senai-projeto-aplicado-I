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

  getRegister(uuid: string): Observable<Register> {
    return this.http.get<Register>(`${this.apiUrl}/${uuid}`);
  }

  createRegister(register: Partial<Register>): Observable<Register> {
    return this.http.post<Register>(this.apiUrl, register);
  }

  updateRegister(uid: string, register: Partial<Register>): Observable<Register> {
    return this.http.put<Register>(`${this.apiUrl}/${uid}`, register);
  }

  deleteRegister(uid: string) {
    return this.http.delete(`${this.apiUrl}/${uid}`);
  }
}