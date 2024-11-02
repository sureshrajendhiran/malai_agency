import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CommonApiService {
  baseUrl = environment.baseUrl;

  private extractData(res: any) {
    const body = res;
    return body || {};
  }
  constructor(private http: HttpClient) { }


  login(obj: any): Observable<any> {
    const requestUrl = this.baseUrl + 'login/';
    return this.http.post(requestUrl, obj).pipe(map(this.extractData));
  }

  // Common row Create
  createRowCommon(body: any, tableName: string): Observable<any> {
    const requestUrl = this.baseUrl + 'create_record/' + tableName + '/';
    return this.http.post(requestUrl, body).pipe(map(this.extractData));
  }

  // Common updates
  updateCommon(body: any, tableName: string): Observable<any> {
    const requestUrl = this.baseUrl + '/update/' + tableName + '/';
    return this.http.put(requestUrl, body).pipe(map(this.extractData));
  }

  // Delete field
  deleteRow(id: string,tableName: string): Observable<any> {
    const requestUrl = this.baseUrl + '/delete/' + tableName + "/" + id + '/';
    return this.http.delete(requestUrl).pipe(map(this.extractData));
  }

  getMasterData(body: { type: string; }): Observable<any> {
    const requestUrl = this.baseUrl + 'get_master/' + body.type + '/';
    return this.http.post(requestUrl, body).pipe(map(this.extractData));
  }


  getSearchOption(obj: any): Observable<any> {
    const params = new HttpParams()
      .set('limit', !!obj.limit ? obj.limit : 10)
      .set('q', !!obj.q ? obj.q : '')
    const requestUrl = this.baseUrl + 'search_option/' + obj.type + "/";
    return this.http.get(requestUrl, { params }).pipe(map(this.extractData));
  }

  UpdateQI(body: { type: string; operation: string; }): Observable<any> {
    const requestUrl = this.baseUrl + 'update_qi/' + body.type + "/" + body.operation + "/";
    return this.http.put(requestUrl, body).pipe(map(this.extractData));
  }

  getFilterCount(type: any,filter:any): Observable<any> {
    const params = new HttpParams()
    const requestUrl = this.baseUrl + 'filter_countr/' + type + "/"+filter+"/";
    return this.http.get(requestUrl, { params }).pipe(map(this.extractData));
  }

  getDataQI(body: { type: string; }): Observable<any> {
    const requestUrl = this.baseUrl + 'get_qi/' + body.type;
    return this.http.post(requestUrl, body).pipe(map(this.extractData));
  }

  getQIInfoById(type: any, operation: any, id: any): Observable<any> {
    const params = new HttpParams()
    const requestUrl = this.baseUrl + 'qi/' + type + "/" + operation + "/" + id;
    return this.http.get(requestUrl, { params }).pipe(map(this.extractData));
  }
  savepdfFile(name: string, data: any) {
    console.log("ss")
    const linkSource = `data:application/pdf;base64,${data}`;
    const downloadLink = document.createElement('a');
    const fileName = name + "_" + (new Date().getTime().toString()) + '.pdf';
    downloadLink.href = linkSource;
    downloadLink.download = fileName;
    downloadLink.click();
  }
}