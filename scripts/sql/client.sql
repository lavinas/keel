CREATE TABLE keel_client.client (
  id VARCHAR(50) NOT NULL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  nickname VARCHAR(50) NOT NULL,
  document DECIMAL(20) NOT NULL,
  phone DECIMAL(20) NOT NULL,
  email VARCHAR(50) NOT NULL,
  INDEX idx_nick (nickname ASC),
  INDEX idx_document (document ASC),
  INDEX idx_email (email ASC),
  INDEX idx_phone (phone ASC)
);

  