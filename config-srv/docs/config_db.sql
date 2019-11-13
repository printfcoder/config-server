create table app
(
	id int unsigned auto_increment
		primary key,
	app_id varchar(200) not null,
	app_name varchar(200) not null,
	deleted tinyint(1) default 0 not null,
	created_time timestamp(3) default CURRENT_TIMESTAMP(3) not null,
	updated_time timestamp(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3),
	constraint app_app_id_uindex
		unique (app_id)
);

create table cluster
(
	id int unsigned auto_increment
		primary key,
	name varchar(200) not null,
	app_id varchar(200) not null,
	deleted tinyint(1) default 0 not null,
	created_time timestamp(3) default CURRENT_TIMESTAMP(3) not null,
	updated_time timestamp(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3)
);

create index cluster_app_id_index
	on cluster (app_id);

create table env
(
	id int unsigned auto_increment
		primary key,
	name varchar(200) not null,
	app_id varchar(200) not null,
	deleted tinyint(1) default 0 not null,
	created_time timestamp(3) default CURRENT_TIMESTAMP(3) not null,
	updated_time timestamp(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3),
	constraint app_env_name_app_id_uindex
		unique (name, app_id)
);

create table instance
(
	id int unsigned auto_increment
		primary key,
	app_id varchar(200) not null,
	cluster_name varchar(200) not null,
	ip varchar(150) not null,
	created_time timestamp(3) default CURRENT_TIMESTAMP(3) not null,
	updated_time timestamp(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3),
	constraint instance_app_id_cluster_name_ip_uindex
		unique (app_id, cluster_name, ip)
);

create table item
(
	id int unsigned auto_increment
		primary key,
	namespace_id int unsigned not null,
	`key` varchar(128) not null,
	value text null,
	deleted tinyint(1) not null,
	created_time timestamp(3) default CURRENT_TIMESTAMP(3) not null,
	updated_time timestamp(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3)
);

create index item_updated_time_index
	on item (updated_time);

create table namespace
(
	id int unsigned auto_increment
		primary key,
	app_id varchar(200) not null,
	name varchar(200) not null,
	cluster_name varchar(200) not null,
	deleted tinyint(1) not null,
	created_time timestamp(3) default CURRENT_TIMESTAMP(3) not null,
	updated_time timestamp(3) default CURRENT_TIMESTAMP(3) not null on update CURRENT_TIMESTAMP(3),
	constraint namespace_app_id_cluster_name_name_uindex
		unique (app_id, cluster_name, name)
);

create index namespace_updated_time_index
	on namespace (updated_time);

