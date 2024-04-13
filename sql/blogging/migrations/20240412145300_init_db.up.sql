CREATE TABLE `accounts` (
    `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
    `username` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL COMMENT 'hash password',
    `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE `roles` (
     `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
     `name` varchar(255) NOT NULL
);

CREATE TABLE `users` (
     `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
     `account_id` VARCHAR(40) NOT NULL,
     `role_id` VARCHAR(40) NOT NULL,
     `name` varchar(255) NOT NULL,
     `email` varchar(255) NOT NULL,
     `phone_number` varchar(255) NOT NULL,
     `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP),
     `updated_at` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
     `deleted_at` datetime
);

CREATE TABLE `categories` (
      `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
      `name` varchar(255) NOT NULL
);

CREATE TABLE `blogs` (
     `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
     `author_id` VARCHAR(40) NOT NULL,
     `title` varchar(255) NOT NULL,
     `thumbnail` text,
     `summary` varchar(255),
     `content` text,
     `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP),
     `updated_at` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
     `deleted_at` datetime
);

CREATE TABLE `blog_categories` (
   `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
   `blog_id` VARCHAR(40) NOT NULL,
   `category_id` VARCHAR(40) NOT NULL
);

CREATE TABLE `comments` (
        `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
        `user_id` varchar(40) NOT NULL,
        `content` text NOT NULL,
        `is_toxicity` bool,
        `created_at` datetime NOT NULL DEFAULT (CURRENT_TIMESTAMP),
        `updated_at` datetime DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
        `deleted_at` datetime
);

CREATE TABLE `blog_comments` (
             `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
             `blog_id` varchar(40) NOT NULL,
             `comment_id` varchar(40) NOT NULL
);

CREATE TABLE `rattings` (
    `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
    `user_id` varchar(40) NOT NULL,
    `level` int NOT NULL
);

CREATE TABLE `blog_rattings` (
     `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
     `blog_id` varchar(40) NOT NULL,
     `ratting_id` varchar(40) NOT NULL
);

ALTER TABLE `users` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `blog_categories` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);

ALTER TABLE `blog_categories` ADD FOREIGN KEY (`blog_id`) REFERENCES `blogs` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `blog_comments` ADD FOREIGN KEY (`blog_id`) REFERENCES `blogs` (`id`);

ALTER TABLE `blog_comments` ADD FOREIGN KEY (`comment_id`) REFERENCES `comments` (`id`);

ALTER TABLE `blog_rattings` ADD FOREIGN KEY (`blog_id`) REFERENCES `blogs` (`id`);

ALTER TABLE `rattings` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `blog_rattings` ADD FOREIGN KEY (`ratting_id`) REFERENCES `rattings` (`id`);

