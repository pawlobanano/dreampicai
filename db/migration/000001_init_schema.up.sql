CREATE TABLE account (
  id uuid UNIQUE NOT NULL,
  user_id varchar(50) UNIQUE NOT NULL,
  username varchar(50) UNIQUE NOT NULL,
  created_at timestamp(3) NOT NULL DEFAULT (now() at time zone 'UTC'),
  PRIMARY KEY ("id")
);