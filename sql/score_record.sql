/*
Navicat MySQL Data Transfer

Source Server         : 3M-老年人
Source Server Version : 50740
Source Host           : 8.134.160.223:3306
Source Database       : oldpp123

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2024-06-15 18:32:25
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for score_record
-- ----------------------------
DROP TABLE IF EXISTS `score_record`;
CREATE TABLE `score_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `senior_id` bigint(20) DEFAULT NULL COMMENT '老者id',
  `assessors_id` bigint(20) DEFAULT NULL COMMENT '所属评估员id',
  `number_id` bigint(20) DEFAULT NULL COMMENT '编号id',
  `kpi_id` bigint(20) DEFAULT NULL COMMENT '指标id',
  `kpiinfo_id` bigint(20) DEFAULT NULL COMMENT '指标分题id',
  `score` int(11) DEFAULT NULL COMMENT '分值',
  `created` int(11) DEFAULT NULL,
  `updated` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=168 DEFAULT CHARSET=utf8 ROW_FORMAT=FIXED;

-- ----------------------------
-- Records of score_record
-- ----------------------------
INSERT INTO `score_record` VALUES ('1', '0', '21', '0', '0', '0', '3', '1715932114', null);
INSERT INTO `score_record` VALUES ('2', '0', '21', '0', '0', '0', '3', '1715932114', null);
INSERT INTO `score_record` VALUES ('3', '0', '21', '0', '0', '0', '3', '1715932114', null);
INSERT INTO `score_record` VALUES ('4', '0', '21', '0', '0', '0', '0', '1715932114', null);
INSERT INTO `score_record` VALUES ('5', '0', '21', '0', '0', '0', '0', '1715932114', null);
INSERT INTO `score_record` VALUES ('6', '0', '21', '0', '0', '0', '0', '1715932114', null);
INSERT INTO `score_record` VALUES ('7', '0', '21', '0', '0', '0', '0', '1715932114', null);
INSERT INTO `score_record` VALUES ('8', '0', '21', '49', '0', '0', '4', '1715934152', null);
INSERT INTO `score_record` VALUES ('9', '0', '21', '49', '0', '0', '3', '1715934152', null);
INSERT INTO `score_record` VALUES ('10', '0', '21', '49', '0', '0', '3', '1715934152', null);
INSERT INTO `score_record` VALUES ('11', '0', '21', '49', '0', '0', '0', '1715934152', null);
INSERT INTO `score_record` VALUES ('12', '0', '21', '49', '0', '0', '0', '1715934152', null);
INSERT INTO `score_record` VALUES ('13', '0', '21', '49', '0', '0', '0', '1715934152', null);
INSERT INTO `score_record` VALUES ('14', '0', '21', '49', '12', '18', '0', '1715934152', null);
INSERT INTO `score_record` VALUES ('15', '0', '21', '49', '12', '18', '4', '1715934729', null);
INSERT INTO `score_record` VALUES ('16', '0', '21', '49', '0', '0', '3', '1715934729', null);
INSERT INTO `score_record` VALUES ('17', '0', '21', '49', '0', '0', '3', '1715934729', null);
INSERT INTO `score_record` VALUES ('18', '0', '21', '49', '0', '0', '0', '1715934729', null);
INSERT INTO `score_record` VALUES ('19', '0', '21', '49', '0', '0', '0', '1715934729', null);
INSERT INTO `score_record` VALUES ('20', '0', '21', '49', '0', '0', '0', '1715934729', null);
INSERT INTO `score_record` VALUES ('21', '0', '21', '49', '0', '0', '0', '1715934729', null);
INSERT INTO `score_record` VALUES ('22', '0', '21', '49', '0', '0', '4', '1715934781', null);
INSERT INTO `score_record` VALUES ('23', '0', '21', '49', '0', '0', '3', '1715934781', null);
INSERT INTO `score_record` VALUES ('24', '0', '21', '49', '0', '0', '3', '1715934781', null);
INSERT INTO `score_record` VALUES ('25', '0', '21', '49', '0', '0', '0', '1715934781', null);
INSERT INTO `score_record` VALUES ('26', '0', '21', '49', '0', '0', '0', '1715934781', null);
INSERT INTO `score_record` VALUES ('27', '0', '21', '49', '0', '0', '0', '1715934781', null);
INSERT INTO `score_record` VALUES ('28', '0', '21', '49', '0', '0', '0', '1715934781', null);
INSERT INTO `score_record` VALUES ('29', '4', '21', '52', '0', '0', '10', '1716951878', null);
INSERT INTO `score_record` VALUES ('30', '4', '21', '52', '0', '0', '5', '1716951878', null);
INSERT INTO `score_record` VALUES ('31', '4', '21', '52', '0', '0', '10', '1716951878', null);
INSERT INTO `score_record` VALUES ('32', '4', '21', '52', '0', '0', '5', '1716951878', null);
INSERT INTO `score_record` VALUES ('33', '4', '21', '52', '0', '0', '5', '1716951878', null);
INSERT INTO `score_record` VALUES ('34', '4', '21', '52', '0', '0', '10', '1716951878', null);
INSERT INTO `score_record` VALUES ('35', '4', '21', '52', '0', '0', '10', '1716951878', null);
INSERT INTO `score_record` VALUES ('36', '4', '21', '52', '0', '0', '10', '1716951878', null);
INSERT INTO `score_record` VALUES ('37', '4', '21', '52', '0', '0', '10', '1716951878', null);
INSERT INTO `score_record` VALUES ('38', '4', '21', '52', '0', '0', '5', '1716951878', null);
INSERT INTO `score_record` VALUES ('39', '5', '21', '53', '0', '0', '4', '1716952298', null);
INSERT INTO `score_record` VALUES ('40', '5', '21', '53', '0', '0', '3', '1716952298', null);
INSERT INTO `score_record` VALUES ('41', '5', '21', '53', '0', '0', '3', '1716952298', null);
INSERT INTO `score_record` VALUES ('42', '5', '21', '53', '0', '0', '0', '1716952298', null);
INSERT INTO `score_record` VALUES ('43', '5', '21', '53', '0', '0', '0', '1716952298', null);
INSERT INTO `score_record` VALUES ('44', '5', '21', '53', '0', '0', '0', '1716952298', null);
INSERT INTO `score_record` VALUES ('45', '5', '21', '53', '0', '0', '0', '1716952298', null);
INSERT INTO `score_record` VALUES ('46', '6', '21', '53', '0', '0', '4', '1716952754', null);
INSERT INTO `score_record` VALUES ('47', '6', '21', '53', '0', '0', '3', '1716952754', null);
INSERT INTO `score_record` VALUES ('48', '6', '21', '53', '0', '0', '2', '1716952754', null);
INSERT INTO `score_record` VALUES ('49', '6', '21', '53', '0', '0', '0', '1716952754', null);
INSERT INTO `score_record` VALUES ('50', '6', '21', '53', '0', '0', '0', '1716952754', null);
INSERT INTO `score_record` VALUES ('51', '6', '21', '53', '0', '0', '0', '1716952754', null);
INSERT INTO `score_record` VALUES ('52', '6', '21', '53', '0', '0', '0', '1716952754', null);
INSERT INTO `score_record` VALUES ('53', '6', '21', '53', '0', '0', '3', '1716953348', null);
INSERT INTO `score_record` VALUES ('54', '6', '21', '53', '0', '0', '2', '1716953348', null);
INSERT INTO `score_record` VALUES ('55', '6', '21', '53', '0', '0', '4', '1716953348', null);
INSERT INTO `score_record` VALUES ('56', '6', '21', '53', '0', '0', '0', '1716953348', null);
INSERT INTO `score_record` VALUES ('57', '6', '21', '53', '0', '0', '0', '1716953348', null);
INSERT INTO `score_record` VALUES ('58', '6', '21', '53', '0', '0', '0', '1716953348', null);
INSERT INTO `score_record` VALUES ('59', '6', '21', '53', '0', '0', '0', '1716953348', null);
INSERT INTO `score_record` VALUES ('60', '6', '21', '53', '0', '0', '3', '1716953559', null);
INSERT INTO `score_record` VALUES ('61', '6', '21', '53', '0', '0', '3', '1716953559', null);
INSERT INTO `score_record` VALUES ('62', '6', '21', '53', '0', '0', '2', '1716953559', null);
INSERT INTO `score_record` VALUES ('63', '6', '21', '53', '0', '0', '0', '1716953559', null);
INSERT INTO `score_record` VALUES ('64', '6', '21', '53', '0', '0', '0', '1716953559', null);
INSERT INTO `score_record` VALUES ('65', '6', '21', '53', '0', '0', '0', '1716953559', null);
INSERT INTO `score_record` VALUES ('66', '6', '21', '53', '0', '0', '0', '1716953559', null);
INSERT INTO `score_record` VALUES ('67', '6', '21', '53', '0', '0', '3', '1716954316', null);
INSERT INTO `score_record` VALUES ('68', '6', '21', '53', '0', '0', '3', '1716954316', null);
INSERT INTO `score_record` VALUES ('69', '6', '21', '53', '0', '0', '3', '1716954316', null);
INSERT INTO `score_record` VALUES ('70', '6', '21', '53', '0', '0', '0', '1716954316', null);
INSERT INTO `score_record` VALUES ('71', '6', '21', '53', '0', '0', '0', '1716954316', null);
INSERT INTO `score_record` VALUES ('72', '6', '21', '53', '0', '0', '0', '1716954316', null);
INSERT INTO `score_record` VALUES ('73', '6', '21', '53', '0', '0', '0', '1716954316', null);
INSERT INTO `score_record` VALUES ('74', '6', '21', '53', '0', '0', '3', '1716964025', null);
INSERT INTO `score_record` VALUES ('75', '6', '21', '53', '0', '0', '2', '1716964025', null);
INSERT INTO `score_record` VALUES ('76', '6', '21', '53', '0', '0', '3', '1716964025', null);
INSERT INTO `score_record` VALUES ('77', '6', '21', '53', '0', '0', '0', '1716964025', null);
INSERT INTO `score_record` VALUES ('78', '6', '21', '53', '0', '0', '0', '1716964025', null);
INSERT INTO `score_record` VALUES ('79', '6', '21', '53', '0', '0', '0', '1716964025', null);
INSERT INTO `score_record` VALUES ('80', '6', '21', '53', '0', '0', '0', '1716964025', null);
INSERT INTO `score_record` VALUES ('81', '6', '21', '53', '0', '0', '1', '1716964330', null);
INSERT INTO `score_record` VALUES ('82', '6', '21', '53', '0', '0', '2', '1716964330', null);
INSERT INTO `score_record` VALUES ('83', '6', '21', '53', '0', '0', '2', '1716964330', null);
INSERT INTO `score_record` VALUES ('84', '6', '21', '53', '0', '0', '0', '1716964330', null);
INSERT INTO `score_record` VALUES ('85', '6', '21', '53', '0', '0', '0', '1716964330', null);
INSERT INTO `score_record` VALUES ('86', '6', '21', '53', '0', '0', '0', '1716964330', null);
INSERT INTO `score_record` VALUES ('87', '6', '21', '53', '0', '0', '0', '1716964330', null);
INSERT INTO `score_record` VALUES ('88', '6', '21', '53', '0', '0', '3', '1716967902', null);
INSERT INTO `score_record` VALUES ('89', '6', '21', '53', '0', '0', '2', '1716967902', null);
INSERT INTO `score_record` VALUES ('90', '6', '21', '53', '0', '0', '2', '1716967902', null);
INSERT INTO `score_record` VALUES ('91', '6', '21', '53', '0', '0', '0', '1716967902', null);
INSERT INTO `score_record` VALUES ('92', '6', '21', '53', '0', '0', '0', '1716967902', null);
INSERT INTO `score_record` VALUES ('93', '6', '21', '53', '0', '0', '0', '1716967902', null);
INSERT INTO `score_record` VALUES ('94', '6', '21', '53', '0', '0', '0', '1716967902', null);
INSERT INTO `score_record` VALUES ('95', '6', '21', '53', '0', '0', '3', '1716968128', null);
INSERT INTO `score_record` VALUES ('96', '6', '21', '53', '0', '0', '3', '1716968128', null);
INSERT INTO `score_record` VALUES ('97', '6', '21', '53', '0', '0', '3', '1716968128', null);
INSERT INTO `score_record` VALUES ('98', '6', '21', '53', '0', '0', '0', '1716968128', null);
INSERT INTO `score_record` VALUES ('99', '6', '21', '53', '0', '0', '0', '1716968128', null);
INSERT INTO `score_record` VALUES ('100', '6', '21', '53', '0', '0', '0', '1716968128', null);
INSERT INTO `score_record` VALUES ('101', '6', '21', '53', '0', '0', '0', '1716968128', null);
INSERT INTO `score_record` VALUES ('102', '6', '21', '53', '0', '0', '2', '1716975829', null);
INSERT INTO `score_record` VALUES ('103', '6', '21', '53', '0', '0', '3', '1716975829', null);
INSERT INTO `score_record` VALUES ('104', '6', '21', '53', '0', '0', '3', '1716975829', null);
INSERT INTO `score_record` VALUES ('105', '6', '21', '53', '0', '0', '0', '1716975829', null);
INSERT INTO `score_record` VALUES ('106', '6', '21', '53', '0', '0', '0', '1716975829', null);
INSERT INTO `score_record` VALUES ('107', '6', '21', '53', '0', '0', '0', '1716975829', null);
INSERT INTO `score_record` VALUES ('108', '6', '21', '53', '0', '0', '0', '1716975829', null);
INSERT INTO `score_record` VALUES ('109', '6', '21', '53', '0', '0', '3', '1717028436', null);
INSERT INTO `score_record` VALUES ('110', '6', '21', '53', '0', '0', '3', '1717028436', null);
INSERT INTO `score_record` VALUES ('111', '6', '21', '53', '0', '0', '2', '1717028436', null);
INSERT INTO `score_record` VALUES ('112', '6', '21', '53', '0', '0', '0', '1717028436', null);
INSERT INTO `score_record` VALUES ('113', '6', '21', '53', '0', '0', '0', '1717028436', null);
INSERT INTO `score_record` VALUES ('114', '6', '21', '53', '0', '0', '0', '1717028436', null);
INSERT INTO `score_record` VALUES ('115', '6', '21', '53', '0', '0', '0', '1717028436', null);
INSERT INTO `score_record` VALUES ('116', '6', '21', '53', '0', '0', '3', '1717028678', null);
INSERT INTO `score_record` VALUES ('117', '6', '21', '53', '0', '0', '3', '1717028678', null);
INSERT INTO `score_record` VALUES ('118', '6', '21', '53', '0', '0', '3', '1717028678', null);
INSERT INTO `score_record` VALUES ('119', '6', '21', '53', '0', '0', '0', '1717028678', null);
INSERT INTO `score_record` VALUES ('120', '6', '21', '53', '0', '0', '0', '1717028678', null);
INSERT INTO `score_record` VALUES ('121', '6', '21', '53', '0', '0', '0', '1717028678', null);
INSERT INTO `score_record` VALUES ('122', '6', '21', '53', '0', '0', '0', '1717028678', null);
INSERT INTO `score_record` VALUES ('123', '6', '21', '53', '0', '0', '3', '1717029248', null);
INSERT INTO `score_record` VALUES ('124', '6', '21', '53', '0', '0', '3', '1717029248', null);
INSERT INTO `score_record` VALUES ('125', '6', '21', '53', '0', '0', '3', '1717029248', null);
INSERT INTO `score_record` VALUES ('126', '6', '21', '53', '0', '0', '0', '1717029248', null);
INSERT INTO `score_record` VALUES ('127', '6', '21', '53', '0', '0', '0', '1717029248', null);
INSERT INTO `score_record` VALUES ('128', '6', '21', '53', '0', '0', '0', '1717029248', null);
INSERT INTO `score_record` VALUES ('129', '6', '21', '53', '0', '0', '0', '1717029248', null);
INSERT INTO `score_record` VALUES ('130', '6', '21', '53', '0', '0', '3', '1717031683', null);
INSERT INTO `score_record` VALUES ('131', '6', '21', '53', '0', '0', '3', '1717031683', null);
INSERT INTO `score_record` VALUES ('132', '6', '21', '53', '0', '0', '4', '1717031683', null);
INSERT INTO `score_record` VALUES ('133', '6', '21', '53', '0', '0', '0', '1717031683', null);
INSERT INTO `score_record` VALUES ('134', '6', '21', '53', '0', '0', '0', '1717031683', null);
INSERT INTO `score_record` VALUES ('135', '6', '21', '53', '0', '0', '0', '1717031683', null);
INSERT INTO `score_record` VALUES ('136', '6', '21', '53', '0', '0', '0', '1717031683', null);
INSERT INTO `score_record` VALUES ('137', '6', '21', '53', '0', '0', '2', '1717031712', null);
INSERT INTO `score_record` VALUES ('138', '6', '21', '53', '0', '0', '1', '1717031712', null);
INSERT INTO `score_record` VALUES ('139', '6', '21', '53', '0', '0', '3', '1717031712', null);
INSERT INTO `score_record` VALUES ('140', '6', '21', '53', '0', '0', '0', '1717031712', null);
INSERT INTO `score_record` VALUES ('141', '6', '21', '53', '0', '0', '0', '1717031712', null);
INSERT INTO `score_record` VALUES ('142', '6', '21', '53', '0', '0', '0', '1717031712', null);
INSERT INTO `score_record` VALUES ('143', '6', '21', '53', '0', '0', '0', '1717031712', null);
INSERT INTO `score_record` VALUES ('144', '6', '21', '53', '0', '0', '3', '1717032808', null);
INSERT INTO `score_record` VALUES ('145', '6', '21', '53', '0', '0', '3', '1717032808', null);
INSERT INTO `score_record` VALUES ('146', '6', '21', '53', '0', '0', '3', '1717032808', null);
INSERT INTO `score_record` VALUES ('147', '6', '21', '53', '0', '0', '0', '1717032808', null);
INSERT INTO `score_record` VALUES ('148', '6', '21', '53', '0', '0', '0', '1717032808', null);
INSERT INTO `score_record` VALUES ('149', '6', '21', '53', '0', '0', '0', '1717032808', null);
INSERT INTO `score_record` VALUES ('150', '6', '21', '53', '0', '0', '0', '1717032808', null);
INSERT INTO `score_record` VALUES ('151', '6', '21', '53', '12', '18', '3', '1717033441', null);
INSERT INTO `score_record` VALUES ('152', '6', '21', '53', '10', '6', '2', '1717033441', null);
INSERT INTO `score_record` VALUES ('153', '6', '21', '53', '14', '25', '2', '1717033441', null);
INSERT INTO `score_record` VALUES ('154', '6', '21', '53', '16', '137', '0', '1717033441', null);
INSERT INTO `score_record` VALUES ('155', '6', '21', '53', '17', '0', '0', '1717033441', null);
INSERT INTO `score_record` VALUES ('156', '6', '21', '53', '18', '0', '0', '1717033441', null);
INSERT INTO `score_record` VALUES ('157', '6', '21', '53', '20', '0', '0', '1717033441', null);
INSERT INTO `score_record` VALUES ('158', '7', '26', '54', '36', '49', '10', '1717342717', null);
INSERT INTO `score_record` VALUES ('159', '7', '26', '54', '37', '54', '0', '1717342717', null);
INSERT INTO `score_record` VALUES ('160', '7', '26', '54', '42', '67', '5', '1717342717', null);
INSERT INTO `score_record` VALUES ('161', '7', '26', '54', '38', '57', '5', '1717342717', null);
INSERT INTO `score_record` VALUES ('162', '7', '26', '54', '39', '0', '0', '1717342717', null);
INSERT INTO `score_record` VALUES ('163', '7', '26', '54', '40', '64', '5', '1717342717', null);
INSERT INTO `score_record` VALUES ('164', '7', '26', '54', '43', '70', '5', '1717342717', null);
INSERT INTO `score_record` VALUES ('165', '7', '26', '54', '44', '75', '5', '1717342717', null);
INSERT INTO `score_record` VALUES ('166', '7', '26', '54', '45', '79', '5', '1717342717', null);
INSERT INTO `score_record` VALUES ('167', '7', '26', '54', '46', '82', '5', '1717342717', null);
