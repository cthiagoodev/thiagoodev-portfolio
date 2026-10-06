DROP POLICY IF EXISTS "Allow public read access" ON work_experience;
ALTER TABLE work_experience DISABLE ROW LEVEL SECURITY;
DROP TABLE IF EXISTS work_experience;
