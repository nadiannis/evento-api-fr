CREATE TABLE IF NOT EXISTS tickets (
  id BIGSERIAL PRIMARY KEY,
  event_id BIGINT NOT NULL,
  ticket_type_id BIGINT NOT NULL,
  quantity INT NOT NULL DEFAULT 0
);

ALTER TABLE tickets ADD CONSTRAINT tickets_fk_event_id_events_id FOREIGN KEY (event_id) REFERENCES events(id);

ALTER TABLE tickets ADD CONSTRAINT tickets_fk_ticket_type_id_ticket_types_id FOREIGN KEY (ticket_type_id) REFERENCES ticket_types(id);

ALTER TABLE tickets ADD CONSTRAINT tickets_quantity_check CHECK (quantity >= 0);