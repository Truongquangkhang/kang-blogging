DELETE FROM blog_categories WHERE id IN (
    SELECT id FROM
        (SELECT
             id,
             ROW_NUMBER() OVER (PARTITION BY blog_id, category_id ORDER BY id) AS row_num
         FROM
             blog_categories
        ) as t
    WHERE t.row_num > 1
);

ALTER TABLE blog_categories
    ADD UNIQUE (blog_id, category_id);
