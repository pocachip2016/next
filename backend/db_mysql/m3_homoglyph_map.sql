-- M3: homoglyph_map — 변형 문자 정규화 사전
CREATE TABLE IF NOT EXISTS homoglyph_map (
  id        INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  from_char VARCHAR(16)  NOT NULL COMMENT '변형 문자',
  to_char   VARCHAR(16)  NOT NULL COMMENT '정규화 대상 문자'
) DEFAULT CHARSET=utf8mb4;

-- 초기 seed: 숫자/특수문자 → 알파벳, 한글 자모 → 알파벳
INSERT INTO homoglyph_map (from_char, to_char) VALUES
  ('0',  'o'),  -- 숫자 0 → o
  ('ο',  'o'),  -- 그리스 소문자 omicron
  ('О',  'o'),  -- 키릴 대문자 O
  ('о',  'o'),  -- 키릴 소문자 o
  ('1',  'l'),  -- 숫자 1 → l
  ('ı',  'i'),  -- 터키어 dotless i
  ('І',  'i'),  -- 키릴 대문자 І
  ('і',  'i'),  -- 키릴 소문자 і
  ('ㅇ', 'o'),  -- 한글 이응
  ('ㅣ', 'i'),  -- 한글 이
  ('ㅡ', 'u'),  -- 한글 으
  ('ㄱ', 'k'),  -- 한글 기역
  ('@',  'a'),  -- @ → a
  ('3',  'e'),  -- 3 → e (leet)
  ('4',  'a'),  -- 4 → a (leet)
  ('5',  's'),  -- 5 → s (leet)
  ('8',  'b');  -- 8 → b (leet)
