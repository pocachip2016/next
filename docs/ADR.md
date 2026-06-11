# Architecture Decision Records: next (rightwatch)

## 철학
- **MVP 속도**: 기존 스택(Go/Scrapy/Angular) 유지, 새 의존성 최소화.
- **규칙 우선**: 변형탐지는 자모분해+homoglyph 규칙기반 시작, ML은 데이터 충분시 후속.
- **데이터 우선**: 크롤 안정성 > 완벽한 매칭. 탐지 누락보다 오탐은 더 심각 (법적 위험).
- **외부화 규칙**: 시크릿/자격증명은 환경변수로만, 코드/git에 평문 금지.

---

### ADR-001: 스택 선택 (Go/Gin + Scrapy + Angular)
**결정**: Go(API), Python/Scrapy(크롤러), Angular 14(FE), MySQL(DB) 유지.
**이유**: 
- 기존 구현이 안정적(ondisk 크롤 운영 중).
- Go는 API 성능, Scrapy는 크롤 성숙도, Angular는 관리자 UI 적합.
- 팀 역량: pocachip 기존 경험 있음.

**트레이드오프**: 
- 모놀리식 repo → 마이크로서비스 이전 비용 (장기 고려).
- 스택 종류 많음 → 배포/모니터링 복잡도 (docker-compose로 완화).

---

### ADR-002: 콘텐츠 마스터 (kta_contents vs contents_list)
**결정**: `kta_contents`를 콘텐츠 SSOT(single source of truth)로, `contents_list`는 M1에서 정리.
**이유**: 
- `kta_contents`: genre, actors, director, synop 등 리치 메타(9컬럼).
- `contents_list`: title만 저장(legacy, 중복).
- 매칭 정확도 향상 (제목만→배우/감독도 비교).

**트레이드오프**: 
- 프론트 `contents-list` 컴포넌트는 현재 `contents_list`에 쿼리 중 → M1 의존성 스캔 & 수정 필요.
- 기존 데이터 마이그레이션 필요 (M1).

---

### ADR-003: 변형 패턴 탐지 — 규칙기반 먼저
**결정**: 자모분해 + homoglyph 정규화 + synonym 사전(규칙) 기반 시작. 유사도/ML은 후속.
**이유**: 
- 빠른 구현 (정규식 > ML 학습).
- "기도"→"7ㅣ도" 같은 명확한 변형은 규칙으로 충분.
- 오탐 위험 낮음 (ML은 신뢰도 낮음 초기 단계).

**트레이드오프**: 
- 정교한 변형("기드오" 등 순서 뒤바뀜)은 규칙 커버 불완전.
- 새 변형 등장시 사전 갱신 필요 (운영 비용).
- 미래에 fuzzy matching/ML로 전환시 재설계.

---

### ADR-004: 모니터링 대상 웹하드 — ondisk 우선
**결정**: M1~M4는 ondisk만 완성. filesun, torrent 등은 M4 이후 추가.
**이유**: 
- ondisk: 이미 크롤 안정적(운영 중), 불법 VOD 점유율 높음.
- scope 제한 → M0~M3 집중 가능.
- 구조 확장 가능 (spider 추가 모듈화).

**트레이드오프**: 
- 초기 커버리지 제한(진출 지연).
- 각 웹하드별 로그인/구조 상이 → 각각 spider 필요 (확장 비용).

---

### ADR-005: 시크릿 외부화 원칙
**결정**: `.env` (git 무시) 기반 환경변수 주입. 하드코딩 금지.
**이유**: 
- 보안 (토큰 폐기 용이, 히스토리 노출 방지).
- 운영 (개발/테스트/프로드 환경별 다른 자격증명).
- 법적 (솔로몬 감시 규정상 개인정보 보호).

**트레이드오프**: 
- 로컬 개발시 `.env` 수동 설정 필요.
- 앱 config 외부화 미완료 (go/python `config.toml`/`settings.py`는 아직 하드코딩) → M1 follow-up.

**상태**: docker-compose ✓ (M0-2 완료), 앱 config는 follow-up (M1).

---

### ADR-006: 데이터 파이프라인 — 비동기 + 배치
**결정**: 크롤(비동기) 분리, 매칭·알람은 rightwatch API 배치 태스크.
**이유**: 
- 크롤 오류 격리 (crashing crawler가 API 다운 아님).
- 매칭 재처리 용이 (post 변경 감지 후 재실행).

**트레이드오프**: 
- 실시간성 제약 (크롤 12h + 매칭 주기적 실행).
- 상태 추적 복잡 (job 로그 필수).

**M4 검토**: 검색기반 신규탐지 추가시 준실시간(1~2h) 전환 검토.
