# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.6.27-log)
# Database: dcron
# Generation Time: 2018-03-09 06:01:01 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table dcron_crons
# ------------------------------------------------------------

DROP TABLE IF EXISTS `dcron_crons`;

CREATE TABLE `dcron_crons` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `hostname` varchar(100) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
  `expr` varchar(200) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
  `shell` varchar(4000) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
  `comment` varchar(500) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
  `contact` varchar(100) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
  `notify` varchar(300) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
  `notify_mode` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `delete` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `overleap` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_group_v` (`hostname`,`delete`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;



# Dump of table dcron_logs
# ------------------------------------------------------------

DROP TABLE IF EXISTS `dcron_logs`;

CREATE TABLE `dcron_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `cron_id` bigint(20) NOT NULL DEFAULT '0',
  `code` int(11) NOT NULL DEFAULT '0',
  `result` longtext COLLATE utf8mb4_unicode_520_ci,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
