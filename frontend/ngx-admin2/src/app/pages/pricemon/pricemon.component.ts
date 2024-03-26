import { Component } from '@angular/core';
import { NbThemeService } from '@nebular/theme';

@Component({
  selector: 'ngx-pricemon',
  template: `<router-outlet></router-outlet>`,
})
export class PricemonComponent {
  constructor(private themeService: NbThemeService) {
    this.themeService.changeTheme('cosmic');
  }
}
