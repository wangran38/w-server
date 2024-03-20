/*
Navicat MySQL Data Transfer

Source Server         : 林峰会展
Source Server Version : 50740
Source Host           : localhost:3306
Source Database       : linfeng

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2023-10-25 14:48:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for usercontact
-- ----------------------------
DROP TABLE IF EXISTS `usercontact`;
CREATE TABLE `usercontact` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `openid` varchar(200) CHARACTER SET utf8 DEFAULT NULL COMMENT '用户openid',
  `companyname` varchar(200) CHARACTER SET utf8 DEFAULT NULL COMMENT '公司名称/个人',
  `out_trade_no` varchar(32) CHARACTER SET utf8 DEFAULT NULL COMMENT '订单号',
  `name` varchar(200) CHARACTER SET utf8 DEFAULT NULL COMMENT '联系人姓名',
  `sex` varchar(200) DEFAULT NULL,
  `job` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `phone` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `email` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `idcard` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `cardnum` varchar(200) CHARACTER SET utf8 DEFAULT NULL,
  `createtime` bigint(20) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of usercontact
-- ----------------------------
INSERT INTO `usercontact` VALUES ('1', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '测试公式', 'ZH202308301139105171', '李清波', '男', null, '18889363060', '252588119@qq.com', '4600000000000000', '', '1693366750', null, null);
INSERT INTO `usercontact` VALUES ('2', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '测试公式', 'ZH202308301139105171', '王冉', '男', null, '18888888888', 'wangran@qq.com', '46111111111111111', '', '1693366750', null, null);
INSERT INTO `usercontact` VALUES ('3', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '', 'ZH202308301553121959', '', '', null, '', '', '', '', '1693381992', null, null);
INSERT INTO `usercontact` VALUES ('4', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '', 'ZH202308301555331226', '', '', null, '', '', '', '', '1693382133', null, null);
INSERT INTO `usercontact` VALUES ('5', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '', 'ZH202308301558041912', 'asdf', '', null, '123123123', '', '', '', '1693382284', null, null);
INSERT INTO `usercontact` VALUES ('6', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '测试公司', 'ZH202308311734492742', '王冉测试', '男', null, '18907693872', '23323232@qq.com', '321523456223512315', '', '1693474489', null, null);
INSERT INTO `usercontact` VALUES ('7', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', 'test', 'ZH202308311747258124', 'zhagnsan', '男', null, '18889363060', '', '4600000000000000', '', '1693475245', null, null);
INSERT INTO `usercontact` VALUES ('8', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', 'test', 'ZH202308311747258124', 'lisi', '女', null, '18888888888', '', '46111111111111111', '', '1693475245', null, null);
INSERT INTO `usercontact` VALUES ('9', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '988', 'ZH202309141406124552', '王冉', '男', null, '18907693872', '23323232@qq.com', '321523456223512315', '', '1694671572', null, null);
INSERT INTO `usercontact` VALUES ('10', 'oNbgG52oRb1fj7deRDRjY63usRWQ', '浙江芯源新材料有限公司', 'ZH202309192252052515', '金星华', '男', null, '18221135601', 'xinghua-jin@163.com', '222401197106061817', '', '1695135125', null, null);
INSERT INTO `usercontact` VALUES ('11', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', 'yyu', 'ZH202309192340262367', '王冉', '男', null, '18907693872', '23323232@qq.com', '321523456223512315', '', '1695138026', null, null);
INSERT INTO `usercontact` VALUES ('12', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', 'yyu', 'ZH202309192340262367', '', '', null, '', '', '', '', '1695138026', null, null);
INSERT INTO `usercontact` VALUES ('13', 'oNbgG58fmji_SsLgyriuzpqK9W54', '中钢集团邢台机械轧辊有限公司', 'ZH202309302125358315', '耿浩东', '男', null, '13211078303', '2536951724@qq.com', '130582199807040014', '', '1696080335', null, null);
INSERT INTO `usercontact` VALUES ('14', 'oNbgG55Qa8Del8sWkLPd7hMw45Zg', '佛山市华磊华扬有限公司', 'ZH202310200942514802', '林峰', '男', null, '18520220700', '', '460033198208160011', '', '1697766171', null, null);
