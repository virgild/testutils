CREATE TABLE users
(
    id         varchar(128) NOT NULL,
    email      varchar(128) NOT NULL,
    created_at datetime     NOT NULL,
    updated_at datetime     NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY users_email_uindex (email),
    UNIQUE KEY users_id_uindex (id)
) ENGINE = InnoDB
DEFAULT CHARSET = utf8mb4;
