CREATE TABLE IF NOT EXISTS users 
(
  id uuid NOT NULL,
  email varchar(45) DEFAULT NULL,
  name varchar(45) DEFAULT NULL,
  password varchar(225) DEFAULT NULL,
  address varchar(225) DEFAULT NULL,
  role varchar(20) DEFAULT NULL,
  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);