/*
Navicat MySQL Data Transfer

Source Server         : 林峰会展
Source Server Version : 50740
Source Host           : localhost:3306
Source Database       : linfeng

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2023-10-25 14:48:37
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for userpay
-- ----------------------------
DROP TABLE IF EXISTS `userpay`;
CREATE TABLE `userpay` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `openid` varchar(255) DEFAULT NULL,
  `uid` bigint(20) DEFAULT NULL,
  `exid` bigint(20) DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `codeurl` varchar(255) DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `jid` bigint(20) DEFAULT NULL,
  `type` varchar(200) DEFAULT NULL,
  `out_trade_no` varchar(32) DEFAULT NULL,
  `ispay` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否启用 默认1 否 2 是',
  `pay_time` bigint(20) DEFAULT NULL,
  `success_time` varchar(255) DEFAULT NULL,
  `transaction_id` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=245 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of userpay
-- ----------------------------
INSERT INTO `userpay` VALUES ('183', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', '', '2023-06-05 18:03:43', '2023-06-05 18:03:43', '0', '', '202306051803439190', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('184', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=3Wl0l3Ezz', '2023-06-05 18:04:26', '2023-06-05 18:04:28', '0', '', '202306051804265252', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('185', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=4OPwOhszz', '2023-06-05 18:06:11', '2023-06-05 18:06:12', '0', '', '202306051806118782', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('186', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '2029', '30', 'weixin://wxpay/bizpayurl?pr=oUwKShwzz', '2023-06-05 18:07:45', '2023-06-05 18:07:45', '0', '', '202306051807457333', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('187', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=pmV0Pffzz', '2023-06-05 18:11:42', '2023-06-05 18:11:43', '0', '', '202306051811427795', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('188', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=xdMnwjmzz', '2023-06-05 18:14:11', '2023-06-05 18:14:12', '0', '', '202306051814114727', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('189', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=2b03qBWzz', '2023-06-05 18:26:33', '2023-06-05 18:26:35', '0', '', '202306051826333496', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('190', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=D8YwAIEzz', '2023-06-05 18:27:38', '2023-06-05 18:27:39', '0', '', '202306051827388928', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('191', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=vTB56XEzz', '2023-06-05 18:29:14', '2023-06-05 18:29:16', '0', '', '202306051829149985', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('192', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=HVgzOT3zz', '2023-06-05 18:30:28', '2023-06-05 18:30:29', '0', '', '202306051830287142', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('193', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=l3fJbSzzz', '2023-06-05 18:31:00', '2023-06-05 18:31:02', '0', '', '202306051831003907', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('194', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '2029', '30', 'weixin://wxpay/bizpayurl?pr=rRj2Jb7zz', '2023-06-05 18:41:15', '2023-06-05 18:41:15', '0', '', '202306051841159596', '0', '0', '0', '');
INSERT INTO `userpay` VALUES ('195', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '2030', '30', 'weixin://wxpay/bizpayurl?pr=L7ya0dSzz', '2023-06-05 18:45:27', '2023-06-05 18:45:41', '0', '', '202306051845271257', '2', '0', '2023-06-05T18:45:41+08:00', '');
INSERT INTO `userpay` VALUES ('196', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '2030', '30', 'weixin://wxpay/bizpayurl?pr=AuIr55izz', '2023-06-05 18:45:44', '2023-06-05 18:45:44', '0', '', '202306051845445369', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('197', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=7xcxGAlzz', '2023-06-06 08:24:03', '2023-06-06 08:24:05', '0', '', '202306060824034676', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('198', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=NQByqbczz', '2023-06-06 08:24:59', '2023-06-06 08:25:01', '0', '', '202306060824595479', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('199', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=FmkOimwzz', '2023-06-06 10:47:13', '2023-06-06 10:47:15', '0', '', '202306061047138474', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('200', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=cVtWXwfzz', '2023-06-06 10:53:01', '2023-06-06 10:53:02', '0', '', '202306061053014950', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('201', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '2152', '30', 'weixin://wxpay/bizpayurl?pr=jcgLsxdzz', '2023-06-06 11:15:34', '2023-06-06 11:15:54', '0', '', '202306061115347149', '2', '0', '2023-06-06T11:15:53+08:00', '');
INSERT INTO `userpay` VALUES ('202', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '3882', '50', 'weixin://wxpay/bizpayurl?pr=Fhr7dZazz', '2023-06-08 18:10:59', '2023-06-08 18:11:00', '0', '', '202306081810594502', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('203', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '3882', '50', 'weixin://wxpay/bizpayurl?pr=5GStiNpzz', '2023-06-08 18:12:50', '2023-06-08 18:13:27', '0', '', 'HK202306081812501823', '2', '0', '2023-06-08T18:13:27+08:00', '');
INSERT INTO `userpay` VALUES ('204', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '2016', '30', 'weixin://wxpay/bizpayurl?pr=oWiVBznzz', '2023-06-12 18:12:25', '2023-06-12 18:12:38', '0', '', 'ZH202306121812259213', '2', '0', '2023-06-12T18:12:37+08:00', '');
INSERT INTO `userpay` VALUES ('205', 'oNbgG55Qa8Del8sWkLPd7hMw45Zg', '9', '7273', '0', '', '2023-06-28 10:51:42', '2023-06-28 10:51:42', '0', '', 'ZH202306281051426247', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('206', 'oNbgG55Qa8Del8sWkLPd7hMw45Zg', '9', '6351', '0', '', '2023-06-28 11:12:11', '2023-06-28 11:12:11', '0', '', 'ZH202306281112111318', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('207', 'oNbgG55Qa8Del8sWkLPd7hMw45Zg', '9', '7686', '0', '', '2023-06-28 11:12:24', '2023-06-28 11:12:24', '0', '', 'ZH202306281112245550', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('208', 'oNbgG55Qa8Del8sWkLPd7hMw45Zg', '9', '5892', '0', '', '2023-06-28 11:14:42', '2023-06-28 11:14:42', '0', '', 'ZH202306281114427942', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('209', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '3977', '0', '', '2023-06-28 11:21:23', '2023-06-28 11:21:24', '0', '', 'ZH202306281121239884', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('210', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '3977', '0', '', '2023-06-28 11:21:35', '2023-06-28 11:21:36', '0', '', 'ZH202306281121351566', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('211', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '5617', '0', '', '2023-06-28 11:25:49', '2023-06-28 11:25:50', '0', '', 'ZH202306281125495050', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('212', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '5617', '30', 'weixin://wxpay/bizpayurl?pr=jSw25Obzz', '2023-06-28 17:12:06', '2023-06-28 17:12:07', '0', '', 'ZH202306281712061849', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('213', 'oNbgG5wPixras3PqloMHnB6YOsuo', '8', '3333', '30', 'weixin://wxpay/bizpayurl?pr=670QA4zzz', '2023-06-29 09:37:58', '2023-06-29 09:37:58', '0', '', 'ZH202306290937585628', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('214', 'oNbgG5wPixras3PqloMHnB6YOsuo', '8', '3333', '30', 'weixin://wxpay/bizpayurl?pr=URZCmdkzz', '2023-06-29 09:38:07', '2023-06-29 09:38:08', '0', '', 'ZH202306290938073421', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('215', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '3333', '30', 'weixin://wxpay/bizpayurl?pr=uvEloSFzz', '2023-06-29 15:32:45', '2023-06-29 15:32:46', '0', '', 'ZH202306291532452750', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('216', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '3333', '30', 'weixin://wxpay/bizpayurl?pr=E6ZVRywzz', '2023-06-29 15:33:39', '2023-06-29 15:34:00', '0', '', 'ZH202306291533399269', '2', '0', '2023-06-29T15:33:59+08:00', '');
INSERT INTO `userpay` VALUES ('217', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '5814', '30', 'weixin://wxpay/bizpayurl?pr=VPzlxKXzz', '2023-06-29 16:11:18', '2023-06-29 16:11:19', '0', '', 'ZH202306291611186642', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('218', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '5617', '30', 'weixin://wxpay/bizpayurl?pr=tfAUz6czz', '2023-06-29 18:19:42', '2023-06-29 18:19:43', '0', '', 'ZH202306291819422756', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('219', 'oNbgG55Qa8Del8sWkLPd7hMw45Zg', '9', '3319', '30', 'weixin://wxpay/bizpayurl?pr=X38lVDCzz', '2023-07-07 10:44:09', '2023-07-07 10:44:28', '0', '', 'ZH202307071044099090', '2', '0', '2023-07-07T10:44:27+08:00', '');
INSERT INTO `userpay` VALUES ('220', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '4020', '600', 'weixin://wxpay/bizpayurl?pr=HgofxyKzz', '2023-07-16 12:27:21', '2023-07-16 12:27:22', '0', '', 'ZH202307161227216888', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('221', 'oNbgG55Qa8Del8sWkLPd7hMw45Zg', '9', '5651', '30', 'weixin://wxpay/bizpayurl?pr=LBAfbZazz', '2023-08-07 16:50:20', '2023-08-07 16:50:46', '0', '', 'ZH202308071650208444', '2', '0', '2023-08-07T16:50:46+08:00', '');
INSERT INTO `userpay` VALUES ('222', 'oNbgG5wPixras3PqloMHnB6YOsuo', '8', '3405', '30', 'weixin://wxpay/bizpayurl?pr=6VxXrNAzz', '2023-08-08 15:08:47', '2023-08-08 15:08:48', '0', '', 'ZH202308081508474811', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('223', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '4094', '30', 'weixin://wxpay/bizpayurl?pr=cHvmhsozz', '2023-08-21 11:04:04', '2023-08-21 11:04:23', '0', '', 'ZH202308211104042261', '2', '0', '2023-08-21T11:04:22+08:00', '');
INSERT INTO `userpay` VALUES ('224', 'oNbgG5_oo7L8Hna5LXTWHUsNEwr8', '14', '4837', '30', 'weixin://wxpay/bizpayurl?pr=PF4EyFRzz', '2023-08-21 11:07:15', '2023-08-21 11:07:16', '0', '', 'ZH202308211107158077', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('225', 'oNbgG5_oo7L8Hna5LXTWHUsNEwr8', '14', '4837', '30', 'weixin://wxpay/bizpayurl?pr=LpnbX13zz', '2023-08-21 11:07:49', '2023-08-21 11:08:06', '0', '', 'ZH202308211107493383', '2', '0', '2023-08-21T11:08:05+08:00', '');
INSERT INTO `userpay` VALUES ('226', 'oNbgG59RVeBUVw0OEwMoVPYmwIKM', '15', '7137', '30', 'weixin://wxpay/bizpayurl?pr=D4LdkyFzz', '2023-08-28 11:18:51', '2023-08-28 11:19:15', '0', '', 'ZH202308281118511212', '2', '0', '2023-08-28T11:19:14+08:00', '');
INSERT INTO `userpay` VALUES ('227', 'oNbgG59RVeBUVw0OEwMoVPYmwIKM', '15', '7137', '30', 'weixin://wxpay/bizpayurl?pr=2aePx8Szz', '2023-08-28 11:21:59', '2023-08-28 11:21:59', '0', '', 'ZH202308281121593760', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('228', 'oNbgG5_EeaXYkNMOMDDpjoxxVVvI', '16', '7137', '30', 'weixin://wxpay/bizpayurl?pr=nQhx1eszz', '2023-08-28 11:41:38', '2023-08-28 11:41:53', '0', '', 'ZH202308281141382592', '2', '0', '2023-08-28T11:41:52+08:00', '');
INSERT INTO `userpay` VALUES ('229', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '5801', '30', 'weixin://wxpay/bizpayurl?pr=odGBumDzz', '2023-08-28 15:33:31', '2023-08-28 15:33:33', '0', '', 'ZH202308281533314481', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('230', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '4328', '30', 'weixin://wxpay/bizpayurl?pr=XJuis8Kzz', '2023-08-30 11:37:35', '2023-08-30 11:37:37', '0', '', 'ZH202308301137351899', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('231', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '4328', '30', 'weixin://wxpay/bizpayurl?pr=Trya9IEzz', '2023-08-30 11:39:10', '2023-08-30 11:39:11', '0', '', 'ZH202308301139105171', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('232', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '8825', '30', 'weixin://wxpay/bizpayurl?pr=XrFg8alzz', '2023-08-30 11:58:27', '2023-08-30 11:58:28', '0', '', 'ZH202308301158273972', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('233', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '8825', '30', 'weixin://wxpay/bizpayurl?pr=UVSJTASzz', '2023-08-30 11:58:50', '2023-08-30 11:58:50', '0', '', 'ZH202308301158503478', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('234', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '4328', '30', 'weixin://wxpay/bizpayurl?pr=uh2xL2Vzz', '2023-08-30 14:42:14', '2023-08-30 14:42:16', '0', '', 'ZH202308301442143409', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('235', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '4328', '30', 'weixin://wxpay/bizpayurl?pr=nxXUBQ1zz', '2023-08-30 15:53:12', '2023-08-30 15:53:16', '0', '', 'ZH202308301553121959', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('236', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '4328', '30', 'weixin://wxpay/bizpayurl?pr=DGjEXA4zz', '2023-08-30 15:55:33', '2023-08-30 15:55:38', '0', '', 'ZH202308301555331226', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('237', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '4328', '30', 'weixin://wxpay/bizpayurl?pr=azNgNrWzz', '2023-08-30 15:58:04', '2023-08-30 15:58:06', '0', '', 'ZH202308301558041912', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('238', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '4328', '30', 'weixin://wxpay/bizpayurl?pr=wupgtKZzz', '2023-08-31 17:34:49', '2023-08-31 17:34:50', '0', '', 'ZH202308311734492742', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('239', 'oNbgG5-1NUNuS6nYDZWA8DLmR2LE', '4', '4323', '30', 'weixin://wxpay/bizpayurl?pr=yg3nB1nzz', '2023-08-31 17:47:25', '2023-09-04 08:45:33', '0', '', 'ZH202308311747258124', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('240', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '4426', '30', 'weixin://wxpay/bizpayurl?pr=RliwBF8zz', '2023-09-14 14:06:12', '2023-09-14 14:06:13', '0', '', 'ZH202309141406124552', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('241', 'oNbgG52oRb1fj7deRDRjY63usRWQ', '20', '6715', '30', 'weixin://wxpay/bizpayurl?pr=1KnaG3vzz', '2023-09-19 22:52:05', '2023-09-19 22:52:54', '0', '', 'ZH202309192252052515', '2', '0', '2023-09-19T22:52:53+08:00', '');
INSERT INTO `userpay` VALUES ('242', 'oNbgG5977BGMvUubqrJXw-DD3NpQ', '7', '4550', '600', 'weixin://wxpay/bizpayurl?pr=ZyY3l7Izz', '2023-09-19 23:40:26', '2023-09-19 23:40:27', '0', '', 'ZH202309192340262367', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('243', 'oNbgG58fmji_SsLgyriuzpqK9W54', '21', '6984', '30', 'weixin://wxpay/bizpayurl?pr=jJRjmiWzz', '2023-09-30 21:25:35', '2023-09-30 21:25:36', '0', '', 'ZH202309302125358315', '0', '0', '', '');
INSERT INTO `userpay` VALUES ('244', 'oNbgG55Qa8Del8sWkLPd7hMw45Zg', '9', '8578', '30', 'weixin://wxpay/bizpayurl?pr=8gmZXRNzz', '2023-10-20 09:42:51', '2023-10-20 09:43:20', '0', '', 'ZH202310200942514802', '2', '0', '2023-10-20T09:43:19+08:00', '');
