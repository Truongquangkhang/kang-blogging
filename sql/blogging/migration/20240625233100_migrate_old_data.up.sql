INSERT INTO violations (user_id, violation_type, violation_target_id, description)
SELECT user_id, 'comment', id, 'Toxic comment detected.'
FROM comments
WHERE is_toxicity = TRUE;
