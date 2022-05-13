  CREATE TABLE IF NOT EXISTS transactions (
    id bigint NOT NULL,
    merchant_id bigint NOT NULL,
    outlet_id bigint NOT NULL,
    bill_total double precision NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by bigint NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by bigint NOT NULL,
    PRIMARY KEY (id)
  );