-- naming relied ON this article: http://citforum.ru/database/articles/naming_rule/

-- users
CREATE TABLE r_user
(
    id     BIGSERIAL PRIMARY KEY,
    gender BOOLEAN  NOT NULL                                 DEFAULT FALSE, -- TRUE for man, FALSE for woman
    age    SMALLINT NOT NULL CHECK (age >= 0 AND age <= 150) DEFAULT 0
);

-- users month expenses
CREATE TABLE r_user_month_expense
(
    user_id  BIGINT REFERENCES r_user (id) ON DELETE CASCADE,
    category TEXT  NOT NULL DEFAULT '',
    amount   FLOAT NOT NULL DEFAULT 0
);
CREATE INDEX idx_r_user_month_expense ON r_user_month_expense (user_id, category);
