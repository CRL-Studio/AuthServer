-- +migrate Up
CREATE TABLE `role` (
    `id` int NOT NULL AUTO_INCREMENT,
    `uuid` VARCHAR(36) NOT NULL,
    `name` VARCHAR(36) NOT NULL COMMENT '名稱',
    `code` VARCHAR(36) NOT NULL COMMENT '代碼',
    `created_at` DATETIME NOT NULL DEFAULT NOW() COMMENT '創建日期',
    `updated_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `created_by` VARCHAR(36) COMMENT '創建者',
    `updated_by` VARCHAR(36) COMMENT '更新者',
    UNIQUE INDEX (`uuid`),
    UNIQUE INDEX (`code`),
    UNIQUE INDEX (`name`),
    PRIMARY KEY (`id`)
) CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色';
-- +migrate Down
DROP TABLE IF EXISTS `role`;