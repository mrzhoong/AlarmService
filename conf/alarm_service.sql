/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50625
Source Host           : localhost:3306
Source Database       : alarm_service

Target Server Type    : MYSQL
Target Server Version : 50625
File Encoding         : 65001

Date: 2017-03-09 21:29:52
*/

CREATE DATABASE IF NOT EXISTS alarm_service;
USE alarm_service;

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for send_method
-- ----------------------------
DROP TABLE IF EXISTS `send_method`;
CREATE TABLE `send_method` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `strategy_infos_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of send_method
-- ----------------------------
INSERT INTO `send_method` VALUES ('3', '电话', '7');

-- ----------------------------
-- Table structure for send_record
-- ----------------------------
DROP TABLE IF EXISTS `send_record`;
CREATE TABLE `send_record` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `station` varchar(255) NOT NULL DEFAULT '',
  `equip_name` varchar(255) NOT NULL DEFAULT '',
  `event_time` bigint(20) NOT NULL DEFAULT '0',
  `event_level` int(11) NOT NULL DEFAULT '0',
  `tos` varchar(255) NOT NULL DEFAULT '',
  `send_type` varchar(255) NOT NULL DEFAULT '',
  `send_time` bigint(20) NOT NULL DEFAULT '0',
  `send_status` tinyint(1) NOT NULL DEFAULT '0',
  `content` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=119 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of send_record
-- ----------------------------
INSERT INTO `send_record` VALUES ('29', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488707239', '1', 'hello');
INSERT INTO `send_record` VALUES ('30', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488707240', '1', 'hello');
INSERT INTO `send_record` VALUES ('31', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488707277', '1', 'hello');
INSERT INTO `send_record` VALUES ('32', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488707277', '1', 'hello');
INSERT INTO `send_record` VALUES ('33', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488707277', '1', 'hello');
INSERT INTO `send_record` VALUES ('34', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488707277', '1', 'hello');
INSERT INTO `send_record` VALUES ('35', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488707413', '1', 'hello');
INSERT INTO `send_record` VALUES ('36', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488707413', '1', 'hello');
INSERT INTO `send_record` VALUES ('37', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488707413', '1', 'hello');
INSERT INTO `send_record` VALUES ('38', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488707413', '1', 'hello');
INSERT INTO `send_record` VALUES ('39', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488711468', '1', 'hello');
INSERT INTO `send_record` VALUES ('40', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488711471', '1', 'hello');
INSERT INTO `send_record` VALUES ('41', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488711472', '0', 'hello');
INSERT INTO `send_record` VALUES ('42', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488711473', '0', 'hello');
INSERT INTO `send_record` VALUES ('43', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488711501', '1', 'hello');
INSERT INTO `send_record` VALUES ('44', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488711501', '1', 'hello');
INSERT INTO `send_record` VALUES ('45', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488711501', '1', 'hello');
INSERT INTO `send_record` VALUES ('46', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488711501', '1', 'hello');
INSERT INTO `send_record` VALUES ('47', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488785291', '1', 'hello');
INSERT INTO `send_record` VALUES ('48', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488785291', '1', 'hello');
INSERT INTO `send_record` VALUES ('49', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488785291', '1', 'hello');
INSERT INTO `send_record` VALUES ('50', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488785291', '1', 'hello');
INSERT INTO `send_record` VALUES ('51', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488871259', '1', 'hello');
INSERT INTO `send_record` VALUES ('52', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488871259', '1', 'hello');
INSERT INTO `send_record` VALUES ('53', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488871259', '1', 'hello');
INSERT INTO `send_record` VALUES ('54', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488871259', '1', 'hello');
INSERT INTO `send_record` VALUES ('55', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488871687', '1', 'hello');
INSERT INTO `send_record` VALUES ('56', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488871687', '1', 'hello');
INSERT INTO `send_record` VALUES ('57', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488871687', '1', 'hello');
INSERT INTO `send_record` VALUES ('58', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488871687', '1', 'hello');
INSERT INTO `send_record` VALUES ('59', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488871754', '1', 'hello');
INSERT INTO `send_record` VALUES ('60', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488871754', '1', 'hello');
INSERT INTO `send_record` VALUES ('61', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488871754', '1', 'hello');
INSERT INTO `send_record` VALUES ('62', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488871754', '1', 'hello');
INSERT INTO `send_record` VALUES ('63', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488871755', '1', 'hello');
INSERT INTO `send_record` VALUES ('64', 'Ali', 'temp', '1488488411', '2', '', '/sms', '1488871755', '1', 'hello');
INSERT INTO `send_record` VALUES ('65', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488871756', '1', 'hello');
INSERT INTO `send_record` VALUES ('66', 'Ali', 'temp', '1488488411', '2', '', '/mail', '1488871756', '1', 'hello');
INSERT INTO `send_record` VALUES ('67', '科学城机房', '2', '1488488411', '0', '', '/sms', '1488872387', '1', '烟感状态异常');
INSERT INTO `send_record` VALUES ('68', '科学城机房', '2', '1488488411', '0', '', '/sms', '1488872387', '1', '烟感状态异常');
INSERT INTO `send_record` VALUES ('69', '科学城机房', '2', '1488488411', '0', '', '/mail', '1488872387', '1', '烟感状态异常');
INSERT INTO `send_record` VALUES ('70', '科学城机房', '2', '1488488411', '0', '', '/mail', '1488872387', '1', '烟感状态异常');
INSERT INTO `send_record` VALUES ('71', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488872465', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('72', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488872465', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('73', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488872465', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('74', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488872465', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('75', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488873097', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('76', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488873097', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('77', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488873097', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('78', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488873097', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('79', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874350', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('80', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874350', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('81', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874350', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('82', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874350', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('83', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874538', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('84', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874538', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('85', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874539', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('86', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874539', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('87', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874620', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('88', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874620', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('89', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874620', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('90', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874620', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('91', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874757', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('92', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874757', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('93', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874757', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('94', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874757', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('95', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874765', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('96', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/sms', '1488874765', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('97', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874765', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('98', 'bbbbbbb', 'ccccccccc', '1488488411', '2', '', '/mail', '1488874765', '1', 'aaaa');
INSERT INTO `send_record` VALUES ('99', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488874866', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('100', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488874866', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('101', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488874866', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('102', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488874866', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('103', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488875003', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('104', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488875003', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('105', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488875003', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('106', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488875003', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('107', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488880054', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('108', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488880054', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('109', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488880054', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('110', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488880054', '1', 'eeeee');
INSERT INTO `send_record` VALUES ('111', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488884834', '0', 'eeeee');
INSERT INTO `send_record` VALUES ('112', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488884835', '0', 'eeeee');
INSERT INTO `send_record` VALUES ('113', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488884836', '0', 'eeeee');
INSERT INTO `send_record` VALUES ('114', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488884837', '0', 'eeeee');
INSERT INTO `send_record` VALUES ('115', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488895077', '0', 'eeeee');
INSERT INTO `send_record` VALUES ('116', 'eeee', 'eeee', '1488488411', '2', '', '/sms', '1488895078', '0', 'eeeee');
INSERT INTO `send_record` VALUES ('117', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488895079', '0', 'eeeee');
INSERT INTO `send_record` VALUES ('118', 'eeee', 'eeee', '1488488411', '2', '', '/mail', '1488895080', '0', 'eeeee');

-- ----------------------------
-- Table structure for send_to
-- ----------------------------
DROP TABLE IF EXISTS `send_to`;
CREATE TABLE `send_to` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `mobile` varchar(255) NOT NULL DEFAULT '',
  `we_chat` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `strategy_infos_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of send_to
