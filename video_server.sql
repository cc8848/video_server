create table user (
  id          int unsigned primary key      auto_increment,
  username    varchar(64) not null unique,
  password    varchar(64) not null,
  create_time timestamp   not null          default current_timestamp,
  update_time timestamp   not null          default current_timestamp
  on update current_timestamp
)
  engine = InnoDB
  default charset = utf8;

create table video_info (
  id            varchar(64) primary key,
  author_id     int unsigned not null,
  name          varchar(128) not null,
  display_ctime varchar(128) not null,
  create_time   timestamp    not null  default current_timestamp,
  update_time   timestamp    not null  default current_timestamp
  on update current_timestamp
)
  engine = InnoDB
  default charset = utf8;

create table comment (
  id        varchar(64) primary key,
  video_id  varchar(64) not null,
  author_id int(10)     not null,
  content   text        not null,
  time      datetime    not null default current_timestamp
)
  engine = InnoDB
  default charset = utf8;
