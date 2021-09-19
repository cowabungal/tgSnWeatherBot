CREATE TABLE users
(
    id       serial       not null unique,
    username varchar(255) not null,
    user_id varchar (30) not null unique
);

CREATE TABLE names
(
    id          serial                                      not null unique,
    name varchar (30),
    user_id     varchar (30) references users (user_id) on delete cascade not null
);