CREATE TYPE job_status AS ENUM ('pending', 'completed', 'failed', 'running');

CREATE TABLE jobs (
    job_id UUID PRIMARY KEY,
    request_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    status job_status DEFAULt 'pending'
);