CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    photo VARCHAR(255) NULL DEFAULT NULL,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    phone VARCHAR(16) NULL DEFAULT NULL,
    token VARCHAR(255) NULL DEFAULT NULL,
    role ENUM('owner', 'employee') NOT NULL DEFAULT 'employee',
    user_id CHAR(36) NULL DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);