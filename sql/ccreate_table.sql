drop table if exists test_case;

create table test_case (
                           id int(11) not null primary key ,
                           identity varchar(36) not null ,
                           problem_identity varchar(36),
                           input text,
                           output text,
                           created_at datetime,
                           updated_at datetime,
                           deleted_at datetime);