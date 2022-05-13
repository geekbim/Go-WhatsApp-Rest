CREATE TABLE IF NOT EXISTS users 
(
  id bigint NOT NULL,
  name varchar(45) DEFAULT NULL,
  user_name varchar(45) DEFAULT NULL,
  password varchar(225) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by bigint NOT NULL,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_by bigint NOT NULL,
  PRIMARY KEY (id)
);