/*
Navicat MySQL Data Transfer

Source Server         : 林峰会展
Source Server Version : 50740
Source Host           : localhost:3306
Source Database       : linfeng

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2023-12-11 14:40:35
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `auth_rule`;
CREATE TABLE `auth_rule` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT '0',
  `type` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `pathname` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `ismenu` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否启用 默认1 菜单 0 文件',
  `created` int(11) DEFAULT NULL,
  `updated` int(11) DEFAULT NULL,
  `deletetime` int(11) DEFAULT NULL,
  `weigh` int(11) DEFAULT NULL,
  `status` varchar(40) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=368 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of auth_rule
-- ----------------------------
INSERT INTO `auth_rule` VALUES ('2', '1', 'file', 'el-icon-notebook-2', '/system/admin', '管理员列表', '用户列表', '0', '1614759419', '1673339018', null, '1', 'normal', 'system/admin/index');
INSERT INTO `auth_rule` VALUES ('3', '1', 'file', 'el-icon-notebook-2', '/system/rules', '菜单管理', '前端菜单', '0', '1614759461', '1662431244', null, '2', 'normal', 'system/rules/index');
INSERT INTO `auth_rule` VALUES ('1', '0', 'file', 'el-icon-menu', '/system', '系统管理', '用户管理', '1', '1621413412', null, null, '0', 'normal', null);
INSERT INTO `auth_rule` VALUES ('5', '0', 'file', 'el-icon-setting', '/share', '公共配置', '公共配置', '1', '1619409666', '1662445753', null, '0', 'normal', null);
INSERT INTO `auth_rule` VALUES ('358', '5', '', 'el-icon-s-flag', '/share/city', '城市管理', '', '0', '1671527625', '1671527625', '0', '0', '', 'share/city');
INSERT INTO `auth_rule` VALUES ('357', '0', 'file', 'el-icon-coin', '/all', '展会管理', '', '1', '1671525883', '1679374904', '0', '0', '', '');
INSERT INTO `auth_rule` VALUES ('359', '357', '', 'el-icon-s-fold', '/law/category', '分类管理', '', '0', '1671528375', '1671528375', '0', '0', '', 'law/category');
INSERT INTO `auth_rule` VALUES ('361', '1', '', 'el-icon-user', '/system/group', '组别管理', '', '0', '1673574878', '1673574936', '0', '0', '', 'system/group/index');
INSERT INTO `auth_rule` VALUES ('363', '357', '', 'el-icon-s-platform', '/exhibition/exhibition', '国内展会', '', '0', '1680521306', '1680521412', '0', '0', '', 'exhibition/index');
INSERT INTO `auth_rule` VALUES ('364', '357', '', 'el-icon-menu', '/law/news', '展会资讯', '', '0', '1681192232', '1681192279', '0', '0', '', 'law/news/index');
INSERT INTO `auth_rule` VALUES ('365', '5', '', 'el-icon-bangzhu', '/share/country', '海外地区', '', '0', '1681957457', '1681957539', '0', '0', '', 'share/country/index');
INSERT INTO `auth_rule` VALUES ('366', '0', '', 'el-icon-money', '/dingdan', '订单管理', '', '1', '1693446059', '1693446195', '0', '0', '', '');
INSERT INTO `auth_rule` VALUES ('367', '366', '', 'el-icon-s-unfold', '/userpay/userpay', '门票订单', '', '0', '1693446135', '1693446169', '0', '0', '', 'userpay/index');
