# CLIENTS

CREATE TABLE keel_client.client (
  id VARCHAR(50) NOT NULL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  nickname VARCHAR(50) NOT NULL,
  document DECIMAL(20) NOT NULL,
  phone DECIMAL(20) NOT NULL,
  email VARCHAR(50) NOT NULL,
  PRIMARY KEY (id)
);

ALTER TABLE keel_client.client 
  ADD INDEX idx_nick (nickname ASC);

ALTER TABLE keel_client.client 
  ADD INDEX idx2 (document ASC);

ALTER TABLE keel_client.client 
  ADD INDEX idx3 (email ASC);

ALTER TABLE keel_client.client 
  ADD INDEX idx4 (phone ASC);

  