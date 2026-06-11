# TODO — next

> **세션 재개 프롬프트**: "TODO.md 확인하고 `## Now`부터 이어서 진행해"

## Now (진행 중, 1~3개)
- **M0-1/2/3**: 보안 위생 + 문서화 (branch: `chore/rightwatch-m0-hygiene`)
  - [x] **M0-1** — read.md 시크릿 제거 + dev-notes.md 이전
  - [x] **M0-2** — docker-compose 환경변수 외부화 (`.env`)
  - [ ] **M0-3** — 문서 작성 (PRD/ARCHITECTURE/ADR)
  - [ ] **M0-4** — TODO/plans 로드맵 확정

## Next (이번 마일스톤)
- **M1**: 도메인 모델 정리
  - CP 개념 도입 (모델/API)
  - `kta_contents` vs `contents_list` 중복 정리
  - ondisk 로그인 자격증명 외부화 (settings.py)
  
- **M2**: 매칭 엔진 (핵심 기능)
  - 제목 정규화 (한글 자모분해 + 공백/특수문자)
  - synonym_words 사전 확장
  - post ↔ kta_contents 매칭 알고리즘
  - check_list 자동 적재
  - API endpoint: POST `/matching/check` 

## Later (백로그)
- **M3**: 변형 패턴 탐지
  - homoglyph 정규화 (O/0, l/1, etc)
  - 변형 사전 DB 확장
  - 탐지 정확도 테스트
  
- **M4**: 스케줄러 + 검색기반 탐지
  - 콘텐츠 갱신주기 기반 매칭 스케줄
  - 검색 기반 신규 콘텐츠 탐지
  - crawler_job 상태 관리 (시작/종료/결과)
  
- **M5**: CP 대시보드
  - CP별 집계 API (탐지 건수/삭제율/웹하드별 분포)
  - Angular rightwatch 페이지 (content-panel → check 결과 표시)
  - CP별 리포트 생성/다운로드
  
- **M6**: 알람 + 삭제 추적
  - 상태 머신 (탐지 → 승인 대기 → 통보 → 삭제확인 → 종결)
  - SMTP 메일 발송 (CP에게 탐지 통보)
  - 삭제 여부 재확인 flow
  
- **M7** (선택): 심화 검증
  - 상세 이미지 캡처 (playwright)
  - 해시 기반 이미지 유사도
  - (선택) 영상 DNA 검증
  
- **확장**: 다수 웹하드 지원
  - filesun 스파이더 완성
  - 추가 웹하드 spider 플러그인화
  - 멀티 웹하드 통합 매칭

## Done (최근 5개만)
- ondisk_spider3 로그인 폼 수정 — www.ondisk.co.kr → ondisk.co.kr로 도메인 통일, start_requests + AJAX 4곳 + Host 헤더 모두 무www로 통일 (쿠키 도메인 일관성 보장)
- pricemon DB 스키마 생성 — content/product/price_list/price_attime 4테이블 (`backend/db_mysql/pricemon_schema.sql`), Docker MySQL 적용 + API 4종 조회 검증
- Docker 전체 스택 구성 (mysql + rightwatch + pricemon + frontend + crawler)
- Angular 14 프론트엔드 빌드 이슈 해결 (babel/runtime, @types/ws, node-sass)
- supercronic 크론 스케줄러 등록 (4개 카테고리 12h 간격)
