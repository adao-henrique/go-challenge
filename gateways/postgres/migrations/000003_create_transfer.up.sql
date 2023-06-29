BEGIN;

create table transfer
(
    id                      uuid    not null    PRIMARY KEY,
	account_origin_id       uuid    not null    REFERENCES account,
	account_destination_id  uuid    not null    REFERENCES account,
	amount                  float8     not null,
	created_at              date    not null
);

COMMIT;