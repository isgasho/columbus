-- Auto-generated at Thu, 06 Sep 2018 18:07:47 CST
-- Please do not change the name attributes

-- name: up

CREATE TABLE IF NOT EXISTS `account` (
  `uid` int unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(128) NOT NULL DEFAULT '',
  `username` varchar(20) NOT NULL COMMENT '用户名',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像(如果为空，则使用http://www.gravatar.com)',
  `city` varchar(10) NOT NULL DEFAULT '' COMMENT '居住地',
  `website` varchar(63) NOT NULL DEFAULT '' COMMENT '个人主页，博客',
  `monlog` varchar(140) NOT NULL DEFAULT '' COMMENT '个人状态，签名，独白',
  `introduce` varchar(2022) NOT NULL COMMENT '个人简介',
  `is_root` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否超级用户，不受权限控制：1-是',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  UNIQUE KEY (`username`),
  UNIQUE KEY (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户信息表';

CREATE TABLE IF NOT EXISTS `authentication` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `open_id` varchar(127) NOT NULL DEFAULT '' COMMENT '用户的标识，对当前公众号/小程序唯一',
  `union_id` varchar(127) NOT NULL DEFAULT '' COMMENT '用户的标识，对开放者唯一',
  `nickname` varchar(127) NOT NULL DEFAULT '' COMMENT '用户的昵称',
  `session_key` varchar(127) NOT NULL DEFAULT '' COMMENT '小程序返回的 session_key',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户微信头像',
  `open_info` varchar(1024) NOT NULL DEFAULT '' COMMENT '用户微信的其他信息，json格式',
  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户UID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `open_id` (`open_id`),
  KEY `uid` (`uid`),
  KEY `updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='微信用户绑定表';

-- name: down
-- DROP TABLE IF EXISTS account;
-- DROP TABLE IF EXISTS authentication;
