CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS projects_skills
(
    project_id UUID NOT NULL,
    skill_id   UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),

    PRIMARY KEY (project_id, skill_id),

    CONSTRAINT fk_project
        FOREIGN KEY (project_id)
        REFERENCES projects (uuid)
        ON DELETE CASCADE,

    CONSTRAINT fk_skill
        FOREIGN KEY (skill_id)
        REFERENCES skills (uuid)
        ON DELETE CASCADE
);