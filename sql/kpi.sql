/*
 Navicat MySQL Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : oldpp

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 07/04/2024 09:22:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for kpi
-- ----------------------------
DROP TABLE IF EXISTS `kpi`;
CREATE TABLE `kpi`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pid` bigint(20) NULL DEFAULT NULL COMMENT '姓名',
  `kname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '指标名称',
  `level` int(3) NULL DEFAULT NULL COMMENT '级别',
  `kcontent` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '级别',
  `created` int(11) NULL DEFAULT NULL COMMENT '创建时间',
  `updated` int(11) NULL DEFAULT NULL COMMENT '修改时间',
  `weigh` int(11) NULL DEFAULT NULL COMMENT '排序',
  `status` int(11) NULL DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 58 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of kpi
-- ----------------------------
INSERT INTO `kpi` VALUES (1, 0, 'gb2022标准', 0, 'gb2022标准', 1711013411, NULL, 1000, 0);
INSERT INTO `kpi` VALUES (4, 1, '自理能力评估', 1, '进食、修饰、洗澡、穿/脱上衣、穿/脱裤子和鞋袜、小便控制、大便控制、如厕', 1711014750, NULL, 901, 0);
INSERT INTO `kpi` VALUES (5, 1, '基础运动能力', 1, '床上体位转移、座椅转移、平地行走、上下楼梯', 1711014806, NULL, 801, 0);
INSERT INTO `kpi` VALUES (6, 1, '精神状态评估', 1, '时间定向、空间定向', 1711014849, NULL, 701, 0);
INSERT INTO `kpi` VALUES (7, 1, '感知觉与社会参与', 1, '视力、听力、执行日常事务', 1711014890, NULL, 601, 0);
INSERT INTO `kpi` VALUES (10, 4, '进食：使用适当的器具将食物送入口中咽下', 2, '进食：使用适当的器具将食物送入口中咽下', 1711071972, NULL, 898, 0);
INSERT INTO `kpi` VALUES (11, 0, '企业标准', 1, '企业标准', 1711073674, NULL, 500, 0);
INSERT INTO `kpi` VALUES (12, 4, '修饰：洗脸、刷牙、梳头、刮脸、剪指（趾）甲', 2, '修饰：洗脸、刷牙、梳头、刮脸、剪指（趾）甲', 1711243745, NULL, 899, 0);
INSERT INTO `kpi` VALUES (13, 5, '床上体位转移：卧床翻身及坐下躺下', 2, '床上体位转移：卧床翻身及坐下躺下', 1711244208, NULL, 891, 0);
INSERT INTO `kpi` VALUES (14, 4, '洗澡：清洗和擦干身体', 2, '洗澡：清洗和擦干身体', 1711245596, NULL, 897, 0);
INSERT INTO `kpi` VALUES (16, 4, '穿/脱上衣：穿/脱上衣服、系扣、拉拉链等', 2, '穿/脱上衣：穿/脱上衣服、系扣、拉拉链等', 1711245915, NULL, 896, 0);
INSERT INTO `kpi` VALUES (17, 4, '穿/脱裤子和鞋袜：穿/脱裤子和鞋袜等', 2, '穿/脱裤子和鞋袜：穿/脱裤子和鞋袜等', 1711246009, NULL, 895, 0);
INSERT INTO `kpi` VALUES (18, 4, '小便控制：控制和排除尿液的能力', 2, '小便控制：控制和排除尿液的能力', 1711246141, NULL, 894, 0);
INSERT INTO `kpi` VALUES (19, 1, '大便控制：控制和排出粪便的能力', 2, '大便控制：控制和排出粪便的能力', 1711246204, NULL, 893, 0);
INSERT INTO `kpi` VALUES (20, 4, '如厕：上厕所排泄大小便，并清洁身体', 2, '如厕：上厕所排泄大小便，并清洁身体', 1711246293, NULL, 892, 0);
INSERT INTO `kpi` VALUES (21, 5, '床椅转移：从座位到站位，再从站位到坐位的转换过程', 2, '床椅转移：从座位到站位，再从站位到坐位的转换过程', 1711246458, NULL, 890, 0);
INSERT INTO `kpi` VALUES (22, 5, '平地行走：双脚交互的方式在地面行动，总是一只脚在前', 2, '平地行走：双脚交互的方式在地面行动，总是一只脚在前', 1711246552, NULL, 889, 0);
INSERT INTO `kpi` VALUES (23, 5, '上下楼梯：双腿交替完成楼梯台阶连续的上下移动', 2, '上下楼梯：双腿交替完成楼梯台阶连续的上下移动', 1711246642, NULL, 888, 0);
INSERT INTO `kpi` VALUES (24, 6, '时间定向：知道并确认时间的能力', 2, '时间定向：知道并确认时间的能力', 1711246712, NULL, 887, 0);
INSERT INTO `kpi` VALUES (25, 6, '空间定向：知道并确认空间的能力', 2, '空间定向：知道并确认空间的能力', 1711246767, NULL, 886, 0);
INSERT INTO `kpi` VALUES (26, 6, '人物定向：知道并确认人物的能力', 2, '人物定向：知道并确认人物的能力', 1711246832, NULL, 885, 0);
INSERT INTO `kpi` VALUES (27, 6, '记忆：短期、近期和远期记忆能力', 2, '记忆：短期、近期和远期记忆能力', 1711246950, NULL, 884, 0);
INSERT INTO `kpi` VALUES (28, 6, '理解能力：理解语言信息和非语言信息的能力（可借助平时使用助听设备等）及理解别人的话', 2, '理解能力：理解语言信息和非语言信息的能力（可借助平时使用助听设备等）及理解别人的话', 1711247098, NULL, 883, 0);
INSERT INTO `kpi` VALUES (29, 6, '表达能力：表达信息能力，包括口头的和非口头的，及表达自己的想法', 2, '表达能力：表达信息能力，包括口头的和非口头的，及表达自己的想法', 1711247255, NULL, 882, 0);
INSERT INTO `kpi` VALUES (30, 6, '攻击行为：身体攻击行为（如打/踢/推/咬/抓/摔东西）和语言攻击行为（如骂人、语言威胁、尖叫）', 2, '攻击行为：身体攻击行为（如打/踢/推/咬/抓/摔东西）和语言攻击行为（如骂人、语言威胁、尖叫）', 1711247623, NULL, 881, 0);
INSERT INTO `kpi` VALUES (31, 6, '抑郁症状：存在情绪低落、兴趣减退、活力减退等症状，甚至出现妄想、幻觉、自杀念头或自杀行为', 2, '抑郁症状：存在情绪低落、兴趣减退、活力减退等症状，甚至出现妄想、幻觉、自杀念头或自杀行为', 1711247849, NULL, 880, 0);
INSERT INTO `kpi` VALUES (32, 0, '国际标准', 0, '2332', 1712132284, NULL, 0, 0);
INSERT INTO `kpi` VALUES (33, 11, '日常生活活动能力', 0, '日常生活活动能力', 1712132506, NULL, 0, 0);
INSERT INTO `kpi` VALUES (34, 11, '认知能力', 0, '', 1712132517, NULL, 0, 0);
INSERT INTO `kpi` VALUES (35, 11, '精神状态与社会交流能力', 0, '', 1712132536, NULL, 0, 0);
INSERT INTO `kpi` VALUES (36, 33, 'B1进食', 0, '', 1712132605, NULL, 0, 0);
INSERT INTO `kpi` VALUES (37, 33, 'B2洗澡', 0, '过程中需他人帮助', 1712133609, NULL, 0, 0);
INSERT INTO `kpi` VALUES (42, 33, 'B6小便控制', 0, '', 1712134299, NULL, 0, 0);
INSERT INTO `kpi` VALUES (38, 33, 'B3修饰', 0, '5 分，可自己独立完成\n0 分，需他人帮助', 1712133657, NULL, 0, 0);
INSERT INTO `kpi` VALUES (39, 33, 'B4穿脱衣裤，鞋子', 0, '10分，可独立完成\n5分，需部分帮助(能自己穿脱，但需他人帮助整理衣物、系扣/鞋带拉拉链)', 1712133685, NULL, 0, 0);
INSERT INTO `kpi` VALUES (40, 33, 'B5大便控制', 0, '10分，可控制大便\n5分，偶尔失控(每周<1次)，或需要他人提示\n0分，完全失控', 1712133728, NULL, 0, 0);
INSERT INTO `kpi` VALUES (43, 33, 'B7如厕', 0, '', 1712134396, NULL, 0, 0);
INSERT INTO `kpi` VALUES (44, 33, 'B8 床椅转移', 0, '', 1712134468, NULL, 0, 0);
INSERT INTO `kpi` VALUES (45, 33, 'B9 平地行走', 0, '', 1712134577, NULL, 0, 0);
INSERT INTO `kpi` VALUES (46, 33, 'B10 上下楼梯', 0, '', 1712134716, NULL, 0, 0);
INSERT INTO `kpi` VALUES (47, 35, 'D1 攻击行为', 0, '', 1712134914, NULL, 0, 0);
INSERT INTO `kpi` VALUES (48, 35, 'D2 抑郁状态', 0, '', 1712134981, NULL, 0, 0);
INSERT INTO `kpi` VALUES (49, 35, 'D3 意识水平', 0, '', 1712135035, NULL, 0, 0);
INSERT INTO `kpi` VALUES (50, 35, 'D4 沟通交流', 0, '', 1712135311, NULL, 0, 0);
INSERT INTO `kpi` VALUES (51, 35, 'D5 社会交往能力', 0, '', 1712135414, NULL, 0, 0);
INSERT INTO `kpi` VALUES (52, 35, 'D6 视力功能', 0, '', 1712135508, NULL, 0, 0);
INSERT INTO `kpi` VALUES (53, 35, 'D7 听力功能', 0, '', 1712135585, NULL, 0, 0);
INSERT INTO `kpi` VALUES (54, 34, 'C1 评估人员说出3样东西：“皮球”、“国旗”、“树木”，一秒说一项。 请评估对象记住这3样东西。过30秒，请评估对象重复一遍。', 0, '', 1712451600, NULL, 0, 0);
INSERT INTO `kpi` VALUES (55, 34, 'C2 现在是在哪一年？几月？什么季节？哪个城市？什么地方？ ', 0, '', 1712451766, NULL, 0, 0);
INSERT INTO `kpi` VALUES (56, 34, 'C3 请评估对象画一钟表盘面，把表示时间的数字写在正确的位置，再画上时针分针表示指定的时间', 0, '', 1712452478, NULL, 0, 0);
INSERT INTO `kpi` VALUES (57, 34, 'C4 （与第1题相隔3分钟）请您把刚才我们说过的3样东西说一遍？', 0, '', 1712452523, NULL, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
