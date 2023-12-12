create database digitalwallet2;
create user digitalwallet2 with password 'digitalwallet2';
grant all privileges on database digitalwallet2 to digitalwallet2;;

create table accounts(
    id uuid not null primary key,
    balance bigint,
    created_at timestamp
);