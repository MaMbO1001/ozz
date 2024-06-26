-- up
create table post
(
    id          int       not null default gen_random_int(),
    user_id     int       not null,
    text        text      not null,
    created_at  timestamp not null default now(),
    ban_comment boolean   not null default false,
)

create table comment
(
    id                int       not null default gen_random_int(),
    answer_comment_id int null,
    user_id           int       not null,
    post_id           int       not null,
    text              text      not null,
    created_at        timestamp not null default now(),
)
-- down
drop table comment;
drop table post;