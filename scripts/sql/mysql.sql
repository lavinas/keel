CREATE SCHEMA `cbs_client` ;

CREATE TABLE `cbs_client`.`client` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `nickname` VARCHAR(50) NOT NULL,
  `document` DECIMAL(20) NOT NULL,
  `phone` DECIMAL(20) NOT NULL,
  `email` VARCHAR(50) NOT NULL,
  PRIMARY KEY (`id`));

ALTER TABLE `cbs_client`.`client` 
  ADD INDEX `idx1` (`nickname` ASC) VISIBLE;
;

