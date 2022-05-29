CREATE TABLE IF NOT EXISTS orders 
(
  id uuid NOT NULL,
  buyer_id uuid NULL,
  seller_id uuid NULL,
  product_id uuid NULL,
  qty int DEFAULT NULL,
  total_price int DEFAULT NULL,
  status int DEFAULT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);