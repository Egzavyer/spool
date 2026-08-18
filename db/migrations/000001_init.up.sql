DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'job_status') THEN
        CREATE TYPE job_status AS ENUM ('pending', 'completed', 'failed', 'running');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS jobs (
    job_id UUID PRIMARY KEY,
    request_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    status job_status DEFAULt 'pending'
);