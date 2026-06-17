-- M2: 매칭 엔진 — check_list 스키마 정정
-- 날짜: 2026-06-11
-- 목적: content_id varchar(256) → INT (kta_contents.id 타입 통일)

-- 1. 기존 UNIQUE KEY 제거
ALTER TABLE `check_list`
  DROP INDEX `unique_index`;

-- 2. content_id 타입 변경: varchar(256) → INT
-- 기존 데이터가 숫자 문자열이면 자동 캐스팅됨
-- 예: "1234" → 1234
ALTER TABLE `check_list`
  MODIFY COLUMN `content_id` INT NULL DEFAULT NULL COMMENT 'CP 콘텐츠 ID (kta_contents.id)';

-- 3. UNIQUE KEY 재추가
-- post_id는 varchar 유지 (website:idx 조합 문자열)
ALTER TABLE `check_list`
  ADD UNIQUE KEY `unique_index` (`content_id`, `post_id`(191));

-- 4. FK 추가 (기존 없음, M1 이후 추가 고려)
-- ALTER TABLE `check_list`
--   ADD CONSTRAINT `fk_check_list_kta_contents`
--   FOREIGN KEY (`content_id`) REFERENCES `kta_contents` (`id`)
--   ON DELETE CASCADE;

-- 검증 쿼리:
-- SELECT COUNT(*) FROM check_list WHERE content_id IS NULL;
-- SELECT COUNT(*), content_id FROM check_list GROUP BY content_id LIMIT 5;
