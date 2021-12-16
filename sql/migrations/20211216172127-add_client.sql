
-- +migrate Up
CREATE TABLE client (
  id VARCHAR NOT NULL,
  PRIMARY KEY (id)
);

-- +migrate Down
TRUNCATE client CASCADE;
DROP TABLE client;
