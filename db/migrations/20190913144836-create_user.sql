-- +migrate Up
CREATE TABLE `user` (
    `id` int NOT NULL AUTO_INCREMENT,
    `uuid` VARCHAR(36) NOT NULL,
    `role_uuid`  VARCHAR(36) NOT NULL,
    `account` VARCHAR(36) NOT NULL COMMENT '帳號',
    `password` BINARY(60) NOT NULL COMMENT '密碼',
    `name` VARCHAR(36) NOT NULL COMMENT '姓名',
    `email` VARCHAR(36) NOT NULL COMMENT '信箱',
    `status` ENUM('enabled', 'disabled') DEFAULT 'enabled' NOT NULL COMMENT '狀態',
    `score` int NOT NULL COMMENT '分數',
    `verification` boolean NOT NULL COMMENT '驗證',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `created_by` VARCHAR(36) COMMENT '創建者',
    `updated_by` VARCHAR(36) COMMENT '更新者',
    UNIQUE INDEX (`account`),
    UNIQUE INDEX (`email`),
    UNIQUE INDEX (`uuid`),
    PRIMARY KEY (`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='使用者';
-- +migrate Down
DROP TABLE IF EXISTS `user`;