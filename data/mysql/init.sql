CREATE DATABASE IF NOT EXISTS xinclockin;

use xinclockin;

CREATE TABLE
    IF NOT EXISTS `user` (
        `id` INT(20) NOT NULL AUTO_INCREMENT,
        `identity_number` varchar(50) NOT NULL,
        `uid` varchar(36) NOT NULL,
        `full_name` varchar(255) DEFAULT NULL,
        `avatar_url` varchar(255) DEFAULT NULL,
        `session_id` varchar(255) DEFAULT NULL,
        `biography` varchar(255) DEFAULT NULL,
        `private_key` varchar(255) DEFAULT NULL,
        `client_id` VARCHAR(255) DEFAULT NULL,
        `contract` VARCHAR(255) DEFAULT NULL,
        `is_mvm_user` TINYINT(1) DEFAULT 0,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        UNIQUE KEY `idx_identity_number` (`identity_number`),
        UNIQUE KEY `idx_uid` (`uid`),
        KEY `idx_full_name` (`full_name`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 打卡项目表

CREATE TABLE
    `clockin` (
        `cid` VARCHAR(36) NOT NULL,
        `uid` VARCHAR(36) NOT NULL,
        `title` VARCHAR(200) NOT NULL,
        `description` VARCHAR(256) DEFAULT NULL,
        `word_count` INT UNSIGNED DEFAULT 0,
        `cover` LONGTEXT,
        `start_day` TIMESTAMP NOT NULL COMMENT 'Start time',
        `end_day` TIMESTAMP NOT NULL,
        `clockin_start_daily` TIMESTAMP NOT NULL COMMENT 'every day clockin start time',
        `clockin_end_daily` TIMESTAMP NOT NULL COMMENT 'every day clockin end time',
        `total_fee` INT UNSIGNED DEFAULT 0,
        `bonus_daily` INT UNSIGNED DEFAULT 0,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `deleted_at` TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (`uid`),
        KEY `created_at` (`created_at`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 打卡话题表

CREATE TABLE
    `content` (
        `uid` VARCHAR(36) NOT NULL,
        `content` longtext NOT NULL,
        `clockin_id` VARCHAR(36) NOT NULL,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `deleted_at` TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (`uid`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE TABLE
    `user_clockin` (
        `uid` VARCHAR(36) NOT NULL,
        `cid` VARCHAR(36) NOT NULL,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `deleted_at` TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;