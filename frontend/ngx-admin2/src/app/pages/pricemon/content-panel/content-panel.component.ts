import { Component, OnDestroy, Directive, Input, ViewChild } from '@angular/core';
import { ContentListComponent } from './content-list/content-list.component';

@Component({
  selector: 'ngx-content-panel',
  templateUrl: './content-panel.component.html',
  styleUrls: ['./content-panel.component.scss']
})

export class ContentPanelComponent implements OnDestroy {

  @ViewChild(ContentListComponent) contents_list: any; 
//  @ViewChild(ContentDetailComponent) content_detail: any; 

  private alive = true;
  ID :string = '';

  constructor() {
   }

  changeContent(event: any){
    this.ID = event.id;
  }

  ngOnDestroy() {
    this.alive = false;
  }
}
