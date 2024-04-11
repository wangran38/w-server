/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 80012
Source Host           : localhost:3306
Source Database       : oldpp

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2024-04-06 09:36:54
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
DROP TABLE IF EXISTS `dictionary`;
CREATE TABLE `dictionary` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` int(255) DEFAULT NULL,
  `codename` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=35 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of dictionary
-- ----------------------------
INSERT INTO `dictionary` VALUES ('1', '1', 'nation', '汉族');
INSERT INTO `dictionary` VALUES ('2', '1', 'nation', '少数民族');
INSERT INTO `dictionary` VALUES ('3', '2', 'marital', '离婚');
INSERT INTO `dictionary` VALUES ('4', '2', 'marital', '未婚');
INSERT INTO `dictionary` VALUES ('5', '2', 'marital', '离异');
INSERT INTO `dictionary` VALUES ('6', '3', 'payment', '独居');
INSERT INTO `dictionary` VALUES ('7', '3', 'payment', '与配偶居住');
INSERT INTO `dictionary` VALUES ('8', '3', 'payment', '与子女居住');
INSERT INTO `dictionary` VALUES ('9', '3', 'payment', '与父母居住');
INSERT INTO `dictionary` VALUES ('10', '3', 'payment', '与兄弟姐妹居住');
INSERT INTO `dictionary` VALUES ('11', '3', 'payment', '与其他亲属居住');
INSERT INTO `dictionary` VALUES ('12', '3', 'payment', '与非亲属关系的居住养老机构');
INSERT INTO `dictionary` VALUES ('13', '3', 'payment', '养老机构');
INSERT INTO `dictionary` VALUES ('14', '4', 'cost', '城镇职工基本医疗保险');
INSERT INTO `dictionary` VALUES ('15', '4', 'cost', '城镇居民基本医疗保险');
INSERT INTO `dictionary` VALUES ('16', '4', 'cost', '自费');
INSERT INTO `dictionary` VALUES ('17', '4', 'cost', '公务员补助');
INSERT INTO `dictionary` VALUES ('18', '4', 'cost', '企业补充保险');
INSERT INTO `dictionary` VALUES ('19', '4', 'cost', '公费医疗及医疗照顾对象');
INSERT INTO `dictionary` VALUES ('20', '4', 'cost', '医疗救助');
INSERT INTO `dictionary` VALUES ('21', '4', 'cost', '大病保险');
INSERT INTO `dictionary` VALUES ('22', '5', 'title', '小学');
INSERT INTO `dictionary` VALUES ('23', '5', 'title', '初中');
INSERT INTO `dictionary` VALUES ('24', '5', 'title', '高中');
INSERT INTO `dictionary` VALUES ('25', '5', 'title', '大专');
INSERT INTO `dictionary` VALUES ('26', '5', 'title', '本科');
INSERT INTO `dictionary` VALUES ('27', '6', 'religious', '有');
INSERT INTO `dictionary` VALUES ('28', '6', 'religious', '无');
INSERT INTO `dictionary` VALUES ('29', '7', 'economic', '退休金/养老金');
INSERT INTO `dictionary` VALUES ('30', '7', 'economic', '子女补贴');
INSERT INTO `dictionary` VALUES ('31', '7', 'economic', '亲友资助');
INSERT INTO `dictionary` VALUES ('32', '7', 'economic', '国家普惠型补贴');
INSERT INTO `dictionary` VALUES ('33', '7', 'economic', '个人储蓄');
INSERT INTO `dictionary` VALUES ('34', '7', 'economic', '其他补贴');
