ALTER TABLE blog_comments
    ADD COLUMN level integer DEFAULT 0,
    ADD COLUMN reply_comment_id varchar(40);
