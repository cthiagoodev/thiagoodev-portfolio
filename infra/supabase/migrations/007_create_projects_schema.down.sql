DROP POLICY IF EXISTS "Allow public read access" ON projects;
ALTER TABLE projects DISABLE ROW LEVEL SECURITY;
DROP TABLE IF EXISTS projects;
