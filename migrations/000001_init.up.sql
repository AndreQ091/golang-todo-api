create table users (
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    passwordHash varchar(255) not null
);
create table todo_lists (
    id serial primary key,
    title varchar(255) not null,
    description varchar(255)
);
create table user_lists (
    id serial primary key,
    user_id int references users (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);
create table todo_items (
    id serial primary key,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);
create table list_items (
    id serial primary key,
    item_id int references todo_items (id) on delete cascade not null,
    list_id int references todo_lists (id) on delete cascade not null
);