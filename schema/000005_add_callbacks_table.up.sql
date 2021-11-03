CREATE TABLE callbacks
(
    id       serial       not null unique,
    user_id varchar (30) not null unique,
    callback_id varchar (30) not null unique,
    callback_data varchar (30)
);