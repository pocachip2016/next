# TODO — next

> **세션 재개 프롬프트**: "TODO.md 확인하고 `## Now`부터 이어서 진행해"

## Now (진행 중, 1~3개)
- **M7** (선택): 심화 검증 — 상세 이미지 캡처 + 해시 유사도

## Next (이번 마일스톤)
- (없음)

## Later (백로그)
- **M5-3**: CP별 리포트 생성/다운로드 (CSV 다운로드)
- **M7** (선택): 심화 검증
  - 상세 이미지 캡처 (playwright)
  - 해시 기반 이미지 유사도
  - (선택) 영상 DNA 검증
- **확장**: 다수 웹하드 지원
  - filesun 스파이더 완성
  - 추가 웹하드 spider 플러그인화
  - 멀티 웹하드 통합 매칭

## Done (최근 5개만)
- **M6**: 알람 + 삭제 추적 — check_list 상태 머신(0~3) + SMTP dry-run 메일러 + /transition·/confirm-deletion API + Angular StatusPanelComponent + DB 스키마 보강(kta_contents/synonym_words 등)
- **M5**: CP 대시보드 — GET /api/cp/dashboard + Angular CpDashboardComponent (M5-3 리포트는 Later)
- **M4**: 스케줄러 + 검색기반 탐지 — RunMatchingSince(증감분), RunMatchingForContent(신규콘텐츠), crawler_job 이력
- **M3**: 변형 패턴 탐지 — homoglyph_map DB + SetHomoglyphMap/NormalizeHomoglyph + 단위 테스트 4종
- **M2**: 매칭 엔진 — normalize/matcher 패키지, POST /matching/run, GET /matching/status, GetDB()
- **M1**: 도메인 모델 정리 완료 — CP 테이블/API 추가, kta_contents cp_id FK, contents_list 레거시 API 제거, Scrapy 환경변수 외부화
- **M0**: 보안/문서화 완료 — read.md 시크릿 제거, docker-compose 환경변수 외부화 (.env), PRD/ARCHITECTURE/ADR 작성, 로드맵 확정 (M0~M7)
- ondisk_spider3 로그인 폼 수정 — www.ondisk.co.kr → ondisk.co.kr로 도메인 통일, start_requests + AJAX 4곳 + Host 헤더 모두 무www로 통일
- pricemon DB 스키마 생성 — content/product/price_list/price_attime 4테이블, Docker MySQL 적용 + API 4종 조회 검증
- Docker 전체 스택 구성 (mysql + rightwatch + pricemon + frontend + crawler)
