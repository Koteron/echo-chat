--liquibase formatted sql

--changeset author:Koteron failOnError:true

create extension if not exists "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE users (
   id UUID PRIMARY KEY,
   display_name TEXT NOT NULL,
   bio TEXT,
   is_public BOOLEAN,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_user_name_trgm ON users USING gin (display_name gin_trgm_ops);


-----

-- rollback drop table users;
