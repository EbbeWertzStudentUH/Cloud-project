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
('61a2333f-c154-4f53-8f87-67df85bd6217',	'Ebbe',	'Wertz',	'$2a$12$8uD9z5Kjm.NJKSPArC5cO.1JVJCdJg6MKshD1g.bcDj0a6UhkqlWy',	'ebbe.wertz@student.uhasselt.be'),
('7a2fb106-0ab7-4612-9083-f0ab75b116bc',	'Mathias',	'Houwen',	'$2a$12$rMiC4epCRD7ETCQz7gzIMubB3d3Gl9bY9RuGf3BE/tg1oM8r34/wS',	'mathias.houwen@student.uhasselt.be'),
('901d1dfc-a39d-4c49-9f81-27944841ecf5',	'Robin',	'Lambregts',	'$2a$12$DgX5OxJJUrYFUQtsWb59WulLuzQH44T99SpRpLdSRCvCEwmkyaAKW',	'robin.lambregts@student.uhasselt.be');

-- 2024-12-25 20:09:45
