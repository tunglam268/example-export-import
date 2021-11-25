CREATE TABLE `account`.`accounts` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NULL ,
    `address` VARCHAR(255) NULL,
    `phonenumber` VARCHAR(255) NOT NULL ,
    `balance` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);