drop database if exists gin_gorm_oj;

drop table if exists problem_basic;
drop table if exists problem_category;
drop table if exists submit_basic;
drop table if exists test_case;
drop table if exists user_basic;




create database gin_gorm_oj char set = utf8 collate = utf8_general_ci;

use gin_gorm_oj;





create table problem_basic (
                               id int(11) unsigned  not null primary key auto_increment ,
                               identity varchar(36) not null ,
                               title varchar(255),
                               content text,
                               max_runtime int(11),
                               max_mem int(11),
                               pass_num int(11),
                               submit_num int(11),
                               created_at datetime,
                               updated_at datetime,
                               deleted_at datetime
#     foreign key(id) references problem_categories(problem_id),
#     foreign key (identity) references test_case(problem_identity)
);

create table problem_categories (
                                    id int(11) unsigned not null primary key auto_increment ,
                                    problem_id int(11)unsigned,
                                    category_id int(11)unsigned,
                                    created_at datetime,
                                    updated_at datetime,
                                    deleted_at datetime,
                                    foreign key (problem_id) references problem_basic(id)
);









create table user_basic (
                               id int(11) unsigned not null primary key auto_increment,
                               identity varchar(36) not null ,
                               name varchar(100),
                               password varchar(32),
                               phone varchar(20),
                               mail varchar(100),
                               pass_num int(11),
                               submit_num int(11),
                               is_admin tinyint(1) default 0,
                               created_at datetime,
                               updated_at datetime,
                               deleted_at datetime
);





create table submit_basic (
                               identity varchar(36) not null ,
                               problem_identity varchar(36),
                               user_identity varchar(36),
                               path varchar(255),
                               status tinyint(1) default -1
#                                foreign key(user_identity) references user_basic(identity),
#                                foreign key (identity) references test_case(problem_identity)
);



create table test_case (
                           id int(11) unsigned not null primary key auto_increment,
                           identity varchar(255) not null ,
                           problem_identity varchar(255),
                           input text,
                           output text,
                           created_at datetime,
                           updated_at datetime,
                           deleted_at datetime
);


## 添加外键

alter table problem_category
    add constraint 关联问题分类表
        foreign key (problem_id) references problem_basic (id);

alter table problem_basic
    add constraint 关联测试用例表
        foreign key (identity) references test_case (identity);




alter table submit_basic
    add constraint 关联问题基础表
        foreign key (problem_identity) references user_basic (identity);

alter table submit_basic
    add constraint 关联用户基础表
        foreign key (user_identity) references user_basic (identity);


