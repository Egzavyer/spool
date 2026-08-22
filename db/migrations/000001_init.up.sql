DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'job_status') THEN
        CREATE TYPE job_status AS ENUM ('pending', 'completed', 'failed', 'running');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS job_types (
    job_type_id SMALLSERIAL PRIMARY KEY,
    job_type_name VARCHAR(63) CONSTRAINT uk_job_types_name UNIQUE 
);

CREATE TABLE IF NOT EXISTS jobs (
    job_id UUID PRIMARY KEY DEFAULT uuidv7(),
    requested_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    status job_status DEFAULt 'pending' NOT NULL,
    job_type_id SMALLINT CONSTRAINT fk_jobs_job_type_id REFERENCES job_types(job_type_id) ON DELETE SET NULL
);