-- +goose Up
CREATE TABLE product (
  id SERIAL PRIMARY KEY,
  product_name VARCHAR(50) NOT NULL,
  price VARCHAR(255) NOT NULL,
  create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE product;


