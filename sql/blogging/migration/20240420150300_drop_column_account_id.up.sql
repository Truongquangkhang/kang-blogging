ALTER TABLE `users`
    DROP CONSTRAINT users_ibfk_1,
    DROP COLUMN account_id,
    ADD FOREIGN KEY (`id`) REFERENCES `accounts` (`id`);
