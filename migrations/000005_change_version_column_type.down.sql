ALTER TABLE movies ALTER COLUMN version DROP DEFAULT;
ALTER TABLE movies ALTER COLUMN version TYPE integer USING 1;
ALTER TABLE movies ALTER COLUMN version SET DEFAULT 1;