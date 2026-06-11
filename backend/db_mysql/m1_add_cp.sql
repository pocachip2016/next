-- M1: CP(저작권자) 도메인 모델 추가
-- 날짜: 2026-06-11
-- 목적: CP 테이블 생성 + kta_contents에 cp_id FK 추가

-- 1. CP 테이블 생성
-- 저작권자(영화사, 배급사 등) 정보 관리
CREATE TABLE IF NOT EXISTS `cp` (
  `id`           INT            NOT NULL AUTO_INCREMENT,
  `name`         VARCHAR(256)   NOT NULL COMMENT '저작권자명',
  `email`        VARCHAR(256)   DEFAULT NULL COMMENT '연락처 이메일',
  `phone`        VARCHAR(64)    DEFAULT NULL COMMENT '연락처 전화',
  `created_at`   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP COMMENT '생성 시각',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 2. kta_contents에 cp_id FK 추가
-- 기존 데이터와의 호환성 유지 (NULL 허용)
ALTER TABLE `kta_contents`
  ADD COLUMN `cp_id` INT NULL DEFAULT NULL COMMENT 'CP ID (저작권자)' AFTER `id`;

ALTER TABLE `kta_contents`
  ADD CONSTRAINT `fk_kta_contents_cp`
  FOREIGN KEY (`cp_id`) REFERENCES `cp` (`id`)
  ON DELETE SET NULL
  ON UPDATE CASCADE;

-- 3. 기존 kta_contents 데이터는 cp_id=NULL로 유지 (운영자가 나중에 수동 매핑)

-- 검증 쿼리:
-- SELECT * FROM cp;
-- SELECT id, title, cp_id FROM kta_contents LIMIT 5;
