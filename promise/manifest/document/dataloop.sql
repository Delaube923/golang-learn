/*
Navicat MySQL Data Transfer

Source Server         : server84
Source Server Version : 80030
Source Host           : 118.25.144.84:6607
Source Database       : dataloop

Target Server Type    : MYSQL
Target Server Version : 80030
File Encoding         : 65001

Date: 2022-08-16 17:45:04
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for carinfo
-- ----------------------------
DROP TABLE IF EXISTS `carinfo`;
CREATE TABLE `carinfo` (
  `vehicle_number` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '车辆编号',
  `vehicle_model` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '车辆型号',
  `slice_name` varchar(255) NOT NULL DEFAULT '' COMMENT '切片名称',
  `slice_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '切片存储地址',
  `slice_size` int NOT NULL COMMENT '切片大小',
  `slice_md5` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '切片md5值',
  `eventId` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '事件id',
  `vin` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '车架号',
  `vehicle_use` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '车辆用途',
  `vehicle_area` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '所属地区',
  KEY `eventid` (`eventId`) USING BTREE,
  CONSTRAINT `carinfo_ibfk_1` FOREIGN KEY (`eventId`) REFERENCES `eventmax` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of carinfo
-- ----------------------------

-- ----------------------------
-- Table structure for eventmax
-- ----------------------------
DROP TABLE IF EXISTS `eventmax`;
CREATE TABLE `eventmax` (
  `id` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件id',
  `event_time` datetime NOT NULL COMMENT '事件发生日期/时间',
  `trigger_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件触发方式',
  `event_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件类型',
  `event_description` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件描述',
  `start_time` datetime NOT NULL COMMENT '切片数据开始时间',
  `duration` int NOT NULL COMMENT '切片数据的持续时间(s)',
  `vehicle_number` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '车辆编号',
  PRIMARY KEY (`id`),
  KEY `Id` (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of eventmax
-- ----------------------------
INSERT INTO `eventmax` VALUES ('1', '2022-08-19 09:37:02', '1', '1', '1', '2022-08-20 09:37:43', '1', '');
INSERT INTO `eventmax` VALUES ('2', '2022-08-10 10:49:30', '2', '2', '2', '2022-08-11 10:49:36', '2', '');

-- ----------------------------
-- Table structure for eventmiddle
-- ----------------------------
DROP TABLE IF EXISTS `eventmiddle`;
CREATE TABLE `eventmiddle` (
  `id` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件id',
  `event_time` datetime NOT NULL COMMENT '事件发生日期/时间',
  `trigger_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件触发方式',
  `event_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件类型',
  `event_description` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件描述',
  `start_time` datetime NOT NULL COMMENT '数据开始时间',
  `duration` int NOT NULL COMMENT '切片数据的持续时间(s)',
  `vehicle_number` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '车辆编号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of eventmiddle
-- ----------------------------
INSERT INTO `eventmiddle` VALUES ('1', '2022-08-28 16:35:42', '1', '1', '1', '2022-08-17 16:35:55', '1', '');

-- ----------------------------
-- Table structure for eventsmall
-- ----------------------------
DROP TABLE IF EXISTS `eventsmall`;
CREATE TABLE `eventsmall` (
  `id` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件id',
  `event_time` datetime NOT NULL COMMENT '事件发生日期/时间',
  `trigger_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件触发方式',
  `event_type` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件类型',
  `event_description` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '事件描述',
  `start_time` datetime NOT NULL COMMENT '切片数据开始时间',
  `duration` int NOT NULL COMMENT '切片数据的持续时间(s)',
  `vehicle_number` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '车辆编号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of eventsmall
-- ----------------------------
INSERT INTO `eventsmall` VALUES ('1', '2022-08-28 16:35:42', '1', '1', '1', '2022-08-17 16:35:55', '1', '');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(255) NOT NULL COMMENT '用户唯一标识',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户姓名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `picture` varchar(255) DEFAULT NULL COMMENT '用户头像 url',
  `email` varchar(255) DEFAULT NULL COMMENT '用户邮箱'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'admin', '', null, null);

-- ----------------------------
-- Event structure for month_delete
-- ----------------------------
DROP EVENT IF EXISTS `month_delete`;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` EVENT `month_delete` ON SCHEDULE EVERY 1 MONTH STARTS '2022-08-08 20:06:55' ON COMPLETION NOT PRESERVE ENABLE COMMENT '每一个月执行一次\r\n向总表添加新数据\r\n清除eventmiddle（月表）一个月之前的数据（从2022-08-08开始）\r\n' DO BEGIN
INSERT INTO eventmax SELECT * FROM eventmiddle WHERE not EXISTS( select * from eventmax WHERE eventmax.id=eventmiddle.id);
delete from eventsmiddle where DATE(event_time)<=DATE(DATE_SUB(NOW(),INTERVAL 1 MONTH));
END
;;
DELIMITER ;

-- ----------------------------
-- Event structure for month_insert
-- ----------------------------
DROP EVENT IF EXISTS `month_insert`;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` EVENT `month_insert` ON SCHEDULE EVERY 10 DAY STARTS '2022-08-08 20:06:55' ON COMPLETION NOT PRESERVE DISABLE COMMENT '每隔10天向总表写入新数据' DO INSERT INTO eventmax SELECT * FROM eventmiddle WHERE not EXISTS( select * from eventmax WHERE eventmax.id=eventmiddle.id)
;;
DELIMITER ;

-- ----------------------------
-- Event structure for schedule_delete
-- ----------------------------
DROP EVENT IF EXISTS `schedule_delete`;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` EVENT `schedule_delete` ON SCHEDULE EVERY 1 WEEK STARTS '2022-08-08 20:06:55' ON COMPLETION NOT PRESERVE ENABLE COMMENT '每周执行一次\r\n向月表添加数据\r\n清除eventsmall（周表）一周前的数据\r\n' DO BEGIN
INSERT INTO eventmiddle SELECT * FROM eventsmall WHERE not EXISTS( select * from eventmiddle WHERE eventmiddle.id=eventsmall.id);
delete from eventsmall where DATE(event_time)<=DATE(DATE_SUB(NOW(),INTERVAL 1 WEEK));
END
;;
DELIMITER ;

-- ----------------------------
-- Event structure for week_insert
-- ----------------------------
DROP EVENT IF EXISTS `week_insert`;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` EVENT `week_insert` ON SCHEDULE EVERY 3 DAY STARTS '2022-08-08 20:06:55' ON COMPLETION NOT PRESERVE ENABLE COMMENT '每隔3天向月表写入新数据' DO INSERT INTO eventmiddle SELECT * FROM eventsmall WHERE not EXISTS( select * from eventmiddle WHERE eventmiddle.id=eventsmall.id)
;;
DELIMITER ;
