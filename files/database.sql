create database digitalwallet2;
create user digitalwallet2 with password 'digitalwallet2';
grant all privileges on database digitalwallet2 to digitalwallet2;;

create table accounts(
    id uuid not null primary key,
    balance bigint not null,
    created_at timestamp default now()
);

create table transactions (
    id uuid not null primary key,
    account_id uuid not null,
    external_id varchar(128),
    amount bigint not null,
    type varchar(64) not null,
    status varchar(64) not null,
    constraint transactions_accountid_fk foreign key(account_id) references accounts(id)
);
create unique index transactions_externalid_idx on transactions(external_id);

create table locked_accounts (
    account_id uuid not null primary key,
    process_number varchar(32)
);