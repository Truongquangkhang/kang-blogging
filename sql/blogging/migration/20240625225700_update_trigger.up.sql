DROP TRIGGER IF EXISTS insert_toxicity_comment;
DROP TRIGGER IF EXISTS update_toxic_comment;
DROP TRIGGER IF EXISTS update_nontoxic_comment;

-- create trigger insert_toxicity_comment
CREATE TRIGGER insert_toxic_comment BEFORE INSERT ON comments
    FOR EACH ROW
BEGIN
    IF NEW.is_toxicity = TRUE THEN
    UPDATE users
    SET is_active = IF(total_violation + 1 > (SELECT value FROM policies WHERE type = "MaxViolation"), FALSE, TRUE),
        total_violation = total_violation + 1,
        expire_warning_time = unix_timestamp() + (total_violation + 1) * (SELECT value FROM policies WHERE type = "ExpireWarningTime")
    WHERE id = NEW.user_id;

    INSERT INTO violations (user_id, violation_type, violation_target_id, description)
    VALUES (NEW.user_id, 'comment', NEW.id, 'Toxic comment detected.');

END IF;
END;

-- create trigger update_toxic_comment
CREATE TRIGGER update_toxic_comment BEFORE UPDATE ON comments
    FOR EACH ROW
BEGIN
    IF NEW.is_toxicity = TRUE AND OLD.is_toxicity = FALSE THEN
    UPDATE users
    SET is_active = IF(total_violation + 1 > (SELECT value FROM policies WHERE type = "MaxViolation"), FALSE, TRUE),
        total_violation = total_violation + 1,
        expire_warning_time = unix_timestamp() + (total_violation + 1) * (SELECT value FROM policies WHERE type = "ExpireWarningTime")
    WHERE id = NEW.user_id;

    INSERT INTO violations (user_id, violation_type, violation_target_id, description)
    VALUES (NEW.user_id, 'comment', NEW.id, 'Toxic comment detected.');

END IF;
END;

-- create trigger update_nontoxic_comment
CREATE TRIGGER update_nontoxic_comment BEFORE UPDATE ON comments
    FOR EACH ROW
BEGIN
    IF NEW.is_toxicity = FALSE AND OLD.is_toxicity = TRUE THEN
    UPDATE users
    SET is_active = IF(total_violation - 1 > (SELECT value FROM policies WHERE type = "MaxViolation"), FALSE, TRUE),
        total_violation = total_violation - 1,
        expire_warning_time = unix_timestamp()
    WHERE id = NEW.user_id;

    DELETE FROM violations WHERE violation_target_id = NEW.id;

END IF;
END;
