create table users
(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique ,
    password_hash text not null
);

create table todo_lists
(
    id serial primary key,
    title varchar(255) not null,
    description text
);

create table users_lists
(
    id serial not null unique ,
    user_id int references users (id) on delete cascade not null ,
    list_id int references todo_lists (id) on delete cascade not null
);

create table todo_items
(
    id serial primary key,
    title varchar(255) not null,
    description text,
    done boolean not null default false
);

create table lists_items
(
    id serial not null unique ,
    list_id int references todo_lists (id) on delete cascade not null ,
    item_id int references todo_items (id) on delete cascade not null
);
