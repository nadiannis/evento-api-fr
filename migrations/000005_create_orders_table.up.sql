CREATE TABLE IF NOT EXISTS orders (
  id BIGSERIAL PRIMARY KEY,
  customer_id BIGINT NOT NULL,
  ticket_id BIGINT NOT NULL,
  quantity INT NOT NULL DEFAULT 0,
  total_price NUMERIC NOT NULL DEFAULT 0,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

ALTER TABLE orders ADD CONSTRAINT orders_fk_customer_id_customers_id FOREIGN KEY (customer_id) REFERENCES customers(id);

ALTER TABLE orders ADD CONSTRAINT orders_fk_ticket_id_tickets_id FOREIGN KEY (ticket_id) REFERENCES tickets(id);
