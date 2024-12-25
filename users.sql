-- Adminer 4.8.1 MySQL 9.1.0 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `first_name` varchar(20) NOT NULL,
  `last_name` varchar(20) NOT NULL,
  `password_hash` varchar(256) NOT NULL,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `users` (`id`, `first_name`, `last_name`, `password_hash`, `email`) VALUES
('81952a37-3e69-445f-af99-f2007f74cbf0',	'Robin',	'Lambregts',	'IkHashGeenWachtwoorden',	'robin.lambregts@student.uhasselt.be'),
('a9078026-c29d-11ef-b558-0242ac180005',	'Mathias',	'IsGay',	'hafbkhfdsbkhfdshbhbfsakhbhkfsbh',	'mathias@isgay.be'),
('d0633f5b-c29c-11ef-b558-0242ac180005',	'Ebbe',	'Wertz',	'hashed_password',	'ebbe.wertz@student.uhasselt.be');

-- 2024-12-25 10:23:11
