create table app
(
    id           int unsigned auto_increment
        primary key,
    app_id       varchar(200)                              not null,
    app_name     varchar(200)                              not null,
    deleted      tinyint(1)   default 0                    not null,
    created_time timestamp(3) default CURRENT_TIMESTAMP(3) not null,
    updated_time timestamp(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3),
    constraint app_app_id_uindex
        unique (app_id)
);

create table app_env
(
    id           int unsigned auto_increment
        primary key,
    name         varchar(200)                              not null,
    app_id       varchar(200)                              not null,
    deleted      tinyint(1)   default 0                    not null,
    created_time timestamp(3) default CURRENT_TIMESTAMP(3) not null,
    updated_time timestamp(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3),
    constraint app_env_name_app_id_uindex
        unique (name, app_id)
);

create table cluster
(
    id           int unsigned auto_increment
        primary key,
    name         varchar(200)                              not null,
    app_id       varchar(200)                              not null,
    deleted      tinyint(1)   default 0                    not null,
    created_time timestamp(3) default CURRENT_TIMESTAMP(3) not null,
    updated_time timestamp(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3)
);

create index cluster_app_id_index
    on cluster (app_id);

