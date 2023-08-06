CREATE TABLE IF NOT EXISTS todos (
    `key` VARCHAR(36),
    `content` VARCHAR(255),
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`key`)
);