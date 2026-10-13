CREATE TABLE IF NOT EXISTS projects (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    external_id INTEGER,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    url TEXT,
    repo_url TEXT,
    languages VARCHAR(255)[] DEFAULT '{}',
    pushed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
