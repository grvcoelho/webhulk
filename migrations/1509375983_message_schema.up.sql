CREATE TABLE IF NOT EXISTS messages (
  id varchar(255) PRIMARY KEY,
  webhook_id varchar(255) references webhooks(id) NOT NULL,
  headers json NOT NULL DEFAULT '{}'::json,
  payload json NOT NULL,
  signature varchar(255) NOT NULL,
  created_at date NOT NULL,
  updated_at date NOT NULL
);
