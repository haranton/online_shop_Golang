-- init.sql
DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'db') THEN
        CREATE DATABASE db;
    END IF;
END
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'db_test') THEN
        CREATE DATABASE db_test;
    END IF;
END
$$;