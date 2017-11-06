CREATE TYPE "deliveries_status" AS ENUM (
  'processing',
  'failed',
  'success'
);

CREATE TABLE IF NOT EXISTS deliveries (
  id varchar(255) PRIMARY KEY,
  message_id varchar(255) references messages(id) NOT NULL,
  status "deliveries_status" DEFAULT 'processing'::"deliveries_status" NOT NULL,
  latency integer,
  status_code varchar(255),
  created_at date NOT NULL,
  updated_at date NOT NULL
);
