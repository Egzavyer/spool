DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'job_status') THEN
        CREATE TYPE job_status AS ENUM ('pending', 'completed', 'failed', 'running');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS job_types (
    job_type_id SMALLSERIAL PRIMARY KEY,
    job_type_name VARCHAR(63)
);

CREATE TABLE IF NOT EXISTS jobs (
    job_id UUID PRIMARY KEY,
    requested_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    status job_status DEFAULt 'pending',
    job_type_id SMALLINT REFERENCES job_types(job_type_id) ON DELETE SET NULL
);