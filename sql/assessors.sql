/*
Navicat MySQL Data Transfer

Source Server         : 3M-老年人
Source Server Version : 50740
Source Host           : 8.134.160.223:3306
Source Database       : oldpp123

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2024-06-15 18:57:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for assessors
-- ----------------------------
DROP TABLE IF EXISTS `assessors`;
CREATE TABLE `assessors` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(200) DEFAULT NULL,
  `nickname` varchar(200) DEFAULT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `email` varchar(200) DEFAULT NULL,
  `salt` varchar(200) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `avatar` text,
  `loginfailure` int(11) DEFAULT NULL,
  `logintime` int(11) DEFAULT NULL,
  `loginip` varchar(200) DEFAULT NULL,
  `token` varchar(59) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of assessors
-- ----------------------------
INSERT INTO `assessors` VALUES ('15', '', '', '13976762545', '', 'jNKH', '0', '', '0', '0', '', '', 'e9f6cea7874931932417c230cebe532d', '2024-04-05 16:20:36', '2024-04-05 16:20:36');
INSERT INTO `assessors` VALUES ('14', '', '', '123451678', '', 'V6Xt', '0', '', '0', '0', '', '', '376281788eb49c061f4acd9306d897ef', '2024-04-05 15:53:29', '2024-04-05 15:53:29');
INSERT INTO `assessors` VALUES ('13', '', '', '12345167', '', 'cbE2', '0', '', '0', '0', '', '', 'f5026719cb5f11343ff90b6ddf197046', '2024-04-05 14:15:57', '2024-04-05 14:15:57');
INSERT INTO `assessors` VALUES ('12', '', '', '1234516', '', 'NULa', '0', '', '0', '0', '', '', 'bdf8347021963d1370c6bd67a7b14054', '2024-04-05 10:16:10', '2024-04-05 10:16:10');
INSERT INTO `assessors` VALUES ('11', '', '', '123451', '', 'nUEe', '0', '', '0', '0', '', '', '1a5a71d078b57fcd6ff6e261ce4d1b55', '2024-04-05 09:57:47', '2024-04-05 09:57:47');
INSERT INTO `assessors` VALUES ('16', '', '', '111', '', 'qlkV', '0', '', '0', '0', '', '', '89c49c3bc1ef6cf874fde294d2681bc0', '2024-04-05 16:24:12', '2024-04-05 16:24:12');
INSERT INTO `assessors` VALUES ('17', '', '', '123', '', 'QEU1', '0', '', '0', '0', '', '', 'a32308a488cfacfe645052ee6e487e84', '2024-04-05 16:26:48', '2024-04-05 16:26:48');
INSERT INTO `assessors` VALUES ('18', '', '', '1234', '', '5Qo7', '0', '', '0', '0', '', '', '313ce1d651abcc18aeb8b5b66c5bf109', '2024-04-05 16:26:56', '2024-04-05 16:26:56');
INSERT INTO `assessors` VALUES ('19', '', '', '12345', '', 'MMJe', '0', '', '0', '0', '', '', 'e2a82a81bde475f2f5de5a5a5fe15c28', '2024-04-05 16:36:51', '2024-04-05 16:36:51');
INSERT INTO `assessors` VALUES ('20', '', '', '123233', '', 'bTlu', '0', '', '0', '0', '', '', 'c64a2c9f25aa00444120fd160b5d1113', '2024-04-05 16:39:44', '2024-04-05 16:39:44');
INSERT INTO `assessors` VALUES ('21', '', '', '159', '', '8ITm', '0', '', '0', '0', '', '', 'be7164ccdc63031fdbd36ab8940cb18a', '2024-04-05 16:48:40', '2024-04-05 16:48:40');
INSERT INTO `assessors` VALUES ('26', '', '', '988', '', 'V9HD', '0', '', '0', '0', '', '', 'fea81ea8a409a3de5a0fd0dadcdf468c', '2024-06-02 23:35:59', '2024-06-02 23:35:59');
INSERT INTO `assessors` VALUES ('25', '', '', '118', '', 'kRTq', '0', '', '0', '0', '', '', 'c5dc6545ed2b82b6d0aa76fa4aea043d', '2024-04-06 09:26:01', '2024-04-06 09:26:01');
