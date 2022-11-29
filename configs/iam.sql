CREATE DATABASE  IF NOT EXISTS `iam` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `iam`;

--
-- Table structure for table `policy`
--

DROP TABLE IF EXISTS `policy`;

CREATE TABLE `policy` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `instanceID` varchar(32) DEFAULT NULL,
  `name` varchar(45) NOT NULL,
  `username` varchar(255) NOT NULL,
  `policyShadow` longtext DEFAULT NULL,
  `extendShadow` longtext DEFAULT NULL,
  `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
  `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `instanceID_UNIQUE` (`instanceID`),
  KEY `fk_policy_user_idx` (`username`),
  CONSTRAINT `fk_policy_user` FOREIGN KEY (`username`) REFERENCES `user` (`name`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `policy`
--

LOCK TABLES `policy` WRITE;
UNLOCK TABLES;
DELIMITER ;;
DELIMITER ;

--
-- Table structure for table `policy_audit`
--

DROP TABLE IF EXISTS `policy_audit`;

CREATE TABLE `policy_audit` (
  `id` bigint(20) unsigned NOT NULL,
  `instanceID` varchar(32) DEFAULT NULL,
  `name` varchar(45) NOT NULL,
  `username` varchar(255) NOT NULL,
  `policyShadow` longtext DEFAULT NULL,
  `extendShadow` longtext DEFAULT NULL,
  `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
  `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deletedAt` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  KEY `fk_policy_user_idx` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `policy_audit`
--

LOCK TABLES `policy_audit` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `secret`
--

DROP TABLE IF EXISTS `secret`;
CREATE TABLE `secret` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `instanceID` varchar(32) DEFAULT NULL,
  `name` varchar(45) NOT NULL,
  `username` varchar(255) NOT NULL,
  `secretID` varchar(36) NOT NULL,
  `secretKey` varchar(255) NOT NULL,
  `expires` int(64) unsigned NOT NULL DEFAULT 1534308590,
  `description` varchar(255) NOT NULL,
  `extendShadow` longtext DEFAULT NULL,
  `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
  `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `instanceID_UNIQUE` (`instanceID`),
  KEY `fk_secret_user_idx` (`username`),
  CONSTRAINT `fk_secret_user` FOREIGN KEY (`username`) REFERENCES `user` (`name`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `secret`
--

LOCK TABLES `secret` WRITE;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `instanceID` varchar(32) DEFAULT NULL,
  `name` varchar(45) NOT NULL,
  `status` int(1) DEFAULT 1 COMMENT '1:可用，0:不可用',
  `nickname` varchar(30) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(256) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `isAdmin` tinyint(1) unsigned NOT NULL DEFAULT 0 COMMENT '1: administrator\\\\n0: non-administrator',
  `extendShadow` longtext DEFAULT NULL,
  `loginedAt` timestamp NULL DEFAULT NULL COMMENT 'last login time',
  `createdAt` timestamp NOT NULL DEFAULT current_timestamp(),
  `updatedAt` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  UNIQUE KEY `instanceID_UNIQUE` (`instanceID`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
INSERT INTO `user` VALUES (1,'user-lingfei','admin',1,'admin','$2a$10$WnQD2DCfWVhlGmkQ8pdLkesIGPf9KJB7N1mhSOqulbgN7ZMo44Mv2','admin@foxmail.com','1812884xxxx',1,'{}',now(),'2021-05-27 10:01:40','2021-05-05 21:13:14');
UNLOCK TABLES;
DELIMITER ;;
DELIMITER ;