CREATE TABLE IF NOT EXISTS products 
(
  id uuid NOT NULL,
  name varchar(50) DEFAULT NULL,
  description text DEFAULT NULL,
  price int DEFAULT NULL,
  seller_id uuid NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);