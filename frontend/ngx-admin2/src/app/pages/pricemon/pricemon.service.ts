import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { delay, filter, map, tap } from 'rxjs/operators';

export interface IApiData {
  code : number;
  data : any;
  limit : number;
  offset: number;
  total: number;
}

export interface IContentDetail {
  contents_id : string;
  first_time : string;
  id : number;
  post_id: string;
  post_idx: string;
  status : number;
  update_time : string;
}

@Injectable({
  providedIn: 'root'
})
export class PricemonService {
  private baseUrl: string;
  private targetUrl: string;
  private cUrl: string;
  constructor(private http: HttpClient) {}


  setTargetUrl(url: string){
    this.cUrl = url;
  }

  getAll(): Observable<IApiData>{
    return this.http.get<IApiData>(this.cUrl);
    //return this.http.get<IApiData[]>(this.cUrl);
  }
  get(id: any): Observable<IApiData> {
    return this.http.get<IApiData>(`${this.cUrl}/${id}`);
  }

  create(data: any): Observable<any> {
    return this.http.post(this.cUrl, data);
  }

  update(id: any, data: any): Observable<any> {
    return this.http.put(`${this.cUrl}/${id}`, data);
  }

  delete(id: any): Observable<any> {
    return this.http.delete(`${this.cUrl}/${id}`);
  }

  deleteAll(): Observable<any> {
    return this.http.delete(this.cUrl);
  }

  findByTitle(title: any): Observable<IApiData[]> {
    return this.http.get<IApiData[]>(`${this.cUrl}?title=${title}`);
  }

  getData(url: string): Observable<IApiData[]>{
    //console.log("getData.....")
    console.log(url)
    return this.http.get<IApiData[]>(this.cUrl).pipe(tap((data) => console.log("All: "
                    + JSON.stringify(data))));;
  }
}
