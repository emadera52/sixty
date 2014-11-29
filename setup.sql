CREATE DATABASE `sixtyplus` /*!40100 COLLATE 'utf8_general_ci' */

CREATE TABLE `user` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(30) NOT NULL,
	`fullname` VARCHAR(60) NOT NULL,
	`email` VARCHAR(256) NOT NULL,
	`password` VARCHAR(128) NOT NULL,
	`auto_log` TINYINT(1) NOT NULL,
	`interest` TINYINT(2) NOT NULL,
	`reg_key` VARCHAR(60) NOT NULL,
	`reset_key` VARCHAR(60) NOT NULL,
	`created` DATETIME NOT NULL,
	`updated` DATETIME NOT NULL,
	PRIMARY KEY (`id`),
	UNIQUE INDEX `username` (`username`)
)
COMMENT='email is encrypted, password is hashed'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1000;

CREATE TABLE `comment` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`user_id` INT(11) NOT NULL,
	`txt` VARCHAR(4096) NOT NULL,
	`category` TINYINT(2) NOT NULL,
	`archived` TINYINT(1) NOT NULL,
	`created` DATETIME NOT NULL,
	PRIMARY KEY (`id`),
	INDEX `FK__user` (`user_id`),
	CONSTRAINT `FKuser` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
)
COMMENT='category indicates the comment page'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1000;
