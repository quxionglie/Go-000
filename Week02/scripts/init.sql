create table user
(
    id           bigint                                 not null comment '主键'
        primary key,
    username     varchar(64)  default ''                not null comment '用户名',
    pwd          varchar(128) default ''                not null comment '密码',
    gmt_create   datetime     default CURRENT_TIMESTAMP not null comment '添加时间',
    gmt_modified datetime     default CURRENT_TIMESTAMP not null comment '更新时间',
    constraint user_username_uindex
        unique (username)
)
    comment '用户表';