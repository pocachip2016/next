import { HttpClient } from '@angular/common/http';
import { ServerDataSource } from 'ng2-smart-table';
import { map } from 'rxjs/operators';
import { ServerSourceConf } from 'ng2-smart-table/lib/lib/data-source/server/server-source.conf';

export interface IApiData {
  code : number;
  data : any;
  limit : number;
  offset: number;
  total: number;
}

export class GoDataSource extends ServerDataSource {
    
  constructor(http: HttpClient, conf?: ServerSourceConf | {}){
    super(http, conf);

    if (!this.conf.endPoint) {
      throw new Error('At least endPoint must be specified as a configuration of the server data source.');
    }
  }

  getElements(): Promise<any> {
    console.log("here!!!!!!!!!!!!!!!!!!!!!!!")
    let url = 'http://127.0.0.1:5555/api/v1/content?'

    if (this.sortConf) {
      this.sortConf.forEach((fieldConf) => {
        url += `_sort=${fieldConf.field}&_order=${fieldConf.direction.toUpperCase()}&`;
      });
    }

    if (this.pagingConf && this.pagingConf['page'] && this.pagingConf['perPage']) {
      //url += `_page=${this.pagingConf['page']}&_limit=${this.pagingConf['perPage']}&`;
      let offstr:string = (this.pagingConf['perPage'] * (Number(this.pagingConf['page'])-1)).toString() 
      url += `offset=${offstr}&limit=${this.pagingConf['perPage']}`;
    }



    if (this.filterConf.filters) {
      this.filterConf.filters.forEach((fieldConf) => {
        if (fieldConf['search']) {
          url += `${fieldConf['field']}_like=${fieldConf['search']}&`;
        }
      });
    }

    return this.http.get<IApiData>(url, { observe: 'response' })
      .pipe(
        map(res => {
          this.lastRequestCount = +res.body.total;
          return res.body.data;
        })
      ).toPromise();
  }
}
