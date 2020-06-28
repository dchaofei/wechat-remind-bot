create database `remind` character set utf8mb4 collate utf8mb4_general_ci;

create table `room` (
`id` int primary key auto_increment,
`wx_room_id` varchar(50) not null comment '微信群聊id',
`admin_wx_id` varchar(50) not null comment '管理员微信id',
`status` tinyint not null default 0 comment '是否开启签到 0关闭，1开启',
`created_on` timestamp not null default current_timestamp,
`modified_on` timestamp not null default current_timestamp on update current_timestamp,
unique key (`wx_room_id`)
) ENGINE=innoDB DEFAULT CHARSET=utf8mb4 COMMENT='微信群表';

create table `checkin` (
`id` int primary key auto_increment,
`room_id` int not null comment '微信群聊主键id',
`wx_id` varchar(50) not null comment '微信id',
`date` date not null comment '签到日期',
`created_on` timestamp not null default current_timestamp,
unique (`room_id`,`wx_id`,`date`)
) ENGINE=innoDB DEFAULT CHARSET=utf8mb4 COMMENT='签到表';

create table `not_remind` (
`id` int primary key auto_increment,
`room_id` int not null comment '微信群聊主键id',
`wx_id` varchar(50) not null comment '微信id',
`created_on` timestamp not null default current_timestamp,
unique key (`room_id`, `wx_id`)
) ENGINE=innoDB DEFAULT CHARSET=utf8mb4 COMMENT='不提醒表';
