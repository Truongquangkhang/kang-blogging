CREATE TABLE violations (
   `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
   `user_id` varchar(40) NOT NULL,
   `violation_type` varchar(255) NOT NULL,
   `violation_target_id` varchar(40) NOT NULL,
   `description` text,
   `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE `violations`
    MODIFY user_id VARCHAR(40) COLLATE utf8mb4_0900_ai_ci,
    MODIFY violation_target_id VARCHAR(40) COLLATE utf8mb4_0900_ai_ci,
    ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
