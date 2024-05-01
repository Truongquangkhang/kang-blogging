ALTER TABLE users
    MODIFY phone_number varchar(255) NULL,
    ADD display_name varchar(255) NULL AFTER phone_number,
    ADD avatar varchar(255) NULL AFTER display_name,
    ADD gender BOOLEAN DEFAULT true after avatar,
    ADD birth_of_day BIGINT NULL after gender;
