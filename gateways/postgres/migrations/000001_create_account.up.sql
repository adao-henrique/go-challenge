BEGIN;

create table account
(
    id          uuid            not null     constraint  person_pk   primary key,
    name        varchar(255)    not null,    
    cpf         varchar(14)     not null,
    secret      varchar(100)    not null,  
    balance     float8          not null,
    created_at  date            not null
);

COMMIT;