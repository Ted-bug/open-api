/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : openapi

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 29/11/2024 22:46:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for op_ak_sk
-- ----------------------------
DROP TABLE IF EXISTS `op_ak_sk`;
CREATE TABLE `op_ak_sk`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `ak` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sk` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `show` char(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'no' COMMENT '是否已显示过',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态',
  `create_time` datetime(0) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户aksk表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of op_ak_sk
-- ----------------------------
INSERT INTO `op_ak_sk` VALUES (1, 'aaaaccccddddaaaaccccddddaaaacccc', 'abcd', 'edfg', 'no', 1, '2024-04-13 23:19:23');

-- ----------------------------
-- Table structure for op_short_url
-- ----------------------------
DROP TABLE IF EXISTS `op_short_url`;
CREATE TABLE `op_short_url`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `lurl` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `surl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` tinyint(1) UNSIGNED NULL DEFAULT 1,
  `create_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `'short-ind'`(`surl`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 42 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '短链接生成表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of op_short_url
-- ----------------------------
INSERT INTO `op_short_url` VALUES (28, 'http://dypogtpcff.si/bcocy', '71b0721db2cbee4c1ca9dff1adad5bc7', '2yLXUKAPym3', 1, '2024-05-01 00:53:24');

-- ----------------------------
-- Table structure for op_unique_num
-- ----------------------------
DROP TABLE IF EXISTS `op_unique_num`;
CREATE TABLE `op_unique_num`  (
  `id` bigint(32) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '不同渠道的标识',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `type`(`type`) USING BTREE COMMENT '用于构建唯一值'
) ENGINE = InnoDB AUTO_INCREMENT = 4020 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '全局唯一递增id' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of op_unique_num
-- ----------------------------
INSERT INTO `op_unique_num` VALUES (4020, 'short');

-- ----------------------------
-- Table structure for op_user
-- ----------------------------
DROP TABLE IF EXISTS `op_user`;
CREATE TABLE `op_user`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uid` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of op_user
-- ----------------------------
INSERT INTO `op_user` VALUES (3, 'dfasfd', 'a');

SET FOREIGN_KEY_CHECKS = 1;
