
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create sequence id_seq;
create table users (
    id bigint not null default nextval('id_seq'),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    name varchar(50),
    username varchar(50),
    passhash varchar(50),
    constraint pk_users primary key (id),
    constraint uk_users_username unique (username)
);
create table events (
    id bigint not null default nextval('id_seq'),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    user_id bigint not null,

    notes text,
    constraint pk_events primary key (id),
    constraint fk_event_user_id foreign key (user_id) references users (id) on delete cascade
);
create table tags (
    id bigint not null default nextval('id_seq'),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    tag varchar,
    constraint pk_tags primary key (id),
    constraint uk_tags_tag unique (tag)
);
create table event_tags (
    event_id bigint not null,
    tag_id bigint not null,
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
