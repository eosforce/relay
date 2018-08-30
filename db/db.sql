create table if not exists public.accounts
(
	name varchar(32) not null,
	chain varchar(16) not null,
	create_time timestamp default now(),
	constraint accounts_pk
		primary key (chain, name)
);

create table if not exists public.account_tokens
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

create table public.account_permissions
(
  name       varchar(32) not null,
  chain      varchar(16) not null,
  permission varchar(16) not null,
  pubkey     varchar(64) not null,
  constraint account_keys_pk
  primary key (name, chain)
);

create index account_permissions_permission_pubkey_index
  on public.account_permissions (permission, pubkey);

create table if not exists public.chain_data
(
  name varchar(16) not null
    constraint chain_data_pkey
    primary key,
  note varchar(256) default '' :: character varying,
  typ  varchar(16) not null
);

create unique index if not exists chain_data_name_uindex
  on public.chain_data (name);
