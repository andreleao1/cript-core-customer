    create table balance_reserves (
        id TEXT primary key,
        wallet_id TEXT not null,
        amount TEXT not null,
        status TEXT not null default 'pending',
        created_at timestamp not null default now(),
        updated_at timestamp not null default now(),
        constraint fk_balance_reserve_wallet_id foreign key (wallet_id) references wallets(id)
    );