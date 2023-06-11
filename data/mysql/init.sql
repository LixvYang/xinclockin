CREATE DATABASE IF NOT EXISTS xinclockin;

use xinclockin;

CREATE TABLE
    IF NOT EXISTS `user` (
        `id` int NOT NULL AUTO_INCREMENT,
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
        `id` int(20) NOT NULL AUTO_INCREMENT,
        `title` varchar(200) NOT NULL,
        `description` longtext DEFAULT NULL,
        `cover` varchar(255) DEFAULT NULL,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `deleted_at` TIMESTAMP DEFAULT NULL,
        `end_at` timestamp NOT NULL,
        `clock_time_daily` timestamp NOT NULL,
        PRIMARY KEY (`id`),
        KEY `created_at` (`created_at`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- 打卡话题表

CREATE TABLE
    `playground` (
        `id` int(20) NOT NULL AUTO_INCREMENT,
        `uid` varchar(36) NOT NULL,
        `content` longtext NOT NULL,
        `clockin_id` int(20) NOT NULL,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        `deleted_at` TIMESTAMP DEFAULT NULL,
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;