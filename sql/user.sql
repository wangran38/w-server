/*
Navicat MySQL Data Transfer

Source Server         : 本地测试
Source Server Version : 80012
Source Host           : 127.0.0.1:3306
Source Database       : linfeng

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2023-05-11 08:43:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `openid` varchar(200) DEFAULT NULL,
  `nickname` varchar(200) DEFAULT NULL,
  `sex` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否启用 默认1 是 0 无',
  `province` varchar(200) DEFAULT NULL,
  `city` varchar(200) DEFAULT NULL,
  `country` varchar(200) DEFAULT NULL,
  `headimgurl` varchar(200) DEFAULT NULL,
  `unionid` varchar(200) DEFAULT NULL,
  `token` varchar(200) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
