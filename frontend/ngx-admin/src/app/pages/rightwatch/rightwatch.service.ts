import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { delay, map } from 'rxjs/operators';

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

export interface ICpWebhardEntry {
  website: number;
  count: number;
}

export interface ICpDashboardItem {
  cp_id: number;
  cp_name: string;
  detected: number;
  by_webhard: ICpWebhardEntry[];
}

@Injectable()
export class RightwatchService {

  private baseUrl = 'http://127.0.0.1:5555/api/v1';

  constructor(private http: HttpClient) {}

  getCpDashboard(): Observable<any> {
    return this.http.get(`${this.baseUrl}/cp/dashboard`);
  }

  getData(url: string): Observable<IApiData>{
    //console.log("getData.....")
    console.log(url)
    return this.http.get<IApiData>(url);
  }
/*
  load(page: number, pageSize: number): Observable<NewsPost[]> {
    const startIndex = ((page - 1) % TOTAL_PAGES) * pageSize;

    return this.http
      .get<NewsPost[]>('assets/data/news.json')
      .pipe(
        map(news => news.splice(startIndex, pageSize)),
        delay(1500),
      );
  }
  */
}
