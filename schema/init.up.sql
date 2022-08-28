CREATE TABLE admins
(
    id serial not null unique,
    name varchar(255) not null,
    login varchar(255) not null unique,
    password varchar(255) not null
);

CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    login varchar(255) not null unique,
    password varchar(255) not null,
    root int references admins (id) on delete cascade
);