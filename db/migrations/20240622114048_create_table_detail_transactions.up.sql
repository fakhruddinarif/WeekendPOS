CREATE TABLE detail_transactions(
    id INT PRIMARY KEY AUTO_INCREMENT,
    amount INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    transaction_id CHAR(36) NOT NULL,
    product_id CHAR(36) NOT NULL,
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);