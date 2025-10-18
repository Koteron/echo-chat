--liquibase formatted sql

--changeset author:Koteron failOnError:true

create extension if not exists "uuid-ossp";

CREATE TABLE users (
   id UUID PRIMARY KEY,
   display_name TEXT NOT NULL,
   bio TEXT,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-----

-- rollback drop table users;
