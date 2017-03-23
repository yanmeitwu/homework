CREATE DATABASE IF NOT EXISTS `convoy` COLLATE=utf8_unicode_ci;

# create table drivers
CREATE TABLE IF NOT EXISTS `drivers` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `capacity` int(11) NOT NULL,
  `received_offers` int(11) NOT NULL DEFAULT '0',
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted` datetime DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

# create table shipment
CREATE TABLE IF NOT EXISTS `shipments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `weight` int(11) NOT NULL,
  `accepted` smallint(1) NOT NULL DEFAULT '0',
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `modified` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted` datetime DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

# create table offers
CREATE TABLE IF NOT EXISTS `offers`
(
  `id` INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `shipment_id` int NOT NULL,
  `driver_id` int NOT NULL,
  `status` VARCHAR(10),
  `created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `modified` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted` DATETIME DEFAULT 0,
  UNIQUE (`shipment_id`,`driver_id`)
) ENGINE=INNODB CHARACTER SET utf8;
