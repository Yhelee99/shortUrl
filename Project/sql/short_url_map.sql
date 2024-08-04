CREATE TABLE `short_url_map`(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_by` varchar(64) NOT NULL DEFAULT '' COMMENT '创建者',
    `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否已删除：0正常1已删除',

    `lurl` varchar(160) DEFAULT NULL COMMENT '长链接',
    `surl` varchar(11) DEFAULT NULL COMMENT '短链接',
    PRIMARY KEY (`id`),
    index (`is_del`),
    UNIQUE(`lurl`),
    UNIQUE(`surl`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='长短链接映射表';