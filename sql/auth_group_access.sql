/*
Navicat MySQL Data Transfer

Source Server         : 林峰会展
Source Server Version : 50740
Source Host           : localhost:3306
Source Database       : linfeng

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2023-12-11 14:40:28
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for auth_group_access
-- ----------------------------
DROP TABLE IF EXISTS `auth_group_access`;
CREATE TABLE `auth_group_access` (
  `uid` int(10) unsigned NOT NULL COMMENT '会员ID',
  `gid` int(10) NOT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限分组表';

-- ----------------------------
-- Records of auth_group_access
-- ----------------------------
INSERT INTO `auth_group_access` VALUES ('1', '1');
INSERT INTO `auth_group_access` VALUES ('33', '10');
INSERT INTO `auth_group_access` VALUES ('35', '8');
