DROP TRIGGER IF EXISTS update_toxicity_comment;

CREATE TRIGGER update_toxic_comment BEFORE UPDATE ON comments
    FOR EACH ROW
BEGIN
    IF NEW.is_toxicity = TRUE AND OLD.is_toxicity = FALSE THEN
    UPDATE users
    SET is_active = IF(total_violation + 1 > (SELECT value FROM policies WHERE type = "MaxViolation"), FALSE, TRUE),
        total_violation = total_violation + 1,
        expire_warning_time = unix_timestamp() + (total_violation + 1) * (SELECT value FROM policies WHERE type = "ExpireWarningTime")
    WHERE id = NEW.user_id;
END IF;
END;
