/*
 Navicat Premium Data Transfer

 Source Server         : local_root
 Source Server Type    : MySQL
 Source Server Version : 80015
 Source Host           : localhost
 Source Database       : test

 Target Server Type    : MySQL
 Target Server Version : 80015
 File Encoding         : utf-8

 Date: 07/31/2019 17:09:28 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `jdcron_tasks`
-- ----------------------------
DROP TABLE IF EXISTS `jdcron_tasks`;
CREATE TABLE `jdcron_tasks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_id` int(11) NOT NULL DEFAULT '0',
  `spec` varchar(255) DEFAULT NULL,
  `command` varchar(255) DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `handler` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
