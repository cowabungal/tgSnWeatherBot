CREATE TABLE admins
(
    id       serial       not null unique,
    username varchar(255) not null,
    user_id varchar (30) not null unique
);