alter table if EXISTS "accounts" drop CONSTRAINT if EXISTS "owner_currency_key";

alter table if EXISTS "accounts" drop CONSTRAINT if EXISTS "accounts_wner_fkey";

alter table if EXISTS "users" drop ;
