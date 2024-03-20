/*
Navicat MySQL Data Transfer

Source Server         : 本地测试
Source Server Version : 80012
Source Host           : 127.0.0.1:3306
Source Database       : linfeng

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2023-04-21 15:31:37
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for country
-- ----------------------------
DROP TABLE IF EXISTS `country`;
CREATE TABLE `country` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` int(11) DEFAULT NULL,
  `shortname` varchar(200) DEFAULT NULL,
  `name` varchar(200) DEFAULT NULL,
  `mergename` varchar(200) DEFAULT NULL,
  `level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '层级 0 1 2 省市区县',
  `pinyin` varchar(200) DEFAULT NULL,
  `code` varchar(200) DEFAULT NULL,
  `zip` varchar(200) DEFAULT NULL,
  `first` varchar(200) DEFAULT NULL,
  `lng` varchar(200) DEFAULT NULL,
  `lat` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=196 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of country
-- ----------------------------
INSERT INTO `country` VALUES ('1', '0', '亚洲', '亚洲', '', '1', 'yazhou', '3276', '', '', '', '');
INSERT INTO `country` VALUES ('2', '0', '欧洲', '欧洲', '', '1', 'ouzhou', '3277', '', '', '', '');
INSERT INTO `country` VALUES ('3', '0', '北美洲', '北美洲', '', '1', 'beimeizhou', '3278', '', '', '', '');
INSERT INTO `country` VALUES ('4', '0', '非洲', '非洲', '', '1', 'feizhou', '3279', '', '', '', '');
INSERT INTO `country` VALUES ('5', '0', '南美洲', '南美洲', '', '1', 'nanmeizhou', '3280', '', '', '', '');
INSERT INTO `country` VALUES ('6', '0', '大洋洲', '大洋洲', '', '1', 'dayangzhou', '3281', '', '', '', '');
INSERT INTO `country` VALUES ('7', '1', ' 阿富汗', ' 阿富汗', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('8', '1', '叙利亚', '叙利亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('9', '1', '亚美尼亚', '亚美尼亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('10', '1', '朝鲜', '朝鲜', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('11', '1', '以色列', '以色列', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('12', '1', '阿塞拜疆', '阿塞拜疆', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('13', '1', '黎巴嫩', '黎巴嫩', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('14', '1', '巴勒斯坦', '巴勒斯坦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('15', '1', '吉尔吉斯斯坦', '吉尔吉斯斯坦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('16', '1', '科威特', '科威特', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('17', '1', '文莱', '文莱', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('18', '1', '老挝', '老挝', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('19', '1', '土库曼斯坦', '土库曼斯坦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('20', '1', '不丹', '不丹', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('21', '1', '塞浦路斯', '塞浦路斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('22', '1', '斯里兰卡', '斯里兰卡', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('23', '1', '马尔代夫', '马尔代夫', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('24', '1', '也门', '也门', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('25', '1', '东帝汶', '东帝汶', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('26', '2', '梵蒂冈', '梵蒂冈', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('27', '2', '白俄罗斯', '白俄罗斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('28', '2', '乌克兰', '乌克兰', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('29', '2', '丹麦', '丹麦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('30', '2', '斯洛伐克', '斯洛伐克', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('31', '2', '爱尔兰', '爱尔兰', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('32', '2', '圣马力诺', '圣马力诺', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('33', '2', '匈牙利', '匈牙利', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('34', '2', '克罗地亚', '克罗地亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('35', '2', '希腊', '希腊', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('36', '2', '立陶宛', '立陶宛', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('37', '2', '拉脱维亚', '拉脱维亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('38', '2', '瑞士', '瑞士', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('39', '2', '罗马尼亚', '罗马尼亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('40', '2', '卢森堡', '卢森堡', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('41', '2', '摩尔多瓦', '摩尔多瓦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('42', '2', '瑞典', '瑞典', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('43', '2', ' 摩纳哥', ' 摩纳哥', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('44', '2', '马耳他', '马耳他', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('45', '2', '斯洛文尼亚', '斯洛文尼亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('46', '2', '捷克', '捷克', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('47', '2', '马其顿', '马其顿', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('48', '2', '阿尔巴尼亚', '阿尔巴尼亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('49', '2', '安道尔', '安道尔', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('50', '2', '列支敦士登', '列支敦士登', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('51', '2', '葡萄牙', '葡萄牙', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('52', '2', '波斯尼亚和黑塞哥维那', '波斯尼亚和黑塞哥维那', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('53', '2', '比利时', '比利时', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('54', '2', '奥地利', '奥地利', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('55', '2', '爱沙尼亚', '爱沙尼亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('56', '2', '冰岛', '冰岛', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('57', '3', '巴哈马', '巴哈马', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('58', '3', '巴拿马', '巴拿马', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('59', '3', '尼加拉瓜', '尼加拉瓜', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('60', '3', '巴巴多斯', '巴巴多斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('61', '3', '牙买加', '牙买加', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('62', '3', '海地', '海地', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('63', '3', '危地马拉', '危地马拉', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('64', '3', '古巴', '古巴', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('65', '3', '洪都拉斯', '洪都拉斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('66', '3', '格林纳达', '格林纳达', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('67', '3', '哥斯达黎加', '哥斯达黎加', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('68', '3', '多米尼加', '多米尼加', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('69', '3', '伯利兹', '伯利兹', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('70', '3', '萨尔瓦多', '萨尔瓦多', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('71', '3', '圣卢西亚', '圣卢西亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('72', '3', '圣基茨和尼维斯', '圣基茨和尼维斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('73', '3', '圣文森特和格林纳丁斯', '圣文森特和格林纳丁斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('74', '3', '特立尼达和多巴哥', '特立尼达和多巴哥', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('75', '3', '安提瓜和巴布达', '安提瓜和巴布达', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('76', '3', '多米尼克国', '多米尼克国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('77', '4', '贝　宁', '贝　宁', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('78', '4', '吉布提', '吉布提', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('79', '4', '赞比亚', '赞比亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('80', '4', '毛里塔尼亚', '毛里塔尼亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('81', '4', '加　蓬', '加　蓬', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('82', '4', '几内亚', '几内亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('83', '4', '冈比亚', '冈比亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('84', '4', '博茨瓦纳', '博茨瓦纳', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('85', '4', '塞内加尔', '塞内加尔', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('86', '4', '中非共和国', '中非共和国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('87', '4', '安哥拉', '安哥拉', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('88', '4', '科特迪瓦', '科特迪瓦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('89', '4', '塞拉利昂', '塞拉利昂', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('90', '4', '厄立特里亚', '厄立特里亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('91', '4', '马达加斯加', '马达加斯加', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('92', '4', '马　里', '马　里', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('93', '4', '布隆迪', '布隆迪', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('94', '4', '利比亚', '利比亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('95', '4', '索马里', '索马里', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('96', '4', '津巴布韦', '津巴布韦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('97', '4', '斯威士兰', '斯威士兰', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('98', '4', '几内亚比绍', '几内亚比绍', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('99', '4', '刚果共和国', '刚果共和国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('100', '4', '苏　丹', '苏　丹', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('101', '4', '喀麦隆', '喀麦隆', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('102', '4', '莱索托', '莱索托', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('103', '4', '塞舌尔', '塞舌尔', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('104', '4', '刚果民主共和国', '刚果民主共和国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('105', '4', '乍　得', '乍　得', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('106', '4', '科摩罗', '科摩罗', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('107', '4', '突尼斯', '突尼斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('108', '4', '毛里求斯', '毛里求斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('109', '4', '利比里亚', '利比里亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('110', '4', '布基纳法索', '布基纳法索', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('111', '4', '圣多美和普林西比', '圣多美和普林西比', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('112', '4', '多　哥', '多　哥', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('113', '4', '佛得角', '佛得角', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('114', '4', '马拉维', '马拉维', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('115', '4', '乌干达', '乌干达', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('116', '4', '纳米比亚', '纳米比亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('117', '5', '委内瑞拉', '委内瑞拉', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('118', '5', '圭亚那', '圭亚那', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('119', '5', '苏里南', '苏里南', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('120', '5', '厄瓜多尔', '厄瓜多尔', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('121', '5', '玻利维亚', '玻利维亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('122', '5', '巴拉圭', '巴拉圭', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('123', '5', '乌拉圭', '乌拉圭', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('124', '6', '新西兰 ', '新西兰 ', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('125', '6', '巴布亚新几内亚', '巴布亚新几内亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('126', '6', '所罗门群岛', '所罗门群岛', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('127', '6', '瓦努阿图', '瓦努阿图', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('128', '6', '密克罗尼西亚', '密克罗尼西亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('129', '6', '马绍尔群岛', '马绍尔群岛', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('130', '6', '帕劳', '帕劳', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('131', '6', '瑙鲁', '瑙鲁', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('132', '6', '基里巴斯', '基里巴斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('133', '6', '图瓦卢', '图瓦卢', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('134', '6', '萨摩亚', '萨摩亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('135', '6', '斐济群岛', '斐济群岛', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('136', '6', '汤加', '汤加', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('137', '1', '印度', '印度', '', '2', 'yindu', '', '', '', '', '');
INSERT INTO `country` VALUES ('138', '1', '蒙古', '蒙古', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('139', '1', '日本', '日本', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('140', '1', '沙特', '沙特', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('141', '1', '巴林', '巴林', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('142', '1', '巴基斯坦', '巴基斯坦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('143', '1', '新加坡', '新加坡', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('144', '1', '伊拉克', '伊拉克', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('145', '1', '缅甸', '缅甸', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('146', '1', '孟加拉国', '孟加拉国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('147', '1', '韩国', '韩国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('148', '1', '约旦', '约旦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('149', '1', '尼泊尔', '尼泊尔', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('150', '1', '越南', '越南', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('151', '1', '印尼', '印尼', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('152', '1', '伊朗', '伊朗', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('153', '1', '乌兹别克斯坦', '乌兹别克斯坦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('154', '1', '菲律宾', '菲律宾', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('155', '1', '阿曼', '阿曼', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('156', '1', '土耳其', '土耳其', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('157', '1', '哈萨克斯坦', '哈萨克斯坦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('158', '1', '柬埔寨', '柬埔寨', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('159', '1', '马来西亚', '马来西亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('160', '1', '卡塔尔', '卡塔尔', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('161', '1', '阿联酋', '阿联酋', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('162', '1', '塔吉克斯坦', '塔吉克斯坦', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('163', '1', '泰国', '泰国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('164', '2', '波兰', '波兰', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('165', '2', '保加利亚', '保加利亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('166', '2', '西班牙', '西班牙', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('167', '2', '德国', '德国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('168', '2', '法国', '法国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('169', '2', '意大利', '意大利', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('170', '2', '英国', '英国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('171', '2', '挪威', '挪威', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('172', '2', '芬兰', '芬兰', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('173', '2', '荷兰', '荷兰', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('174', '2', '俄罗斯', '俄罗斯', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('175', '3', '墨西哥', '墨西哥', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('176', '3', '加拿大', '加拿大', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('177', '3', '美国', '美国', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('178', '4', '南　非', '南　非', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('179', '4', '尼日尔', '尼日尔', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('180', '4', '尼日利亚', '尼日利亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('181', '4', '阿尔及利亚', '阿尔及利亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('182', '4', '埃　及', '埃　及', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('183', '4', '埃塞俄比亚', '埃塞俄比亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('184', '4', '加　纳', '加　纳', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('185', '4', '肯尼亚', '肯尼亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('186', '4', '卢旺达', '卢旺达', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('187', '4', '莫桑比克', '莫桑比克', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('188', '4', '坦桑尼亚', '坦桑尼亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('189', '4', '摩洛哥', '摩洛哥', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('190', '5', '哥伦比亚', '哥伦比亚', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('191', '5', '秘鲁', '秘鲁', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('192', '5', '巴西', '巴西', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('193', '5', '智利', '智利', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('194', '5', '阿根廷', '阿根廷', '', '2', '', '', '', '', '', '');
INSERT INTO `country` VALUES ('195', '6', '澳大利亚 ', '澳大利亚 ', '', '2', '', '', '', '', '', '');
