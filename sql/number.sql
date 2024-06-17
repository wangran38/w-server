/*
Navicat MySQL Data Transfer

Source Server         : 3M-老年人
Source Server Version : 50740
Source Host           : 8.134.160.223:3306
Source Database       : oldpp123

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2024-06-15 18:51:40
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for number
-- ----------------------------
DROP TABLE IF EXISTS `number`;
CREATE TABLE `number` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) DEFAULT NULL COMMENT '编号',
  `codetime` varchar(255) DEFAULT NULL COMMENT '编号拼音',
  `reasons` varchar(255) DEFAULT NULL,
  `assessors_id` bigint(20) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=56 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of number
-- ----------------------------
INSERT INTO `number` VALUES ('45', '1596', '2024-04-11', '', '21', '2024-04-11 16:15:33', '2024-04-11 16:15:33');
INSERT INTO `number` VALUES ('44', '156', '2024-04-11', '', '21', '2024-04-11 16:11:45', '2024-04-11 16:11:45');
INSERT INTO `number` VALUES ('43', '1525', '2024-04-11', '1', '21', '2024-04-11 16:02:16', '2024-04-11 16:02:16');
INSERT INTO `number` VALUES ('42', '123', '2024-04-10', '首次评估', '12', '2024-04-11 10:03:11', '2024-04-11 10:03:11');
INSERT INTO `number` VALUES ('41', '2024-04-10-6586', '2024-04-10', '', '21', '2024-04-10 10:56:46', '2024-04-10 10:56:46');
INSERT INTO `number` VALUES ('40', '2024-04-100', '2024-04-10', '', '21', '2024-04-10 10:56:14', '2024-04-10 10:56:14');
INSERT INTO `number` VALUES ('39', '2024-04-10', '2024-04-10', '', '21', '2024-04-10 10:55:53', '2024-04-10 10:55:53');
INSERT INTO `number` VALUES ('38', 'undefined', '2024-04-10', '', '21', '2024-04-10 10:54:38', '2024-04-10 10:54:38');
INSERT INTO `number` VALUES ('37', '222', '2024-04-09', '1', '21', '2024-04-10 10:52:06', '2024-04-10 10:52:06');
INSERT INTO `number` VALUES ('36', '1222', '', '无', '21', '2024-04-10 10:48:44', '2024-04-10 10:48:44');
INSERT INTO `number` VALUES ('35', '23', '', '1', '21', '2024-04-10 10:44:06', '2024-04-10 10:44:06');
INSERT INTO `number` VALUES ('34', '256', '2024-04-10', '0', '21', '2024-04-10 10:30:17', '2024-04-10 10:30:17');
INSERT INTO `number` VALUES ('33', '', '2024-04-11', '2', '21', '2024-04-10 10:29:05', '2024-04-10 10:29:05');
INSERT INTO `number` VALUES ('32', '120', '2024-04-10', '首次评估', '12', '2024-04-10 10:29:02', '2024-04-10 10:29:02');
INSERT INTO `number` VALUES ('31', '136', '', '无', '18', '2024-04-09 10:29:43', '2024-04-09 10:29:43');
INSERT INTO `number` VALUES ('46', '202465484946', '2024-05-08', '首次评估', '2', '2024-05-08 16:37:47', '2024-05-08 16:39:08');
INSERT INTO `number` VALUES ('47', '166666', '2024-05-17', '首次评估', '21', '2024-05-16 15:08:00', '2024-05-17 14:25:46');
INSERT INTO `number` VALUES ('48', '1999', '2024-05-17', '首次评估', '21', '2024-05-17 14:45:36', '2024-05-17 14:45:36');
INSERT INTO `number` VALUES ('49', '1777', '2024-05-17', '首次评估', '21', '2024-05-17 14:46:54', '2024-05-17 14:46:54');
INSERT INTO `number` VALUES ('50', 'a2024051822', '2024-05-17', '首次评估', '22', '2024-05-17 17:08:19', '2024-05-17 17:41:29');
INSERT INTO `number` VALUES ('51', '646464646', '2024-05-27', '首次评估', '21', '2024-05-27 20:19:20', '2024-05-27 20:19:20');
INSERT INTO `number` VALUES ('52', '15665626', '2024-05-29', '首次评估', '21', '2024-05-29 11:00:03', '2024-05-29 11:00:03');
INSERT INTO `number` VALUES ('53', '1112', '2024-05-29', '首次评估', '21', '2024-05-29 11:09:43', '2024-05-29 11:09:43');
INSERT INTO `number` VALUES ('54', 'djsjsjj', '2024-06-02', '常规评估', '26', '2024-06-02 23:36:25', '2024-06-02 23:36:25');
INSERT INTO `number` VALUES ('55', '1778', '2024-06-05', '首次评估', '21', '2024-06-05 09:34:42', '2024-06-05 09:34:42');
