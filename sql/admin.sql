/*
Navicat MySQL Data Transfer

Source Server         : 林峰会展
Source Server Version : 50740
Source Host           : localhost:3306
Source Database       : linfeng

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2023-12-11 14:54:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(200) DEFAULT NULL,
  `nickname` varchar(200) DEFAULT NULL,
  `salt` varchar(200) DEFAULT NULL,
  `age` int(2) DEFAULT NULL,
  `avatar` text,
  `loginfailure` int(10) DEFAULT NULL,
  `logintime` int(10) DEFAULT NULL,
  `loginip` varchar(200) DEFAULT NULL,
  `token` varchar(59) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `phone` varchar(200) DEFAULT NULL,
  `email` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=36 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES ('1', 'feng', '峰', 'MRq2', null, 'favicon.ico', null, null, null, null, '052931db8810edf6dc84d8c9a908e0d6', null, '2023-04-23 18:40:25', '189076938272', '707229207@qq.com');
INSERT INTO `admin` VALUES ('33', 'ces2', '测试2', 'OO6T', '0', 'http://file.988cj.com/group1/default/20221216/15/45/3/20201030114158.png', '0', '0', '', '', '8466e5e44de8f8bb4f8a06fac522973a', '2022-12-16 17:24:45', '2022-12-19 10:39:48', '18907693872', '18907693872@qq.com');
INSERT INTO `admin` VALUES ('35', 'wangran', '111', 'sryP', '0', 'http://file.988cj.com/group1/default/20221216/17/17/3/微信图片_20221117152257.jpg', '0', '0', '', '', 'a41ff83f1eaf0e624eb0580ee6170a55', '2022-12-19 10:32:32', '2022-12-19 10:32:32', '1212333423', '23232323');
