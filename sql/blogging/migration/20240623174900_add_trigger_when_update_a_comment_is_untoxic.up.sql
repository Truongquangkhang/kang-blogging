CREATE TRIGGER update_nontoxic_comment BEFORE UPDATE ON comments
    FOR EACH ROW
BEGIN
    IF NEW.is_toxicity = FALSE AND OLD.is_toxicity = TRUE THEN
    UPDATE users
    SET is_active = IF(total_violation - 1 > (SELECT value FROM policies WHERE type = "MaxViolation"), FALSE, TRUE),
        total_violation = total_violation - 1,
        expire_warning_time = unix_timestamp()
    WHERE id = NEW.user_id;
END IF;
END;
