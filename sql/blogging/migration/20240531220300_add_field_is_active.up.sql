ALTER TABLE users
    ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT true AFTER birth_of_day;

ALTER TABLE blogs
    ADD COLUMN published BOOLEAN NOT NULL DEFAULT false AFTER content,
    ADD COLUMN is_deprecated BOOLEAN NOT NULL DEFAULT false AFTER published;

UPDATE blogs SET published = true;
