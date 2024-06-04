CREATE TABLE `policies` (
    `id` VARCHAR(40) DEFAULT (UUID()) NOT NULL PRIMARY KEY,
    `type` VARCHAR(50),
    `value` integer
);

ALTER TABLE comments
    ADD prediction JSON NULL;

ALTER TABLE `users`
    ADD COLUMN `total_violation` integer NOT NULL DEFAULT 0 AFTER birth_of_day,
    ADD COLUMN `expire_warning_time` integer NULL AFTER total_violation;

-- create trigger
CREATE TRIGGER insert_toxicity_comment BEFORE INSERT ON comments
    FOR EACH ROW
BEGIN
    IF NEW.is_toxicity = TRUE THEN
        UPDATE users
        SET is_active = IF(total_violation + 1 > (SELECT value FROM policies WHERE type = "MaxViolation"), FALSE, TRUE),
            total_violation = total_violation + 1,
            expire_warning_time = unix_timestamp() + (total_violation + 1) * (SELECT value FROM policies WHERE type = "ExpireWarningTime")
        WHERE id = NEW.user_id;
    END IF;
END;


-- insert data default for table policies
INSERT INTO policies (type, value) VALUES
("MaxViolation", 5),
("ExpireWarningTime", 900);
