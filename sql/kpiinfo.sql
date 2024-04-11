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

 Date: 07/04/2024 09:22:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for kpiinfo
-- ----------------------------
DROP TABLE IF EXISTS `kpiinfo`;
CREATE TABLE `kpiinfo`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `kpiid` bigint(20) NULL DEFAULT NULL COMMENT '姓名',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '指标名称',
  `score` int(10) NULL DEFAULT NULL COMMENT '分数',
  `created` int(11) NULL DEFAULT NULL COMMENT '创建时间',
  `updated` int(11) NULL DEFAULT NULL COMMENT '修改时间',
  `weigh` int(11) NULL DEFAULT NULL COMMENT '排序',
  `status` int(11) NULL DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 137 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of kpiinfo
-- ----------------------------
INSERT INTO `kpiinfo` VALUES (4, 10, '进食中需要大量接触式协助，经常（每周一次及以上）呛咳', 1, 1711338811, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (5, 10, '完全依赖他人协助进食，或吞咽困难，或留置营养管', 0, 1711358962, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (6, 10, '进食中需要少量接触式协助偶尔（每月一次及以上）呛咳', 2, 1711359178, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (7, 10, '在他人指导或提示下完成，或独立使用辅具，没有呛咳', 3, 1711359621, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (8, 10, '独立使用餐具将食品送进口中并咽下，没有呛咳', 4, 1711413808, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (9, 12, '完全依靠他人协助，且不能给予配合', 0, 1711413866, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (10, 0, '主要依靠他人协助，自身能给予配合', 1, 1711413922, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (11, 0, '主要依靠他人协助，自身能给予配合', 1, 1711413960, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (13, 12, '主要依靠他人协助，自身能给予配合', 1, 1711414037, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (14, 0, '需要他人协助，但以自身完成为主', 2, 1711414087, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (15, 12, '需要他人协助，但以自身完成为主', 2, 1711414141, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (16, 0, '在他人指导或提示下完成', 3, 1711414172, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (17, 0, '', 0, 1711414176, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (18, 12, '在他人指导或提示下完成', 3, 1711414228, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (19, 0, '独立完成，不需要协助', 4, 1711414250, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (20, 12, '独立完成，不需要协助', 4, 1711414281, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (21, 14, '在他人指导或提示下完成', 3, 1711414320, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (22, 0, '独立完成，不需要协助', 4, 1711414340, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (23, 14, '独立完成，不需要协助', 4, 1711414368, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (24, 0, '需要他人协助，但以自身完成为主', 2, 1711414404, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (25, 14, '需要他人协助，但以自身完成为主', 2, 1711414416, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (26, 14, '主要依靠他人协助，自身能给予配合', 1, 1711414450, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (27, 0, '完全依靠他人协助，且不能给予配合', 0, 1711414481, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (28, 0, '完全依靠他人协助，且不能给予配合', 0, 1711414485, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (29, 14, '完全依靠他人协助，且不能给予配合', 0, 1711414492, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (30, 13, '独立完成，不需要协助', 4, 1711414697, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (31, 0, '在他人指导或提示下完成', 3, 1711414726, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (32, 13, '在他人指导或协助下完成', 3, 1711414768, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (33, 0, '需要依靠他人协助，但以自身完成为主', 2, 1711414796, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (34, 0, '需要依靠他人协助，但以自身完成为主', 2, 1711414802, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (35, 13, '需要依靠他人协助，但以自身完成为主', 2, 1711414820, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (36, 0, '主要依靠他人协助，自身能给予配合', 1, 1711414853, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (37, 13, '主要依靠他人协助，自身能给予配合', 1, 1711414867, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (38, 0, '完全依靠他人协助，且不能给予配合', 0, 1711414894, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (39, 13, '完全依靠他人协助，且不能给予配合', 0, 1711414904, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (40, 21, '完全依靠他人协助，且不能给予配合', 0, 1711414927, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (41, 0, '主要依靠他人协助，自身能给予配合', 1, 1711414938, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (42, 21, '主要依靠他人协助，自身能给予配合', 1, 1711414952, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (43, 0, '需要依靠他人协助，但以自身完成为主', 2, 1711414963, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (44, 21, '需要依靠他人协助，但以自身完成为主', 2, 1711414975, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (45, 0, '在他人指导或协助下完成', 3, 1711414998, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (46, 21, '在他人指导或协助下完成', 3, 1711415014, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (47, 21, '独立完成，不需要协助', 4, 1711415044, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (49, 36, '可独立进食', 10, 1712132639, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (50, 0, '需部分帮助', 5, 1712132653, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (51, 37, '准备好洗澡水后，可自己独立完成洗澡过程', 5, 1712134018, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (52, 0, '在洗澡过程中需他人帮助', 0, 1712134030, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (53, 0, '在洗澡过程中需他人帮助', 0, 1712134039, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (54, 37, '在洗澡过程中需他人帮助', 0, 1712134070, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (55, 36, '需部分帮助(进食过程中需要一定帮助，如协助把持餐具)', 5, 1712134090, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (56, 36, '需极大帮助或完全依赖他人，或有留置营养管', 0, 1712134102, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (57, 38, '可自己独立完成', 5, 1712134129, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (58, 38, '需他人帮助', 0, 1712134139, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (59, 39, '可独立完成', 10, 1712134162, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (60, 0, '需部分帮助(能自己穿脱，但需他人帮助整理衣物、系扣/鞋带拉拉链)', 5, 1712134175, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (61, 39, '需部分帮助(能自己穿脱，但需他人帮助整理衣物、系扣/鞋带拉拉链)', 5, 1712134184, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (62, 39, '需极大帮助或完全依赖他人', 0, 1712134199, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (63, 40, '可控制大便', 10, 1712134218, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (64, 40, '偶尔失控(每周<1次)，或需要他人提示', 5, 1712134232, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (65, 40, '完全失控', 0, 1712134242, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (66, 42, '可控制小便', 10, 1712134323, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (67, 42, '偶尔失控(每天<1次，但每周>1次)，或需要他人提示', 5, 1712134339, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (68, 42, '完全失控，或留置导尿管', 0, 1712134350, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (69, 43, '可独立完成', 10, 1712134408, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (70, 43, '需部分帮助(需他人搀扶去厕所、需他人帮忙冲水或整理衣裤等)', 5, 1712134428, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (71, 43, '需极大帮助或完全依赖他人', 0, 1712134437, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (72, 44, '可独立完成', 15, 1712134489, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (73, 0, '需部分帮助(需他人搀扶或使用拐杖)', 10, 1712134506, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (74, 44, '需部分帮助(需他人搀扶或使用拐杖)', 10, 1712134521, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (75, 44, '需极大帮助(较大程度上依赖他人搀扶和帮助)', 5, 1712134536, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (76, 44, '完全依赖他人', 0, 1712134547, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (77, 45, '可独立在平地上行走45m', 15, 1712134594, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (78, 45, '需部分帮助(因肢体残疾、平衡能力差、过度衰弱、视力等问题,在一定程度上需他人地搀扶或使用拐杖、助行器等辅助用具)', 10, 1712134613, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (79, 45, '需极大帮助(因肢体残疾、平衡能力差、过度衰弱、视力等问题,在较大程度上依赖他人搀扶，或坐在轮椅上自行移动)', 5, 1712134685, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (80, 45, '完全依赖他人', 0, 1712134699, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (81, 46, '可独立上下楼梯(连续上下10-15个台阶)', 10, 1712134735, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (82, 46, '需部分帮助(需他人搀扶，或扶着楼梯、使用拐等)', 5, 1712134750, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (83, 46, '需极大帮助或完全依赖他人', 0, 1712134762, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (84, 47, '无身体攻击行为(如打/踢/推/咬/抓/摔东西)和语言攻击行为(如骂人、语言 威胁、尖叫)', 2, 1712134940, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (85, 47, '每月有几次身体攻击行为，或每周有几次语言攻击行为', 1, 1712134956, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (86, 47, '每周有几次身体攻击行为，或每日有语言攻击行为', 0, 1712134967, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (87, 48, '无', 2, 1712134998, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (88, 48, '情绪低落、不爱说话、不爱梳洗、不爱活动', 1, 1712135011, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (89, 48, '有自杀念头或自杀行为', 0, 1712135023, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (90, 49, '神志清醒，对周围环境警觉', 3, 1712135046, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (91, 49, '嗜睡，表现为睡眠状态过度延长。当呼唤或推动其肢体时可唤醒，并能进行正确的交谈或执行指令，停止刺激后又继续入睡', 2, 1712135057, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (94, 49, '昏睡，一般的外界刺激不能使其觉醒，给予较强烈的刺激时可有短时的意识清醒:醒后可简短回答提问，当刺激减弱后又很快进入睡眠状态', 1, 1712135298, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (93, 49, '昏迷，处于浅昏迷时对疼痛刺激有回避和痛苦表情;处于深昏迷时对刺激无反应(若评定为昏迷，直接评定为重度失能，可不进行以下项目的评估)', 0, 1712135233, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (95, 50, '无困难，能与他人正常沟通和交流', 3, 1712135327, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (96, 50, '能够表达自己的需要及理解别人的话，但需要增加时间或给予帮助', 2, 1712135342, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (97, 50, '表达需要或理解有困难，需频繁重复或简化口头表达', 1, 1712135362, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (98, 50, '不能表达需要或理解他人的话', 0, 1712135380, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (99, 51, '参与社会，在社会环境有一定的适应能力，待人接物恰当', 4, 1712135429, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (100, 51, '能适应单纯环境，主动接触人，初见面时难让人发现智力问题，不能理解隐喻语', 3, 1712135446, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (101, 51, '脱离社会，可被动接触，不会主动待人，谈话中很多不适词句，容易上当受骗', 2, 1712135460, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (102, 0, '勉强可与人交往，谈吐内容不清楚，表情不恰当', 1, 1712135470, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (103, 51, '勉强可与人交往，谈吐内容不清楚，表情不恰当', 1, 1712135478, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (104, 51, '难以与人接触', 0, 1712135491, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (105, 52, '能看清书报上的标准字体', 4, 1712135524, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (106, 52, '能看清楚大字体，但看不清书报上的标准字体', 3, 1712135535, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (107, 52, '视力有限，看不清报纸大标题，但能辨认物体', 2, 1712135547, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (108, 52, '辨认物体有困难，但眼睛能跟随物体移动，只能看到光、颜色和形状', 1, 1712135558, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (109, 52, '没有视力，眼睛不能跟随物体移动', 0, 1712135569, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (110, 53, '可正常交谈，能听到电视、电话、门铃的声音', 4, 1712135597, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (111, 53, '在轻声说话或说话距离超过2米时听不清', 3, 1712135609, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (112, 53, '正常交流有些困难，需在安静的环静或大声说话才能听到', 2, 1712135621, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (113, 53, '讲话者大声说话或说话很慢，才能部分听见', 1, 1712135634, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (114, 53, '完全听不见', 0, 1712135647, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (115, 54, '说出一个1分', 1, 1712451648, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (116, 54, '说出两个2分', 2, 1712451674, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (117, 54, '说出三个3分', 3, 1712451707, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (118, 54, '一样都记不住0分', 0, 1712451750, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (119, 55, '回答一个1分', 1, 1712451794, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (120, 55, '回答两个2分', 2, 1712451814, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (121, 0, '回答三个3分', 3, 1712452390, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (122, 55, '回答三个3分', 3, 1712452408, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (123, 55, '回答四个4分', 4, 1712452425, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (124, 55, '回答五个5分', 5, 1712452442, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (125, 55, '一个都答不出0分', 0, 1712452460, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (126, 57, '说出一个1分', 1, 1712452544, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (127, 57, '说出两个2分', 2, 1712452558, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (128, 57, '说出三个3分', 3, 1712452574, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (129, 57, '说不出来0分', 0, 1712452601, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (130, 56, '1分', 1, 1712452829, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (131, 56, '2分', 2, 1712452840, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (132, 0, '3分', 3, 1712452847, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (133, 56, '3分', 3, 1712452869, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (134, 56, '4分', 4, 1712452881, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (135, 56, '5分', 5, 1712452893, NULL, 0, 0);
INSERT INTO `kpiinfo` VALUES (136, 56, '0分', 0, 1712452902, NULL, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
