-- M6: 알람 + 삭제 추적 — 상태 머신 스키마
-- 날짜: 2026-06-11
-- 목적: check_list 상태 타임스탬프 추가 + notification_log 테이블 생성

-- 1. check_list.status COMMENT 갱신 + 타임스탬프 컬럼 추가
-- 상태 코드: 0=detected_pending / 1=notified / 2=delete_confirmed / 3=closed
ALTER TABLE `check_list`
  MODIFY COLUMN `status` TINYINT NOT NULL DEFAULT 0
    COMMENT '0=detected_pending,1=notified,2=delete_confirmed,3=closed';

ALTER TABLE `check_list`
  ADD COLUMN `notified_at`          TIMESTAMP NULL DEFAULT NULL COMMENT '통보 시각' AFTER `status`,
  ADD COLUMN `delete_confirmed_at`  TIMESTAMP NULL DEFAULT NULL COMMENT '삭제확인 시각' AFTER `notified_at`,
  ADD COLUMN `closed_at`            TIMESTAMP NULL DEFAULT NULL COMMENT '종결 시각' AFTER `delete_confirmed_at`;

-- 2. notification_log 테이블 생성
CREATE TABLE IF NOT EXISTS `notification_log` (
  `id`              INT           NOT NULL AUTO_INCREMENT,
  `check_list_id`   INT           NOT NULL COMMENT 'check_list.id',
  `cp_id`           INT           NULL     COMMENT 'cp.id (통보 대상 CP)',
  `to_email`        VARCHAR(256)  NOT NULL COMMENT '수신 이메일',
  `subject`         VARCHAR(256)  NOT NULL COMMENT '메일 제목',
  `body`            TEXT          NOT NULL COMMENT '메일 본문',
  `dry_run`         TINYINT       NOT NULL DEFAULT 0 COMMENT '1=실제 발송 안 함(SMTP 미설정)',
  `result`          VARCHAR(256)  NOT NULL DEFAULT '' COMMENT '발송 결과 메시지',
  `sent_at`         TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '기록 시각',
  PRIMARY KEY (`id`),
  KEY `idx_check_list_id` (`check_list_id`),
  KEY `idx_sent_at` (`sent_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 검증 쿼리:
-- SHOW COLUMNS FROM check_list;
-- SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_NAME='notification_log';
