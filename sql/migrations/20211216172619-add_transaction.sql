
-- +migrate Up
CREATE TABLE transaction(
  id VARCHAR NOT NULL,
  from_account VARCHAR NOT NULL,
  to_account VARCHAR NOT NULL,
  amount INTEGER NOT NULL,
  PRIMARY KEY (id)
);

-- +migrate Down
TRUNCATE account CASCADE;
DROP TABLE account CASCADE;
