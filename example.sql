-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               5.7.33 - MySQL Community Server (GPL)
-- Server OS:                    Win64
-- HeidiSQL Version:             11.2.0.6213
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for example
CREATE DATABASE IF NOT EXISTS `example` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `example`;

-- Dumping structure for table example.employee
CREATE TABLE IF NOT EXISTS `employee` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `age` bigint(20) NOT NULL DEFAULT '0',
  `address` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table example.employee: ~8 rows (approximately)
/*!40000 ALTER TABLE `employee` DISABLE KEYS */;
INSERT IGNORE INTO `employee` (`id`, `name`, `age`, `address`, `created_at`, `updated_at`) VALUES
	(1, 'Hanif', 31, 'Jalan damai 1', '2023-01-04 08:46:25', '2023-01-04 08:46:25'),
	(2, 'Dani', 21, 'Jalan kediri 1', '2023-01-04 08:46:36', '2023-01-04 08:46:36'),
	(3, 'Ade', 25, 'Jalan Saran 18', '2023-01-04 08:46:43', '2023-01-04 08:46:43'),
	(4, 'wandy', 27, 'Jalan kereta api 2', '2023-01-04 08:46:54', '2023-01-04 08:46:54'),
	(5, 'Annas', 25, 'Jalan kereta api 2', '2023-01-04 09:53:21', '2023-01-04 09:53:21'),
	(6, 'Annas dua', 25, 'Jalan kereta api 2', '2023-01-04 09:57:15', '2023-01-04 09:57:15'),
	(7, 'Annas dua', 25, 'Jalan kereta api 2', '2023-01-04 10:14:44', '2023-01-04 10:14:44'),
	(8, 'Annas dua', 25, 'Jalan kereta api 2', '2023-01-04 11:19:12', '2023-01-04 11:19:12');
/*!40000 ALTER TABLE `employee` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
