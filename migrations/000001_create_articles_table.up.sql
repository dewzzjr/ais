CREATE TABLE articles (
    id int NOT NULL AUTO_INCREMENT,
    author text NOT NULL,
    title text NOT NULL,
    body text NOT NULL,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id),
    FULLTEXT (title, body)
);