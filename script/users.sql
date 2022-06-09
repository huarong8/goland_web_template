create table users (
    id int not null primary key auto_increment,
    user_id bigint not null default 0,
    name varchar(128) not null default "",
    email varchar(128) not null default "",
    passwd varchar(64) not null default "",
    `stat` tinyint not null default 1,
    created_at timestamp not null default current_timestamp,
    update_at timestamp not null default current_timestamp on update current_timestamp,
    unique `uniq_user_id`(user_id),
    unique `uniq_email`(email)
) ENGINE=InnoDB default charset = utf8mb4;
