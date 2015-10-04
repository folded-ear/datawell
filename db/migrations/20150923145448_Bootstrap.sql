
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create sequence id_seq start 1 increment by 1;
create table users (
    id bigint not null default nextval('id_seq'),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name varchar(50) not null,
    username varchar(50) not null,
    passhash varchar(50) not null,
    constraint pk_users primary key (id),
    constraint uk_users_username unique (username)
);
create table events (
    id bigint not null default nextval('id_seq'),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    user_id bigint not null,
    timestamp timestamptz not null,
    notes text not null default '',
    constraint pk_events primary key (id),
    constraint fk_events_user_id foreign key (user_id) references users (id) on delete cascade
);
create table tags (
    id bigint not null default nextval('id_seq'),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    user_id bigint not null,
    tag varchar not null,
    constraint pk_tags primary key (id),
    constraint fk_tags_user_id foreign key (user_id) references users (id) on delete cascade,
    constraint uk_tags_user_tag unique (user_id, tag)
);
create table event_tags (
    event_id bigint not null,
    tag_id bigint not null,
    number float8 not null default 1,
    constraint pk_event_tags primary key (event_id, tag_id),
    constraint fk_event_tags_event_id foreign key (event_id) references events (id),
    constraint fk_event_tags_tag_id foreign key (tag_id) references tags (id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table event_tags;
drop table tags;
drop table events;
drop table users;
drop sequence id_seq;
