```mermaid
erDiagram
    JOBS {
        uuid job_id PK
        timestamp requested_at
        timestamp updated_at
        enum status
    }
```
