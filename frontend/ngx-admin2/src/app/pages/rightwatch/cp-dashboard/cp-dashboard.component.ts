import { Component, OnInit } from '@angular/core';
import { RightwatchService, ICpDashboardItem } from '../rightwatch.service';

@Component({
  selector: 'ngx-cp-dashboard',
  templateUrl: './cp-dashboard.component.html',
  styleUrls: ['./cp-dashboard.component.scss'],
})
export class CpDashboardComponent implements OnInit {
  items: ICpDashboardItem[] = [];
  loading = true;
  error = '';

  constructor(private service: RightwatchService) {}

  ngOnInit() {
    this.service.getCpDashboard().subscribe({
      next: (res: any) => {
        this.items = res.data || [];
        this.loading = false;
      },
      error: (err) => {
        this.error = err.message || '데이터를 불러오지 못했습니다.';
        this.loading = false;
      },
    });
  }
}
