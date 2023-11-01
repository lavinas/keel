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

insert into keel_client.client (id, name, nickname, document, phone, email) values ('0', 'business ltda', 'business', 5072212000106, 5511999999999, 'business@test.com');
insert into keel_client.client (id, name, nickname, document, phone, email) values ('1', 'Consumer Doe', 'consumer_doe', 82109271086, 5511999999998, 'consumer@test.com');