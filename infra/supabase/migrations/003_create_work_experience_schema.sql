CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS work_experience
(
    uuid        UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    role        VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    company     VARCHAR(255) NOT NULL,
    start_date  DATE         NOT NULL,
    end_date    DATE,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT now()
);