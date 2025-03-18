START TRANSACTION;

CREATE SCHEMA IF NOT EXISTS auth;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE auth.users (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT pk_users PRIMARY KEY (id)
);


CREATE UNIQUE INDEX ix_users_id ON auth.users (id);
CREATE UNIQUE INDEX ix_users_username ON auth.users (username);

COMMIT;
