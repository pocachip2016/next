-- M0: 누락된 코어 테이블 스키마 (모델만 있고 마이그레이션이 없던 테이블)
-- 목적: synonym_words / crawler_job / screenshots 테이블 정의를 커밋.
--       컬럼은 apiServer/rightwatch/models/model_*.go 기준.
-- 데이터: 비어있어도 동작 (synonym 없으면 동의어 확장만 생략).

-- 1. synonym_words: 동의어 사전 (매칭 시 제목 확장)
CREATE TABLE IF NOT EXISTS `synonym_words` (
  `id`       INT          NOT NULL AUTO_INCREMENT,
  `pair_id`  INT          NOT NULL COMMENT '동의어 그룹 ID',
  `synonym`  VARCHAR(128) NOT NULL COMMENT '동의어',
  PRIMARY KEY (`id`),
  KEY `idx_pair_id` (`pair_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 2. crawler_job: 크롤링/매칭 작업 이력
CREATE TABLE IF NOT EXISTS `crawler_job` (
  `id`          INT         NOT NULL AUTO_INCREMENT,
  `boundary`    TINYINT     DEFAULT NULL COMMENT '크롤 경계 플래그',
  `cat1_code`   VARCHAR(10) DEFAULT NULL COMMENT '카테고리 코드',
  `start_time`  TIMESTAMP   NULL DEFAULT NULL COMMENT '시작 시각',
  `end_time`    TIMESTAMP   NULL DEFAULT NULL COMMENT '종료 시각',
  `result`      VARCHAR(32) DEFAULT NULL COMMENT '결과 메시지',
  `website`     INT         DEFAULT NULL COMMENT '웹하드 코드',
  PRIMARY KEY (`id`),
  KEY `idx_cat1_code` (`cat1_code`),
  KEY `idx_website` (`website`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 3. screenshots: 상세 캡처 (M7 심화 검증용)
CREATE TABLE IF NOT EXISTS `screenshots` (
  `url`      VARCHAR(128) NOT NULL COMMENT '캡처 대상 URL',
  `url_md5`  VARCHAR(32)  NOT NULL COMMENT 'URL MD5 해시',
  `ts`       VARCHAR(256) DEFAULT NULL COMMENT '캡처 시각',
  PRIMARY KEY (`url_md5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 검증 쿼리:
-- SELECT COUNT(*) FROM synonym_words;
-- SELECT COUNT(*) FROM crawler_job;
-- SELECT COUNT(*) FROM screenshots;
