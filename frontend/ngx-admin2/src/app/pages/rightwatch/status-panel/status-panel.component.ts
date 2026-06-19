import { Component, OnInit } from '@angular/core';
import { RightwatchService, ICheckListItem, ICapture } from '../rightwatch.service';

@Component({
  selector: 'ngx-status-panel',
  templateUrl: './status-panel.component.html',
  styleUrls: ['./status-panel.component.scss'],
})
export class StatusPanelComponent implements OnInit {
  items: ICheckListItem[] = [];
  loading = true;
  error = '';
  results: Record<number, string> = {};

  captureLoading: Record<number, boolean> = {};
  captureMsg: Record<number, string> = {};
  showCaptures: Record<number, boolean> = {};
  captures: Record<number, ICapture[]> = {};

  private readonly staticBase = 'http://127.0.0.1:5555/static/captures';

  constructor(private service: RightwatchService) {}

  ngOnInit() {
    this.load();
  }

  refresh() {
    this.results = {};
    this.load();
  }

  private load() {
    this.loading = true;
    this.error = '';
    this.service.getCheckList().subscribe({
      next: (res: any) => {
        this.items = res.data || [];
        this.loading = false;
      },
      error: (err: any) => {
        this.error = err.message || '데이터를 불러오지 못했습니다.';
        this.loading = false;
      },
    });
  }

  statusLabel(s: number): string {
    switch (s) {
      case 0: return '탐지대기';
      case 1: return '통보완료';
      case 2: return '삭제확인';
      case 3: return '종결';
      default: return `status(${s})`;
    }
  }

  statusClass(s: number): string {
    switch (s) {
      case 0: return 'status-warning';
      case 1: return 'status-info';
      case 2: return 'status-success';
      case 3: return 'status-basic';
      default: return '';
    }
  }

  actionLabel(s: number): string {
    switch (s) {
      case 0: return '통보 발송';
      case 1: return '삭제 확인';
      case 2: return '종결';
      default: return '';
    }
  }

  captureImage(item: ICheckListItem) {
    this.captureLoading[item.id] = true;
    delete this.captureMsg[item.id];
    this.service.captureCreate(item.id).subscribe({
      next: (res: any) => {
        const cap: ICapture = res.data;
        this.captureMsg[item.id] = '캡처 완료';
        this.captureLoading[item.id] = false;
        if (!this.captures[item.id]) {
          this.captures[item.id] = [];
        }
        this.captures[item.id] = [...this.captures[item.id], cap];
        this.showCaptures[item.id] = true;
      },
      error: (e: any) => {
        this.captureMsg[item.id] = e.error?.msg || e.message || '캡처 실패';
        this.captureLoading[item.id] = false;
      },
    });
  }

  toggleCaptures(item: ICheckListItem) {
    const id = item.id;
    if (this.showCaptures[id]) {
      this.showCaptures[id] = false;
      return;
    }
    if (this.captures[id]) {
      this.showCaptures[id] = true;
      return;
    }
    this.captureLoading[id] = true;
    this.service.captureList(id).subscribe({
      next: (res: any) => {
        this.captures[id] = res.data || [];
        this.showCaptures[id] = true;
        this.captureLoading[id] = false;
      },
      error: () => {
        this.captureLoading[id] = false;
      },
    });
  }

  captureUrl(cap: ICapture): string {
    return `${this.staticBase}/${cap.id}.jpg`;
  }

  performAction(item: ICheckListItem) {
    item.loadingAction = true;
    delete this.results[item.id];

    const done = (msg: string) => {
      this.results[item.id] = msg;
      item.loadingAction = false;
      this.load();
    };

    if (item.status === 0) {
      this.service.transition(item.id, 0, 1).subscribe({
        next: () => done('통보 발송됨'),
        error: (e: any) => { this.results[item.id] = e.error?.msg || e.message; item.loadingAction = false; },
      });
    } else if (item.status === 1) {
      this.service.confirmDeletion(item.id).subscribe({
        next: (res: any) => {
          const d = res.data;
          done(d?.deleted ? '삭제 확인됨 → 삭제확인 상태로 전이' : `게시물 잔존 (미전이): ${d?.reason || ''}`);
        },
        error: (e: any) => { this.results[item.id] = e.error?.msg || e.message; item.loadingAction = false; },
      });
    } else if (item.status === 2) {
      this.service.transition(item.id, 2, 3).subscribe({
        next: () => done('종결 처리됨'),
        error: (e: any) => { this.results[item.id] = e.error?.msg || e.message; item.loadingAction = false; },
      });
    }
  }
}
