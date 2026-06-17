# 아키텍처: next (rightwatch)

## 디렉토리 구조
```
next/
├── backend/
│   ├── crawler/
│   │   └── scrapy/ondisk/  (Scrapy 프로젝트, 웹하드 크롤링)
│   ├── apiServer/
│   │   ├── rightwatch/     (모니터링 메인 API, Go/Gin 포트 5555)
│   │   └── pricemon/       (가격 모니터링 API, Go/Gin 포트 5556)
│   └── db_mysql/           (SQL 스키마, 마이그레이션)
├── frontend/
│   └── ngx-admin2/         (Angular 14, 포트 4200)
├── docker-compose.yml      (통합 스택 정의)
├── .env.example            (환경변수 템플릿)
├── docs/
│   ├── PRD.md, ARCHITECTURE.md, ADR.md
│   └── dev-notes.md        (개발 메모, 커맨드)
└── TODO.md, CLAUDE.md
```

## 데이터 흐름
```
웹하드(ondisk)
    ↓ [Scrapy 크롤러, 12h 주기]
Post 테이블
    ↓ [매칭엔진, 규칙기반]
CheckList 테이블 (제목+변형패턴 매칭 결과)
    ↓ [CP 집계, 이메일 발송]
CP 대시보드 / 알림
    ↓ [운영자 재확인 & 종결]
최종 리포트
```

## 컴포넌트
- **Crawler (Scrapy, Python 3.11)**: Playwright 기반 로그인 → 게시판 전수 크롤 → 상세페이지 파싱 → `post` 테이블 적재. supercronic 스케줄러.
- **rightwatch API (Go/Gin)**: CRUD API (post, check_list, kta_contents, synonym_words, crawler_job). 매칭엔진·상태머신 구현 (M2~M6).
- **pricemon API (Go/Gin)**: 가격 모니터링 (현재 미사용).
- **Frontend (Angular 14, ngx-admin)**: rightwatch 페이지 (content-panel, check-panel, post-list). CP 대시보드 (M5).
- **MySQL 8.0**: 통합 데이터베이스. `post`(캐시), `kta_contents`(콘텐츠마스터), `check_list`(매칭결과), `synonym_words`(변형사전), `crawler_job`(스케줄), `screenshots`(이미지).

## 외부 의존성
| 서비스 | 용도 | 비고 |
|---|---|---|
| ondisk.co.kr | 웹하드 게시판 | 포트 443(HTTPS), robots.txt 준수, 요청간격 2s |
| filesun.net | 웹하드 게시판 | (M4 이후) |
| SMTP (미정) | CP 이메일 통보 | (M6, 설정 필요) |

## 설계 패턴
- **CRUD 스캐폴딩**: ginbro 자동생성 (Go API), 비즈니스 로직은 핸들러에 추가.
- **규칙기반 매칭**: 정규화(자모분해+homoglyph) → synonym 사전 확장 → 문자열 매칭.
- **비동기 처리**: Scrapy 크롤러는 async, 후처리(매칭)는 rightwatch API 태스크.
- **테이블 정규화**: `kta_contents`(SSOT, 리치 메타) vs `contents_list`(legacy, title-only) — M1 정리 예정.

## 주요 제약
- **CRITICAL**: 시크릿(DB 비밀번호, API 키)은 `.env` 기반 환경변수로만 주입. 하드코딩 금지.
- **CRITICAL**: 웹하드 robots.txt/이용약관 준수. 요청 간격 ≥ 2초 (법적 준수 + 차단 방지).
- **CRITICAL**: ondisk 로그인 자격증명은 `.env`로 외부화 (M1 후속).
- **성능**: 크롤 건당 게시물 1000+ 파싱 — 정규식/병렬화로 최적화 필요.
- **데이터 품질**: `post.txt` 제목 길이 제한(256자) — 긴 제목은 truncate 또는 전문 필드 추가 (M1 검토).
