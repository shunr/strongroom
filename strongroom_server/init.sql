DROP TABLE if exists accounts
CASCADE;
CREATE TABLE
if not exists accounts
(
	id serial PRIMARY KEY,
	username VARCHAR
(32) UNIQUE NOT NULL,
    display_name TEXT,
	auth_verifier BYTEA UNIQUE NOT NULL,
    auth_salt BYTEA NOT NULL,
    muk_salt BYTEA NOT NULL
);

DROP TABLE if exists sessions
CASCADE;
CREATE TABLE
if not exists sessions
(
	id UUID PRIMARY KEY,
	account_id serial REFERENCES accounts
(id),
    secret BYTEA NOT NULL
);