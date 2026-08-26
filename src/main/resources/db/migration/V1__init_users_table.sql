create table users (

    id              UUID            primary key     default gen_random_uuid(),
    name            varchar(128)    not null,
    surname         varchar(256)    not null,
    email           varchar(512)    not null        unique ,
    password_hash   varchar(512)    not null,

    created_at      timestamp with time zone        default current_timestamp not null,
    updated_at      timestamp with time zone        default current_timestamp not null


)