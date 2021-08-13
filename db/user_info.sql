/*
 Navicat Premium Data Transfer

 Source Server         : malldb
 Source Server Type    : MariaDB
 Source Server Version : 100604
 Source Host           : localhost:3306
 Source Schema         : malldb

 Target Server Type    : MariaDB
 Target Server Version : 100604
 File Encoding         : 65001

 Date: 13/08/2021 11:14:04
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户序号',
    `first_name`  varchar(255)      DEFAULT '' COMMENT '用户名字',
    `last_name`   varchar(255)      DEFAULT '' COMMENT '用户姓氏',
    `phone`       varchar(255)      DEFAULT '' COMMENT '手机号',
    `age`         tinyint(5) DEFAULT 0 COMMENT '年龄',
    `sex`         tinyint(5) DEFAULT 1 COMMENT '性别',
    `avatar`      varchar(255)      DEFAULT '' COMMENT '头像',
    `password`    varchar(255)      DEFAULT '' COMMENT '登录密码',
    `create_time` datetime          DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE current_timestamp () COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Records of user_info
-- ----------------------------
BEGIN;
COMMIT;

SET
FOREIGN_KEY_CHECKS = 1;
