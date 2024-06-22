ALTER TABLE comments
    ADD COLUMN is_deprecated bool NOT NULL DEFAULT false AFTER deleted_at;
