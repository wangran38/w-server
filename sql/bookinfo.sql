/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 80012
Source Host           : localhost:3306
Source Database       : lsk

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2024-04-18 14:31:06
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for bookinfo
-- ----------------------------
DROP TABLE IF EXISTS `bookinfo`;
CREATE TABLE `bookinfo` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) DEFAULT NULL COMMENT '姓名',
  `chaptername` varchar(64) DEFAULT NULL COMMENT '章节名称',
  `level` int(3) DEFAULT NULL COMMENT '级别',
  `name` varchar(64) DEFAULT NULL COMMENT '事件名称',
  `content` longtext COMMENT '哪本书',
  `created` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated` int(11) DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of bookinfo
-- ----------------------------
INSERT INTO `bookinfo` VALUES ('1', '1', '第一章', '1', '聊天', '初级', null, null);
INSERT INTO `bookinfo` VALUES ('2', '2', '第二章', '2', '洗澡', '中级', null, null);
INSERT INTO `bookinfo` VALUES ('3', '3', '第三章', '3', '用药', '高级', null, null);
