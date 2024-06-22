CREATE TABLE transactions (
    id CHAR(36) PRIMARY KEY,
    customer VARCHAR(255) NOT NULL,
    date DATETIME NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    user_id CHAR(36) NOT NULL,
    employee_id CHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (employee_id) REFERENCES employees(id)
);