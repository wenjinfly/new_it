/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.0.169
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 192.168.0.169:33069
 Source Schema         : new_it

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 18/04/2023 00:47:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_dict
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict`;
CREATE TABLE `sys_dict`  (
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `type_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `cn_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `fixed` tinyint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`code`) USING BTREE,
  INDEX `type_code`(`type_code`) USING BTREE,
  CONSTRAINT `sys_dict_ibfk_1` FOREIGN KEY (`type_code`) REFERENCES `sys_dict_type` (`code`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict
-- ----------------------------
INSERT INTO `sys_dict` VALUES ('101', '100', NULL, '软件策划', 1);
INSERT INTO `sys_dict` VALUES ('102', '100', NULL, '软件需求', 1);
INSERT INTO `sys_dict` VALUES ('103', '100', NULL, '概要设计', 1);
INSERT INTO `sys_dict` VALUES ('104', '100', NULL, '详细设计', 1);
INSERT INTO `sys_dict` VALUES ('105', '100', NULL, '单元测计划', 1);
INSERT INTO `sys_dict` VALUES ('106', '100', NULL, '编码和单测', 1);
INSERT INTO `sys_dict` VALUES ('107', '100', NULL, '系统测试计划', 1);
INSERT INTO `sys_dict` VALUES ('108', '100', NULL, '系统测试', 1);
INSERT INTO `sys_dict` VALUES ('109', '100', NULL, '用户测试', 1);
INSERT INTO `sys_dict` VALUES ('110', '100', NULL, '软件发行', 1);
INSERT INTO `sys_dict` VALUES ('111', '100', NULL, '部署运维', 1);

SET FOREIGN_KEY_CHECKS = 1;
