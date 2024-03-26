import { NgModule } from '@angular/core';
import { NbCardModule, NbIconModule, NbInputModule, NbTreeGridModule } from '@nebular/theme';
import { Ng2SmartTableModule } from 'ng2-smart-table';
import { ThemeModule } from '../../@theme/theme.module';
import { PricemonRoutingModule, routedComponents } from './pricemon-routing.module';
import { PricemonService } from './pricemon.service';
import { ContentListComponent } from './content-panel/content-list/content-list.component';
import { ServerTableComponent } from './common/server-table/server-table.component';
@NgModule({
  imports: [
    NbCardModule,
    NbTreeGridModule,
    NbIconModule,
    NbInputModule,
    ThemeModule,
    Ng2SmartTableModule,
    PricemonRoutingModule,
  ],
  declarations: [
    ...routedComponents,
    ContentListComponent,
    ServerTableComponent,
  ],
  providers: [
    PricemonService,
  ],
})
export class PricemonModule { }