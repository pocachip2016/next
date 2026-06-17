-- M7: 증거 캡처 히스토리 테이블
-- 목적: check_list 항목별 화면 캡처를 시간순으로 누적 (목록/상세/삭제확인).
--       레거시 screenshots(url_md5 PK)는 단일 캡처용이라 히스토리를 담을 수 없어 신규 테이블 분리.
-- 컬럼은 apiServer/rightwatch/models/model_capture.go 기준.

CREATE TABLE IF NOT EXISTS `capture` (
  `id`            INT          NOT NULL AUTO_INCREMENT,
  `check_list_id` INT          NOT NULL COMMENT 'check_list.id FK',
  `capture_type`  VARCHAR(16)  NOT NULL COMMENT 'list|detail|takedown',
  `file_path`     VARCHAR(512) NOT NULL COMMENT '/app/screenshots/{id}.jpg',
  `page_url`      VARCHAR(512) DEFAULT NULL COMMENT '캡처 대상 URL',
  `status_at`     TINYINT      DEFAULT NULL COMMENT '캡처 시점 check_list.status',
  `trigger_by`    VARCHAR(8)   DEFAULT NULL COMMENT 'auto|manual',
  `captured_at`   TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP COMMENT '캡처 시각',
  PRIMARY KEY (`id`),
  KEY `idx_check_list` (`check_list_id`),
  KEY `idx_captured_at` (`captured_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 검증 쿼리:
-- SELECT id, check_list_id, capture_type, trigger_by, captured_at FROM capture ORDER BY captured_at DESC LIMIT 10;
