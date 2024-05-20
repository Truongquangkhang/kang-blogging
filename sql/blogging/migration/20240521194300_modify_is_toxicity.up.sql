ALTER TABLE comments
    MODIFY COLUMN is_toxicity tinyint(1) DEFAULT false;

ALTER TABLE comments
    ADD COLUMN level integer DEFAULT 0,
    ADD COLUMN reply_comment_id varchar(40);

UPDATE comments
    JOIN blog_comments ON comments.id = blog_comments.comment_id
    SET comments.level = blog_comments.level,
        comments.reply_comment_id = blog_comments.reply_comment_id;

ALTER TABLE blog_comments
    DROP COLUMN level,
    DROP COLUMN reply_comment_id;
