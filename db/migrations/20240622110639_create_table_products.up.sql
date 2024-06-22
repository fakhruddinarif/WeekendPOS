CREATE TABLE products (
    id CHAR(36) PRIMARY KEY,
    sku VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    category_id CHAR(36) NOT NULL,
    buy_price DECIMAL(10, 2) NOT NULL,
    sell_price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    photo VARCHAR(255),
    user_id CHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id)
    FOREIGN KEY (user_id) REFERENCES users(id)
);