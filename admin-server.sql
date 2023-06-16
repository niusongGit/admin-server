
CREATE TABLE `admin` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `admin_name` varchar(40) NOT NULL DEFAULT '',
    `password` varchar(255) NOT NULL DEFAULT '',
    `last_login_time` timestamp NOT NULL,
    `status` tinyint NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='管理员表';

CREATE TABLE `banner` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `sport_type_id` bigint NOT NULL DEFAULT 0,
    `name` varchar(40) NOT NULL DEFAULT '',
    `banner_src` varchar(255) NOT NULL DEFAULT '' COMMENT 'banner图片地址',
    `status` tinyint NOT NULL DEFAULT 1,
    PRIMARY KEY (`id`),
    INDEX idx_sport_type_id (`sport_type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='banner表';

CREATE TABLE `announcement` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `sport_type_id` bigint NOT NULL DEFAULT 0,
    `title` varchar(40) NOT NULL DEFAULT '',
    `content` text NOT NULL,
    `status` tinyint NOT NULL DEFAULT 1,
    PRIMARY KEY (`id`),
    INDEX idx_sport_type_id (`sport_type_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='公告表';

CREATE TABLE `blogroll` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `category` varchar(40) NOT NULL DEFAULT '',
    `name` varchar(40) NOT NULL DEFAULT '' COMMENT '友情链接名字',
    `icon` varchar(255) NOT NULL DEFAULT '' COMMENT 'icon图片地址',
    `link` varchar(255) NOT NULL DEFAULT '' COMMENT '链接地址',
    `status` tinyint NOT NULL DEFAULT 1,
    PRIMARY KEY (`id`),
    INDEX idx_sport_type_id (`sport_type_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='友情链接表';


CREATE TABLE `sport_type` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(40) NOT NULL DEFAULT '' COMMENT '比赛名字',
    `status` tinyint NOT NULL DEFAULT 1,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='运动类型表';

CREATE TABLE `competition_type` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `sport_type_id` bigint NOT NULL DEFAULT 0,
    `name` varchar(40) NOT NULL DEFAULT '' COMMENT '比赛名字',
    `icon` varchar(255) NOT NULL DEFAULT '' COMMENT 'icon图片地址',
    `status` tinyint NOT NULL DEFAULT 1,
    PRIMARY KEY (`id`),
    INDEX idx_sport_type_id (`sport_type_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='比赛类型表';