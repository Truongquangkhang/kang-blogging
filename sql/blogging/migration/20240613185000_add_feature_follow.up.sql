CREATE TABLE `follows` (
     `follower_id` varchar(40) NOT NULL,
     `followed_id` varchar(40) NOT NULL,
     `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE follows
    MODIFY follower_id VARCHAR(40) COLLATE utf8mb4_0900_ai_ci,
    MODIFY followed_id VARCHAR(40) COLLATE utf8mb4_0900_ai_ci,
    ADD PRIMARY KEY (follower_id, followed_id),
    ADD FOREIGN KEY (`follower_id`) REFERENCES `users` (`id`),
    ADD FOREIGN KEY (`followed_id`) REFERENCES `users` (`id`);


