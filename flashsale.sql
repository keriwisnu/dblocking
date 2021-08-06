CREATE TABLE `flashsale` (
	`id` Int( 8 ) AUTO_INCREMENT NOT NULL,
	`name` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
	`stock` BigInt( 14 ) NOT NULL,
	CONSTRAINT `unique_id` UNIQUE( `id` ) )
CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB
AUTO_INCREMENT = 9;

INSERT INTO `flashsale`(`id`,`name`,`stock`) VALUES
( '1', 'Kulkas', '100'),
( '2', 'Brankas', '10');