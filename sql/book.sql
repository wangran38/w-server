/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 80012
Source Host           : localhost:3306
Source Database       : lsk

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2024-04-18 14:31:00
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for book
-- ----------------------------
DROP TABLE IF EXISTS `book`;
CREATE TABLE `book` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) DEFAULT NULL COMMENT '姓名',
  `chaptername` varchar(64) DEFAULT NULL COMMENT '章节名称',
  `level` int(3) DEFAULT NULL COMMENT '级别',
  `name` varchar(64) DEFAULT NULL COMMENT '事件名称',
  `content` longtext COMMENT '哪本书',
  `created` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated` int(11) DEFAULT NULL COMMENT '修改时间',
  `isdel` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否启用 默认1 正常 2 已删除',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=94 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of book
-- ----------------------------
INSERT INTO `book` VALUES ('14', '13', '第一节', '2', '饮食照料', '饮食照料', '1713259654', null, '0');
INSERT INTO `book` VALUES ('13', '12', '第一章', '1', '生活照料', '生活照料', '1713259618', null, '0');
INSERT INTO `book` VALUES ('12', '0', '初级', '0', '初级', '初级', '1713259587', null, '0');
INSERT INTO `book` VALUES ('15', '14', '单元一', '3', '为老年人摆放进食体位', '为老年人摆放进食体位', '1713259696', null, '0');
INSERT INTO `book` VALUES ('16', '14', '单元三', '3', '帮助老年人进食、水种类和量的观察与记录', '帮助老年人进食、水种类和量的观察与记录', '1713260131', null, '0');
INSERT INTO `book` VALUES ('17', '14', '单元四', '3', '为老年人发放治疗饮食', '为老年人发放治疗饮食', '1713260164', null, '0');
INSERT INTO `book` VALUES ('18', '14', '单元二', '3', '帮助老年人进食、进水', '帮助老年人进食、进水', '1713260217', null, '0');
INSERT INTO `book` VALUES ('19', '13', '第二节', '2', '排泄照料', '排泄照料', '1713317036', null, '0');
INSERT INTO `book` VALUES ('20', '19', '单元一', '3', '帮助老年人如厕', '帮助老年人如厕', '1713317281', null, '0');
INSERT INTO `book` VALUES ('21', '19', '单元二', '3', '帮助卧床老年人使用便盆及尿壶排泄', '帮助卧床老年人使用便盆及尿壶排泄', '1713317372', null, '0');
INSERT INTO `book` VALUES ('22', '19', '单元三', '3', '为老年人更换尿垫、纸尿裤', '为老年人更换尿垫、纸尿裤', '1713317434', null, '0');
INSERT INTO `book` VALUES ('23', '19', '单元四', '3', '采集老年人的二便标本', '采集老年人的二便标本', '1713317477', null, '0');
INSERT INTO `book` VALUES ('24', '19', '单元五', '3', '使用开塞露辅助老年人排便', '使用开塞露辅助老年人排便', '1713317511', null, '0');
INSERT INTO `book` VALUES ('25', '19', '单元六', '3', '协助老年人呕吐时变换体位', '协助老年人呕吐时变换体位', '1713317543', null, '0');
INSERT INTO `book` VALUES ('26', '13', '第三节', '2', '睡眠照料', '睡眠照料', '1713317603', null, '0');
INSERT INTO `book` VALUES ('27', '26', '单元一', '3', '为老年人布置睡眠环境', '为老年人布置睡眠环境', '1713317655', null, '0');
INSERT INTO `book` VALUES ('28', '26', '单元二', '3', '老年人睡眠状况观察记录', '老年人睡眠状况观察记录', '1713317686', null, '0');
INSERT INTO `book` VALUES ('29', '13', '第四节', '2', '清洁照料', '清洁照料', '1713317846', null, '0');
INSERT INTO `book` VALUES ('30', '29', '单元一', '3', '为老年人整理更换床单位', '为老年人整理更换床单位', '1713317880', null, '0');
INSERT INTO `book` VALUES ('31', '29', '单元二', '3', '为老年人清洁口腔', '为老年人清洁口腔', '1713317912', null, '0');
INSERT INTO `book` VALUES ('32', '29', '单元三', '3', '为老年人摘戴、清洗义齿', '为老年人摘戴、清洗义齿', '1713317947', null, '0');
INSERT INTO `book` VALUES ('33', '29', '单元四', '3', '头发清洁与梳理', '头发清洁与梳理', '1713318091', null, '0');
INSERT INTO `book` VALUES ('34', '29', '单元五', '3', '身体清洁', '身体清洁', '1713318117', null, '0');
INSERT INTO `book` VALUES ('35', '29', '单元六', '3', '为老年人修饰仪容仪表', '为老年人修饰仪容仪表', '1713318157', null, '0');
INSERT INTO `book` VALUES ('36', '29', '单元七', '3', '为老年人更衣', '为老年人更衣', '1713318189', null, '0');
INSERT INTO `book` VALUES ('37', '29', '单元八', '3', '卧床老年人预防压疮', '卧床老年人预防压疮', '1713318247', null, '0');
INSERT INTO `book` VALUES ('38', '12', '第二章', '1', '基础护理', '基础护理', '1713318586', null, '0');
INSERT INTO `book` VALUES ('39', '38', '第一节', '2', '用药照料', '用药照料', '1713318619', null, '0');
INSERT INTO `book` VALUES ('40', '39', '单元一', '3', '协助老年人服药', '协助老年人服药', '1713318687', null, '0');
INSERT INTO `book` VALUES ('41', '39', '单元二', '3', '常用药的管理及注意事项', '常用药的管理及注意事项', '1713318757', null, '0');
INSERT INTO `book` VALUES ('42', '38', '第二节', '2', '冷热应用', '冷热应用', '1713318813', null, '0');
INSERT INTO `book` VALUES ('43', '42', '单元一', '3', '老年人热水袋的应用', '老年人热水袋的应用', '1713318848', null, '0');
INSERT INTO `book` VALUES ('44', '42', '单元二', '3', '老年人湿热敷的应用', '老年人湿热敷的应用', '1713319049', null, '0');
INSERT INTO `book` VALUES ('45', '42', '单元三', '3', '老年人皮肤观察', '老年人皮肤观察', '1713319623', null, '0');
INSERT INTO `book` VALUES ('46', '38', '第三节', '2', '遗体照料', '遗体照料', '1713319655', null, '0');
INSERT INTO `book` VALUES ('47', '46', '单元一', '3', '清洁遗体', '清洁遗体', '1713319699', null, '0');
INSERT INTO `book` VALUES ('48', '46', '单元二', '3', '整理遗物', '整理遗物', '1713319731', null, '0');
INSERT INTO `book` VALUES ('49', '12', '第三章', '1', '康复护理', '康复护理', '1713319834', null, '0');
INSERT INTO `book` VALUES ('50', '49', '第一节', '2', '康乐活动', '康乐活动', '1713319873', null, '0');
INSERT INTO `book` VALUES ('51', '50', '单元一', '3', '教老年人手工活动', '教老年人手工活动', '1713319902', null, '0');
INSERT INTO `book` VALUES ('52', '50', '单元二', '3', '为老年人示范娱乐游戏活动', '为老年人示范娱乐游戏活动', '1713320002', null, '0');
INSERT INTO `book` VALUES ('53', '49', '第二节', '2', '活动保护', '活动保护', '1713320162', null, '0');
INSERT INTO `book` VALUES ('54', '53', '单元一', '3', '教老年人使用拐杖进行活动', '教老年人使用拐杖进行活动', '1713320193', null, '0');
INSERT INTO `book` VALUES ('55', '53', '单元二', '3', '救老年人使用轮椅进行活动', '救老年人使用轮椅进行活动', '1713320223', null, '0');
INSERT INTO `book` VALUES ('56', '53', '单元三', '3', '使用平车转运老年人', '使用平车转运老年人', '1713320249', null, '0');
INSERT INTO `book` VALUES ('57', '0', '中级', '0', '中级', '中级', '1713322520', null, '0');
INSERT INTO `book` VALUES ('58', '57', '第一章', '1', '生活照料', '生活照料', '1713322614', null, '0');
INSERT INTO `book` VALUES ('59', '58', '第一节', '2', '饮食照料', '饮食照料', '1713322713', null, '0');
INSERT INTO `book` VALUES ('60', '59', '单元一', '3', '带鼻饲管老年人的进食照料', '带鼻饲管老年人的进食照料', '1713322757', null, '0');
INSERT INTO `book` VALUES ('61', '59', '单元二', '3', '救助噎食、误吸的老年人', '救助噎食、误吸的老年人', '1713322792', null, '0');
INSERT INTO `book` VALUES ('62', '58', '第二节', '2', '排泄照料', '排泄照料', '1713322885', null, '0');
INSERT INTO `book` VALUES ('63', '62', '单元一', '3', '使用人工取便的方法辅助老年人排便', '使用人工取便的方法辅助老年人排便\n', '1713322916', null, '0');
INSERT INTO `book` VALUES ('64', '62', '单元二', '3', '为留置导尿的老年人更换一次性引流袋(尿袋)', '为留置导尿的老年人更换一次性引流袋(尿袋)', '1713322987', null, '0');
INSERT INTO `book` VALUES ('65', '62', '单元三', '3', '观察留置导尿的老年人尿量及颜色并记录异常', '观察留置导尿的老年人尿量及颜色并记录异常', '1713323346', null, '0');
INSERT INTO `book` VALUES ('66', '62', '单元四', '3', '为有肠造的老年人更换粪袋', '为有肠造的老年人更换粪袋', '1713323388', null, '0');
INSERT INTO `book` VALUES ('67', '58', '第三节', '2', '睡眠照料', '睡眠照料', '1713323413', null, '0');
INSERT INTO `book` VALUES ('68', '67', '单元一', '3', '识别并改善影响老年人睡眠的环境因素', '识别并改善影响老年人睡眠的环境因素', '1713323442', null, '0');
INSERT INTO `book` VALUES ('69', '67', '单元二', '3', '照料有睡眠障碍的老年人入睡', '照料有睡眠障碍的老年人入睡', '1713323480', null, '0');
INSERT INTO `book` VALUES ('70', '67', '单元三', '3', '指导老年人改变不良睡眠习惯', '指导老年人改变不良睡眠习惯', '1713323505', null, '0');
INSERT INTO `book` VALUES ('71', '58', '第四节', '2', '清洁照料', '清洁照料', '1713323547', null, '0');
INSERT INTO `book` VALUES ('72', '71', '单元一', '3', '为老年人进行口腔护理', '为老年人进行口腔护理', '1713323619', null, '0');
INSERT INTO `book` VALUES ('73', '71', '单元二', '3', '对老年人进行床旁隔离', '对老年人进行床旁隔离', '1713323664', null, '0');
INSERT INTO `book` VALUES ('74', '71', '单元三', '3', '老年人房间终末消毒', '老年人房间终末消毒', '1713323707', null, '0');
INSERT INTO `book` VALUES ('75', '57', '第二章', '1', '基础护理', '基础护理', '1713323834', null, '0');
INSERT INTO `book` VALUES ('76', '75', '第一节', '2', '用药照料', '用药照料', '1713323897', null, '0');
INSERT INTO `book` VALUES ('77', '76', '单元一', '3', '为老年人进行雾化吸入', '为老年人进行雾化吸入', '1713323925', null, '0');
INSERT INTO `book` VALUES ('78', '76', '单元二', '3', '为老年人应用眼、耳、鼻等外用药', '为老年人应用眼、耳、鼻等外用药', '1713324001', null, '0');
INSERT INTO `book` VALUES ('79', '76', '单元三', '3', '为I度压疮老年人提供护理', '为I度压疮老年人提供护理', '1713324076', null, '0');
INSERT INTO `book` VALUES ('80', '75', '第二节', '2', '冷热应用', '冷热应用', '1713324101', null, '0');
INSERT INTO `book` VALUES ('81', '80', '单元一', '3', '为老年人测量体温', '为老年人测量体温', '1713324143', null, '0');
INSERT INTO `book` VALUES ('82', '80', '单元二', '3', '使用冰袋为高热老年人进行物理降温', '使用冰袋为高热老年人进行物理降温', '1713324168', null, '0');
INSERT INTO `book` VALUES ('83', '80', '单元三', '3', '使用温水擦浴为高热老年人进行物理降温', '使用温水擦浴为高热老年人进行物理降温', '1713324206', null, '0');
INSERT INTO `book` VALUES ('84', '75', '第三节', '2', '临终关怀', '临终关怀', '1713324273', null, '0');
INSERT INTO `book` VALUES ('85', '84', '单元一', '3', '运用肢体语言为临终老年人提供慰藉支持', '运用肢体语言为临终老年人提供慰藉支持', '1713324293', null, '0');
INSERT INTO `book` VALUES ('86', '84', '单元二', '3', '为临终老年人及家属提供精神安慰支持', '为临终老年人及家属提供精神安慰支持', '1713324325', null, '0');
INSERT INTO `book` VALUES ('87', '57', '第三章', '0', '康复护理', '康复护理', '1713324490', null, '0');
INSERT INTO `book` VALUES ('88', '87', '第一节', '2', '康乐活动照护', '康乐活动照护', '1713336736', null, '0');
INSERT INTO `book` VALUES ('89', '88', '单元一', '3', '教老年人使用健身器材', '教老年人使用健身器材', '1713336773', null, '0');
INSERT INTO `book` VALUES ('90', '88', '单元二', '3', '帮助老年人进行床上体位转换', '帮助老年人进行床上体位转换', '1713336803', null, '0');
INSERT INTO `book` VALUES ('91', '87', '第二节', '2', '功能锻炼', '功能锻炼', '1713336829', null, '0');
INSERT INTO `book` VALUES ('92', '91', '单元一', '3', '帮助老年人进行穿脱衣服训练', '帮助老年人进行穿脱衣服训练', '1713336850', null, '0');
INSERT INTO `book` VALUES ('93', '91', '单元二', '3', '帮助老年人进行站立、行走等活动', '帮助老年人进行站立、行走等活动', '1713336877', null, '0');
