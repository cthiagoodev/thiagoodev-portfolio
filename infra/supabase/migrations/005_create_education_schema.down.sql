DROP POLICY IF EXISTS "Allow public read access" ON education;
ALTER TABLE education DISABLE ROW LEVEL SECURITY;
DROP TABLE IF EXISTS education;
