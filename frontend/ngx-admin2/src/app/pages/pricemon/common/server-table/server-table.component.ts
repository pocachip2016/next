import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ServerDataSource } from 'ng2-smart-table';
import { GoDataSource } from './go.data-source';

@Component({
  selector: 'ngx-server-table',
  template:  `
      <ng2-smart-table [settings]="settings" [source]="source"></ng2-smart-table>
  `,
})

export class ServerTableComponent {
  conf ={ 
    endPoint: "http://127.0.0.1:5555/api/v1/content",
    totalKey: "total",
    dataKey: "data",
    pagerPageKey: "page",
    pagerLimitKey: "limit",
    /*
    filterFieldKey: 'title' 
    filterFieldKey
    sortFieldKey: "id",
    sortDirKey: "order",
    */
  }

  //settings = {
    //columns: {
      //id: {
        //title: 'ID',
      //},
      //albumId: {
        //title: 'Album',
      //},
      //title: {
        //title: 'Title',
      //},
      //url: {
        //title: 'Url',
      //},
    //},
  //};
  settings = {
    add: {
      addButtonContent: '<i class="nb-plus"></i>',
      createButtonContent: '<i class="nb-checkmark"></i>',
      cancelButtonContent: '<i class="nb-close"></i>',
    },
    edit: {
      editButtonContent: '<i class="nb-edit"></i>',
      saveButtonContent: '<i class="nb-checkmark"></i>',
      cancelButtonContent: '<i class="nb-close"></i>',
    },
    delete: {
      deleteButtonContent: '<i class="nb-trash"></i>',
      confirmDelete: true,
    },
    columns: {
      id: {
        title: 'ID',
        type: 'string',
        filter: false,
        editable: false,
      },
      title: {
        title: '콘텐츠명',
        type: 'string',
        filter: false,
      },
    },
    pager: {
      display : true,
      perPage: 20,
    },
  };
  source: GoDataSource;

  constructor(http: HttpClient) { 
  //  this.source = new ServerDataSource(http, { endPoint:'https://jsonplaceholder.typicode.com/photos' })
    //this.source = new ServerDataSource(http, { endPoint:'http://127.0.0.1:5555/api/v1/content?limit=100&offset=0' })
    //this.source = new GoDataSource(http, this.conf);
    this.source = new GoDataSource(http, { endPoint:'http://127.0.0.1:5555/api/v1/content?limit=1000' })
  }

}
