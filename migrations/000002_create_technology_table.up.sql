CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS technology
(
    uuid       UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    name       VARCHAR(255) NOT NULL,
    about_uuid UUID         NOT NULL REFERENCES about (uuid) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);