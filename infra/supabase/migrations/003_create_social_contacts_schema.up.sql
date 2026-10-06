CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS social_contacts
(
    uuid       UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    label      VARCHAR(255) NOT NULL UNIQUE,
    url        TEXT         NOT NULL,
    image_path TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

ALTER TABLE social_contacts ENABLE ROW LEVEL SECURITY;
CREATE POLICY "Allow public read access" ON social_contacts FOR SELECT TO anon USING (true);
