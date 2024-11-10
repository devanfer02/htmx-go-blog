CREATE TABLE blogs (
    id serial primary key,
    title varchar(200) not null,
    slug varchar(200) unique,
    content text default '',
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);