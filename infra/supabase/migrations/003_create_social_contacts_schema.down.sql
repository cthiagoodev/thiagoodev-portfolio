DROP POLICY IF EXISTS "Allow public read access" ON social_contacts;
ALTER TABLE social_contacts DISABLE ROW LEVEL SECURITY;
DROP TABLE IF EXISTS social_contacts;
