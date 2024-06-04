ALTER TABLE `toxicity_comments`
    DROP CONSTRAINT toxicity_comments_ibfk_1;

DROP TABLE `toxicity_comments`;

ALTER TABLE `users`
    DROP COLUMN `total_violation`,
    DROP COLUMN `expire_warning_time`;

DROP TABLE `policies`;

DROP TRIGGER IF EXISTS insert_toxicity_comment;