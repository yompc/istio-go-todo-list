/*
 Navicat Premium Data Transfer

 Source Server         : okteto
 Source Server Type    : MySQL
 Source Server Version : 100318
 Source Host           : localhost:3307
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 100318
 File Encoding         : 65001

 Date: 02/07/2020 14:25:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_permission
-- ----------------------------
DROP TABLE IF EXISTS `tb_permission`;
CREATE TABLE `tb_permission`  (
  `id` bigint NOT NULL,
  `created` date NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `epermission_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `parent_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `permission_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `updated` date NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_permission
-- ----------------------------
INSERT INTO `tb_permission` VALUES (1, NULL, NULL, '控制面板', NULL, 'dashboard', NULL, '/dashboard');
INSERT INTO `tb_permission` VALUES (2, NULL, NULL, '测试', NULL, 'test', NULL, '/test');

-- ----------------------------
-- Table structure for tb_role
-- ----------------------------
DROP TABLE IF EXISTS `tb_role`;
CREATE TABLE `tb_role`  (
  `id` bigint NOT NULL,
  `created` date NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `erole_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `parent_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `updated` date NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_role
-- ----------------------------
INSERT INTO `tb_role` VALUES (1, NULL, NULL, '管理员', NULL, 'admin', NULL);
INSERT INTO `tb_role` VALUES (2, NULL, NULL, '测试', NULL, 'test', NULL);

-- ----------------------------
-- Table structure for tb_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `tb_role_permission`;
CREATE TABLE `tb_role_permission`  (
  `id` bigint NOT NULL,
  `permission_id` bigint NULL DEFAULT NULL,
  `role_id` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_role_permission
-- ----------------------------
INSERT INTO `tb_role_permission` VALUES (1, 1, 1);
INSERT INTO `tb_role_permission` VALUES (2, 2, 1);
INSERT INTO `tb_role_permission` VALUES (3, 2, 2);

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`  (
  `id` bigint NOT NULL,
  `created` date NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `refresh_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `updated` date NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES (1, NULL, NULL, '$2a$10$byHOajCrGuITq0PPCxq2feKMyFQbXAvRuraEgP8kvFn3xWuS3UL.S', NULL, NULL, NULL, 'admin');
INSERT INTO `tb_user` VALUES (2, NULL, NULL, '$2a$10$byHOajCrGuITq0PPCxq2feKMyFQbXAvRuraEgP8kvFn3xWuS3UL.S', NULL, NULL, NULL, 'test');

-- ----------------------------
-- Table structure for tb_user_role
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_role`;
CREATE TABLE `tb_user_role`  (
  `id` bigint NOT NULL,
  `role_id` bigint NULL DEFAULT NULL,
  `user_id` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_user_role
-- ----------------------------
INSERT INTO `tb_user_role` VALUES (1, 1, 1);
INSERT INTO `tb_user_role` VALUES (2, 2, 2);

-- ----------------------------
-- Table structure for todo
-- ----------------------------
DROP TABLE IF EXISTS `todo`;
CREATE TABLE `todo`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of todo
-- ----------------------------
INSERT INTO `todo` VALUES (11, '你好世界', 0);
INSERT INTO `todo` VALUES (12, '小刘爱笔记', 0);
INSERT INTO `todo` VALUES (13, '小黄今天吃啥', 0);
INSERT INTO `todo` VALUES (14, '你好世界 小刘爱笔记', 0);

SET FOREIGN_KEY_CHECKS = 1;
