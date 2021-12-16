
-- +migrate Up
CREATE TABLE account(
  id VARCHAR NOT NULL,
  clientId VARCHAR NOT NULL REFERENCES client(id),
  currency VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

-- +migrate Down
TRUNCATE account CASCADE;
DROP TABLE account CASCADE;
