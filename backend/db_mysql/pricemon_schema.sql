-- pricemon schema
-- Generated to match gorm models in backend/apiServer/pricemon/models/.
-- AutoMigrate() is disabled in models/db.go, so tables are created manually.
-- Apply: docker exec -i next-mysql mysql -upocachip -p'media2015!' pricemon < pricemon_schema.sql

SET NAMES utf8mb4;

CREATE TABLE IF NOT EXISTS `content` (
  `id`    BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `product` (
  `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name`        VARCHAR(256) NOT NULL DEFAULT '',
  `description` VARCHAR(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `price_list` (
  `id`            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `price_name`    VARCHAR(20) NOT NULL DEFAULT '',
  `product_price` INT NOT NULL DEFAULT 0,
  `product_line`  INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `price_attime` (
  `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `content_id` INT NOT NULL DEFAULT 0,
  `price_id`   INT NOT NULL DEFAULT 0,
  `start_date` DATETIME NULL,
  `end_date`   DATETIME NULL,
  PRIMARY KEY (`id`),
  KEY `idx_price_attime_content_id` (`content_id`),
  KEY `idx_price_attime_price_id` (`price_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
