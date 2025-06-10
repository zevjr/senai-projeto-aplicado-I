import { Injectable } from '@angular/core';
import {Observable} from "rxjs";
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class FileService {

  private readonly imageUploadUrl = '/api/images';
  private readonly audioUploadUrl = '/api/audios';

  constructor(private http: HttpClient) {}

  uploadImage(file: File): Observable<{ uid: string }> {
    const formData = new FormData();
    formData.append('file', file);
    return this.http.post<{ uid: string }>(this.imageUploadUrl, formData);
  }

  uploadAudio(file: File): Observable<{ uid: string }> {
    const formData = new FormData();
    formData.append('file', file);
    return this.http.post<{ uid: string }>(this.audioUploadUrl, formData);
  }
}
