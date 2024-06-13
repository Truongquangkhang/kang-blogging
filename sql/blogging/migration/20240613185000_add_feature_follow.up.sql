CREATE TABLE `follows` (
     `follower_id` VARCHAR(40) NOT NULL,
     `followed_id` VARCHAR(40) NOT NULL,
     `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP),
     PRIMARY KEY (follower_id, followed_id)
);
