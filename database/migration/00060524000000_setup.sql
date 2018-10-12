-- Auto-generated at Thu, 06 Sep 2018 18:07:47 CST
-- Please do not change the name attributes

-- name: up

CREATE TABLE IF NOT EXISTS `account` (
  `uid` int unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(128) NOT NULL DEFAULT '',
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `avatar` varchar(128) NOT NULL DEFAULT '' COMMENT '头像(如果为空，则使用http://www.gravatar.com)',
  `city` varchar(10) NOT NULL DEFAULT '' COMMENT '居住地',
  `introduce` varchar(2022) NOT NULL COMMENT '个人简介',
  `open_id` varchar(127) NOT NULL DEFAULT '' COMMENT '用户的标识，对当前公众号/小程序唯一',
  `union_id` varchar(127) NOT NULL DEFAULT '' COMMENT '用户的标识，对开放者唯一',
  `is_root` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否超级用户，不受权限控制：1-是',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  UNIQUE KEY (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户信息表';

-- name: down
-- DROP TABLE IF EXISTS account;
