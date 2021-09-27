CREATE DATABASE IF NOT EXISTS `easyblog` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

/* 使用了ent中的自动迁移功能，一般不需要手动建表 */
USE `easyblog`;

DROP TABLE IF EXISTS `blog_index`;

CREATE TABLE `blog_index` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `content_id` INT,
    `title` VARCHAR(50) DEFAULT '' COMMENT '标题',
    `desc` VARCHAR(255) DEFAULT '' COMMENT '简介',
    `view` INT COMMENT '浏览数量',
    `attr` INT COMMENT '属性',
    `created_at` DATETIME COMMENT '创建时间',
    `updated_at` DATETIME COMMENT '修改时间',
    `state` TINYINT DEFAULT '1' COMMENT '是否可用',
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`content_id`)
) ENGINE=InnoDB, DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `blog_content`;

CREATE TABLE `blog_content` (
    `id` INT NOT NULL AUTO_INCREMENT, 
    `index_id` INT,
    `content` TEXT COMMENT '内容',
    `created_at` DATETIME COMMENT '创建时间',
    `updated_at` DATETIME COMMENT '修改时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`index_id`)
) ENGINE=InnoDB, DEFAULT CHARSET=utf8mb4;
