CREATE TABLE `user`
(
    `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `passWord` varchar(50)     NOT NULL DEFAULT '' COMMENT '用户密码,bcrypt加密',
    `user_Nick`     varchar(100)    NOT NULL DEFAULT 'momo' COMMENT '用户昵称',
    `user_Face`     varchar(255)    NOT NULL DEFAULT '' COMMENT '用户头像地址',
    `User_Sex`      tinyint(1)      NOT NULL DEFAULT 0 COMMENT '用户性别:0男,1女,未知',
    `user_Email`    varchar(255)    NOT NULL DEFAULT ''COMMENT '用户邮箱',
    `user_Phone`    varchar(11)     NOT NULL DEFAULT '' COMMENT '手机号',
    `create_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_time` datetime      ,
    PRIMARY KEY (`user_id`),
    INDEX `userPhone` (`user_Phone`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';


CREATE TABLE `user_auth`
(
    `id` bigint NOT NULL AUTO_INCREMENT,
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` datetime ,
    `user_id` bigint NOT NULL DEFAULT '0',
    `provider_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台唯一id',
    `provider` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台类型',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT ='用户鉴权表';

-- goctl model mysql ddl -src="user.sql" -dir="./" --cache=true --prefix="user:"