-- ----------------------------
INSERT INTO `send_to` VALUES ('3', 'sam2', '18503036522', 'eqwead_ew1_2', 'xxx@zktz.com', '7');

-- ----------------------------
-- Table structure for send_type
-- ----------------------------
DROP TABLE IF EXISTS `send_type`;
CREATE TABLE `send_type` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `send_type` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of send_type
-- ----------------------------
INSERT INTO `send_type` VALUES ('1', '电话');
INSERT INTO `send_type` VALUES ('2', '短信');
INSERT INTO `send_type` VALUES ('3', '微信');
INSERT INTO `send_type` VALUES ('4', '邮件');

-- ----------------------------
-- Table structure for strategy
-- ----------------------------
DROP TABLE IF EXISTS `strategy`;
CREATE TABLE `strategy` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `strategy_name` varchar(255) NOT NULL DEFAULT '',
  `strategy_condition` varchar(255) NOT NULL DEFAULT '',
  `send_method` varchar(255) NOT NULL DEFAULT '',
  `send_to` varchar(255) NOT NULL DEFAULT '',
  `strategy_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of strategy
-- ----------------------------
INSERT INTO `strategy` VALUES ('1', '运维', '实时', '短信', '张三，李四，王五', '1,2,3,4,5');
INSERT INTO `strategy` VALUES ('2', '科技部', '实时', '微信', '1,2,3,4', '1,2');
INSERT INTO `strategy` VALUES ('3', '值班', '实时', '4', '1,2', '1');

-- ----------------------------
-- Table structure for strategy_info
-- ----------------------------
DROP TABLE IF EXISTS `strategy_info`;
CREATE TABLE `strategy_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `condition` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of strategy_info
-- ----------------------------
INSERT INTO `strategy_info` VALUES ('7', '运维组', 'Condition');

-- ----------------------------
-- Table structure for strategy_list
-- ----------------------------
DROP TABLE IF EXISTS `strategy_list`;
CREATE TABLE `strategy_list` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `strategy_infos_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of strategy_list
-- ----------------------------
INSERT INTO `strategy_list` VALUES ('7', '暖通', '7');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `tel` varchar(255) NOT NULL DEFAULT '',
  `email` varchar(255) NOT NULL DEFAULT '',
  `wechat` varchar(255) NOT NULL DEFAULT '',
  `organization_code` bigint(20) NOT NULL DEFAULT '0',
  `organization_name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', '张三', '18500001111', 'zhangsan@zktz.com', 'wechat1', '1', '科技部');
INSERT INTO `user` VALUES ('2', '李四', '18655556666', 'lisi@zktz.com', 'wechat2', '2', '信息部');
INSERT INTO `user` VALUES ('3', '一笑', '18522221111', 'yixiao@zktz.com', 'wechat3', '3', '运维');
INSERT INTO `user` VALUES ('4', '笑笑', '13255556666', 'xiao@zktc.com', 'wechat4', '4', '中科');
INSERT INTO `user` VALUES ('5', '珊瑚', '13655554444', 'hu@zktz.com', 'wechat5', '5', '运维外包');
