-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    password VARCHAR(255) NOT NULL COMMENT 'パスワード',
    image VARCHAR(255) NULL COMMENT 'プロフィール画像',
    status TINYINT UNSIGNED NOT NULL COMMENT 'ステータス',
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    PRIMARY KEY(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
