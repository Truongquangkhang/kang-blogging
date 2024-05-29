ALTER TABLE comments
    ADD COLUMN blog_id varchar(40) AFTER user_id;

UPDATE comments
    JOIN blog_comments ON comments.id = blog_comments.comment_id
    SET comments.blog_id = blog_comments.blog_id;

ALTER TABLE `comments`
    ADD FOREIGN KEY (`blog_id`) REFERENCES `blogs` (`id`);

DROP TABLE blog_comments;