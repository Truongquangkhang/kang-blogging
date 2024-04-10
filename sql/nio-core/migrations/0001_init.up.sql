SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for voucher
-- ----------------------------

DROP TABLE IF EXISTS `vouchers`;
CREATE TABLE `vouchers`
(
    `id`             BIGINT PRIMARY KEY AUTO_INCREMENT,
    `voucher_code`   VARCHAR(255) NOT NULL,
    `voucher_source` VARCHAR(255) NOT NULL,
    `voucher_type`   VARCHAR(255) NOT NULL,
    `summary`        VARCHAR(255) NOT NULL,
    `start_time`     DATETIME,
    `end_time`       DATETIME,
    `deleted`        BOOLEAN      NOT NULL DEFAULT FALSE,
    `created_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_voucher_code (`voucher_code`)
)
    CHARACTER SET = utf8mb4
    COMMENT = 'Store all vouchers';

SET FOREIGN_KEY_CHECKS = 1;
