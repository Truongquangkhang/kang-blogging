CREATE TABLE reports (
    `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
    `reporter_id` varchar(40) NOT NULL,
    `report_type` varchar(255) NOT NULL,
    `report_target_id` varchar(40) NOT NULL,
    `reason` text NOT NULL,
    `description` text,
    `is_closed` bool NOT NULL  DEFAULT FALSE,
    `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE `reports`
    MODIFY reporter_id VARCHAR(40) COLLATE utf8mb4_0900_ai_ci,
    MODIFY report_target_id VARCHAR(40) COLLATE utf8mb4_0900_ai_ci,
    ADD FOREIGN KEY (`reporter_id`) REFERENCES `users` (`id`);
