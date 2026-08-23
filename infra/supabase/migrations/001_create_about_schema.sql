CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS about
(
    uuid        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(255) NOT NULL,
    phone       VARCHAR(14),
    description VARCHAR(255),
    text        TEXT
);