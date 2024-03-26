import { RouterModule, Routes } from '@angular/router';
import { NgModule } from '@angular/core';
import { PricemonComponent } from './pricemon.component';
import { ContentPanelComponent }  from './content-panel/content-panel.component';
import { PricePanelComponent } from './price-panel/price-panel.component';
import { ProductPanelComponent } from './product-panel/product-panel.component';

const routes: Routes = [{
  path: '',
  component: PricemonComponent,
  children: [
    {
      path: 'content-panel',
      component: ContentPanelComponent,
    },
    {
      path: 'price-panel',
      component: PricePanelComponent,
    },
    {
      path: 'product-panel',
      component: ProductPanelComponent,
    },
  ],
}];


@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class PricemonRoutingModule { }

export const routedComponents = [
  PricemonComponent, 
  ContentPanelComponent,
  PricePanelComponent,
  ProductPanelComponent,
];

