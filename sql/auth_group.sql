/*
Navicat MySQL Data Transfer

Source Server         : 林峰会展
Source Server Version : 50740
Source Host           : localhost:3306
Source Database       : linfeng

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2023-12-11 14:40:20
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for auth_group
-- ----------------------------
DROP TABLE IF EXISTS `auth_group`;
CREATE TABLE `auth_group` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父组别',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '组名',
  `rules` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '规则ID',
  `createtime` int(10) DEFAULT NULL COMMENT '创建时间',
  `updatetime` int(10) DEFAULT NULL COMMENT '更新时间',
  `status` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '状态',
  `created` int(11) DEFAULT NULL,
  `updated` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分组表';

-- ----------------------------
-- Records of auth_group
-- ----------------------------
INSERT INTO `auth_group` VALUES ('1', '0', 'Admin group', '*', '1490883540', '149088354', '1', null, null);
INSERT INTO `auth_group` VALUES ('12', '1', '管理组', '1,2,3,4,5,347,6,351,352,0', null, null, '0', '1671438417', null);
INSERT INTO `auth_group` VALUES ('13', '12', '管理二组', '1,2,3,361,0', null, null, '0', '1671438458', null);
