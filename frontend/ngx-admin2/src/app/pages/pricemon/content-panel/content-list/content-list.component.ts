import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ServerDataSource, LocalDataSource} from 'ng2-smart-table';
import { GoDataSource } from '../../common/server-table/go.data-source';
import { HttpClient } from '@angular/common/http';
import { PricemonService } from '../../pricemon.service';
import { IApiData } from '../../pricemon.service';

@Component({
  selector: 'ngx-content-list',
  templateUrl: './content-list.component.html',
  styleUrls: ['./content-list.component.scss']
})
export class ContentListComponent  implements OnInit {
  @Input()
    get contentID(): string { return this._contentID;}
    set contentID(c_id: string){
      this._contentID = c_id;
    }
  
  private _contentID  = '';

  @Output() selectedEvent = new EventEmitter<string>();

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
  updateflag: boolean;
  dataService: PricemonService;

  conf ={ 
    //endPoint: "http://127.0.0.1:5555/api/v1/content?limit=1000&offset=0",
    endPoint: "http://127.0.0.1:5555/api/v1/content?",
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

  constructor(private http: HttpClient){
    //this.source = new LocalDataSource();
    //this.source = new ServerDataSource(http, this.conf);
    this.source = new GoDataSource(http, this.conf);
    this.dataService = new PricemonService(http);
    this.updateflag = false;
  }

  ngOnInit(): void {
    console.log("ngOnInit called!!!")
    console.log(this);

    this.source.onChanged().subscribe((change) => {
    //if (change.action === 'sort') {
     // this.sortingChange(change.sort);
    //}
    //else 
    if (change.action === 'page') {
      this.pageChange(change.paging.page);
    }
    }
    );
  }

  pageChange(pagenumber : string): void {
//    console.log("pageChange");
    let offstr:string = (this.settings.pager.perPage * (Number(pagenumber)-1)).toString() 

    let targetUrl: string ="http://127.0.0.1:5555/api/v1/content?limit=20&offset="+offstr;
    this.conf.endPoint = targetUrl;
    this.dataService.setTargetUrl(targetUrl);
    this.dataService.getAll().subscribe({
      next: (val) => {
        let _data: any[];
        _data = val.data;
        this.source.empty();
    //    this.source.load(_data);
     //   this.source.refresh();
        console.log(">>>>>>>>>>>>>>>>>")
        console.log(_data);
      }
    });

    //this.apiService.getAll().subscribe(val=>
      //{
        //console.log(val);
        //this.source.empty();
        //this.source.load(val as []);
        //this.source.refresh();
        //console.log(this.source);
      //});
  }

  getData(query:string): void {
    let urlin = '';
    if(query){
      urlin = "http://127.0.0.1:5555/api/v1/content?where=title%3A"+query;
    } else {
      urlin = "http://127.0.0.1:5555/api/v1/content";
    }
    this.conf.endPoint = urlin;
    this.source = new ServerDataSource(this.http, this.conf)
  }

  reload(query:string){
    this.getData(query);
  }

  setEndPoint(query: string){
    if(query){
      this.conf.endPoint = "http://127.0.0.1:5555/api/v1/content?where=title%3A"+query;
    } else {
      this.conf.endPoint = "http://127.0.0.1:5555/api/v1/content";
    }
    this.reload(query);
  }

  onSearch(query: string = '') {
    this.source.setFilter([
      // fields we want to include in the search
      {
        field: 'id',
        search: query
      },
      {
        field: 'title',
        search: query
      },
    ], false); 
    
//    this.setEndPoint(query);
//    this.getData(query);

//    this.source.reset(true);  // reset your old filtered data 
//    this.source.setPage(1, false); // set page to 1 to start from beginning 
//    console.log(this);
//    this.source.setFilter([
//      // fields we want to include in the search
//      {
//        field: 'title',
//        search: query,
//      },
//    ], false); 
//    this.source.refresh();
//    
//    console.log(this);
//    console.log(query);
  // second parameter specifying whether to perform 'AND' or 'OR' search 
  // (meaning all columns should contain search query or at least one)
  // 'AND' by default, so changing to 'OR' by setting false here
  }

  onUserRowSelect(event): void {
    console.log(event.data.id);
    console.log("onUserRowSelect");
    //console.log(event);
    //this._selectedID= event.data.id;
    this._contentID= event.data.id;
    this.selectedEvent.emit(this._contentID)
  }

  onDeleteConfirm(event): void {
    if (window.confirm('Are you sure you want to delete?')) {
      event.confirm.resolve();
    } else {
      event.confirm.reject();
    }
  }
 
  FilterData():void {
    this.source.reset(true);  // reset your old filtered data 
//    this.source.setPage(1, false); // set page to 1 to start from beginning 
/*
    let filterArr = this.getFilterArray(); // add a new filter data, but be careful to not sent any empty data, as it throws an exception 
    if (filterArr.length)
      this.source.setFilter(filterArr, false, false);
*/
    this.source.refresh(); // this will call the server with new filter and paginatio data
  }
/*
  getFilterArray() {  // setup new filter 
    let filterArray = [];
    if (this.filter.id)
      filterArray.push({ field: 'id', search: this.filter.id });
    if (this.filter.name)
      filterArray.push({ field: 'name', search: this.filter.name});

    return filterArray;  
  }
  */
  onCustomAction(event) {  // custom buttons code 
    switch (event.action) {
      case 'view-something':
      // put your code here 
      break;
      default:
      console.log('Not Implemented Action');
      break;
    }
  }
}
