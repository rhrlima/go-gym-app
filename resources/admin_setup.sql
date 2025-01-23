-- Setup script for admin and app user
-- Use admin credentials to create an app user with database creation permissions
DO $$
BEGIN
    -- Create the app user if it doesn't already exist
    IF NOT EXISTS (
        SELECT FROM pg_roles WHERE rolname = 'gymapp'
    ) THEN
        CREATE ROLE gymapp WITH LOGIN PASSWORD 'securepassword';
        GRANT CREATEDB TO gymapp;
    END IF;
END $$;

-- Grant the necessary permissions for the app user
GRANT CONNECT ON DATABASE postgres TO gymapp;
GRANT ALL PRIVILEGES ON SCHEMA public TO gymapp;
