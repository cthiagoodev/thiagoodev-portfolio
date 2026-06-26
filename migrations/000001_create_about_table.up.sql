CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS about
(
    uuid       UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    name       VARCHAR(255) NOT NULL,
    bio        TEXT         NOT NULL,
    photo      VARCHAR(500) NOT NULL,
    curriculum VARCHAR(500) NOT NULL,
    linkedin   VARCHAR(255) NOT NULL,
    github     VARCHAR(255) NOT NULL,
    city       VARCHAR(100) NOT NULL,
    state      VARCHAR(2)   NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);