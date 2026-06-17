-- M0: kta_contents(콘텐츠 카탈로그) 베이스 스키마
-- 목적: 누락돼 있던 kta_contents 테이블 정의를 마이그레이션으로 커밋.
--       컬럼은 apiServer/rightwatch/models/model_kta_contents.go 기준.
--       cp_id 컬럼/FK는 m1_add_cp.sql이 ALTER로 추가하므로 여기서는 제외.
-- 데이터: kta_contents.json (2135건) 은 별도 적재 스크립트로 INSERT.

CREATE TABLE IF NOT EXISTS `kta_contents` (
  `id`        INT          NOT NULL AUTO_INCREMENT,
  `genre`     VARCHAR(256) DEFAULT NULL COMMENT '장르',
  `title`     VARCHAR(256) DEFAULT NULL COMMENT '제목',
  `actors`    TEXT                  COMMENT '출연진',
  `director`  VARCHAR(256) DEFAULT NULL COMMENT '감독',
  `price`     VARCHAR(256) DEFAULT NULL COMMENT '가격',
  `enddate`   VARCHAR(256) DEFAULT NULL COMMENT '서비스 종료일',
  `synop`     TEXT                  COMMENT '줄거리',
  `p_url`     VARCHAR(256) DEFAULT NULL COMMENT '포스터 URL',
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 검증 쿼리:
-- SELECT COUNT(*) FROM kta_contents;
-- SELECT id, title, genre FROM kta_contents LIMIT 5;
