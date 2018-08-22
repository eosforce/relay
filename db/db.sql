create table if not exists accounts
(
	name varchar(32) not null,
	chain varchar(16) not null,
	create_time timestamp default now(),
	constraint accounts_pk
		primary key (chain, name)
);

create table if not exists account_tokens
(
	name varchar(32) not null,
	chain varchar(16) not null,
	create_time timestamp default now(),
	update_time timestamp default now(),
	token_chain varchar(16) not null,
	symbol varchar(16) not null,
	amount bigint,
	constraint accounts_tokens_pk
		primary key (chain, name, token_chain, symbol)
);