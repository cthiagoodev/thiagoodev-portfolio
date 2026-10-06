DROP POLICY IF EXISTS "Allow public read access" ON projects_skills;
ALTER TABLE projects_skills DISABLE ROW LEVEL SECURITY;
DROP TABLE IF EXISTS projects_skills;
