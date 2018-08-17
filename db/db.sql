create table if not exists accounts
(
	name varchar(32) not null,
	chain varchar(16) not null,
	create_time timestamp default now(),
	constraint accounts_pk
		primary key (chain, name)
);