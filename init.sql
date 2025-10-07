-- Создаем базу данных если не существует
SELECT 'CREATE DATABASE db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'db')\gexec

SELECT 'CREATE DATABASE db_test'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'db_test')\gexec