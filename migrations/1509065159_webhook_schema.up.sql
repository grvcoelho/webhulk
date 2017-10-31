CREATE TABLE IF NOT EXISTS webhooks (
  id varchar(255) PRIMARY KEY,
  name varchar(255),
  url varchar(255) NOT NULL,
  enabled boolean NOT NULL DEFAULT true,
  created_at date NOT NULL,
  updated_at date NOT NULL
);
