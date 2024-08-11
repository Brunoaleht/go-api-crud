-- Active: 1722212882979@@127.0.0.1@5432@postgres
CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(50) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);


INSERT INTO product (product_name, price) VALUES ('Apple', 1.00);


SELECT * FROM product;