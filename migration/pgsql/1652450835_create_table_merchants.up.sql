CREATE TABLE IF NOT EXISTS merchants 
(
    id bigint NOT NULL,
    user_id int NOT NULL,
    merchant_name varchar(40) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by bigint NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by bigint NOT NULL,
    PRIMARY KEY (id)
);