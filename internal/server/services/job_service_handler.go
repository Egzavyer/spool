package services

import (
	"context"
	"errors"
	jobv1 "spool/gen/job/v1"
	"time"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type JobServiceHandler struct {
	DbPool *pgxpool.Pool
}

var (
	ErrUnknownStatus = errors.New("job has invalid status")

	ErrJobTypeNotFound   = errors.New("job type not found")
	ErrJobTypeNameExists = errors.New("job type name already exists")
)

func mapDatabaseError(err error) *connect.Error {
	var (
		pgErr  *pgconn.PgError
		newErr error
	)
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // Unique constraint violation
			{
				switch pgErr.ConstraintName {
				case "uk_job_types_name":
					newErr = ErrJobTypeNameExists
				default:
					newErr = err
				}
				return connect.NewError(connect.CodeAlreadyExists, newErr)
			}
		case "23503": // Foreign key constraint violated
			{

				switch pgErr.ConstraintName {
				case "fk_jobs_job_type_id":
					newErr = ErrJobTypeNotFound
				default:
					newErr = err
				}
				return connect.NewError(connect.CodeFailedPrecondition, newErr)
			}
		}
	}
	return connect.NewError(connect.CodeUnknown, err)
}

func pgStatusToProto(tempStatus string) (status jobv1.JobStatus, err *connect.Error) {
	switch tempStatus {
	case "pending":
		{
			return jobv1.JobStatus_JOB_STATUS_PENDING, nil
		}
	case "completed":
		{
			return jobv1.JobStatus_JOB_STATUS_COMPLETED, nil
		}
	case "failed":
		{
			return jobv1.JobStatus_JOB_STATUS_FAILED, nil
		}
	case "running":
		{
			return jobv1.JobStatus_JOB_STATUS_RUNNING, nil
		}
	default:
		{
			return jobv1.JobStatus_JOB_STATUS_UNSPECIFIED, connect.NewError(connect.CodeUnknown, ErrUnknownStatus)
		}
	}
}

func (s *JobServiceHandler) CreateJobType(
	ctx context.Context,
	req *jobv1.CreateJobTypeRequest,
) (*jobv1.CreateJobTypeResponse, error) {
	query := `
		INSERT INTO Job_Types(job_type_name)
		VALUES ($1) 
		RETURNING job_type_id, job_type_name`
	var jobType jobv1.JobType
	if err := s.DbPool.QueryRow(ctx, query, req.Name).Scan(&jobType.Id, &jobType.Name); err != nil {
		return nil, mapDatabaseError(err)
	}
	res := &jobv1.CreateJobTypeResponse{JobType: &jobType}
	return res, nil
}

func (s *JobServiceHandler) CreateJob(
	ctx context.Context,
	req *jobv1.CreateJobRequest,
) (*jobv1.CreateJobResponse, error) {
	query := `
		WITH inserted_rows AS (
			INSERT INTO Jobs(job_type_id)
			VALUES ($1)
			RETURNING job_id, requested_at, updated_at, status, job_type_id
		)	
		
		SELECT i.job_id, i.requested_at, i.updated_at, i.status, j.job_type_id, j.job_type_name
		FROM inserted_rows AS i
		JOIN job_types AS j
		ON j.job_type_id = i.job_type_id`
	var (
		job             jobv1.Job
		tempRequestedAt time.Time
		tempUpdatedAt   time.Time
		tempStatus      string
		tempJobType     jobv1.JobType
	)
	if err := s.DbPool.QueryRow(ctx, query, req.JobTypeId).Scan(&job.Id, &tempRequestedAt, &tempUpdatedAt, &tempStatus, &tempJobType.Id, &tempJobType.Name); err != nil {
		return nil, mapDatabaseError(err)
	}

	// Cannot scan timestamptz directly into a Protobuf timestamp,
	// instead scan into time.Time and initialize a new Protobuf timestamp with it
	job.RequestedAt = timestamppb.New(tempRequestedAt)
	job.UpdatedAt = timestamppb.New(tempUpdatedAt)

	// Cannot scan Postgres enum directly into a Go enum
	status, err := pgStatusToProto(tempStatus)
	if err != nil {
		return nil, err
	}
	job.Status = status

	job.JobType = &tempJobType

	res := &jobv1.CreateJobResponse{Job: &job}
	return res, nil
}

func (s *JobServiceHandler) GetJobType(
	_ context.Context,
	req *jobv1.GetJobTypeRequest,
) (*jobv1.GetJobTypeResponse, error) {
	res := &jobv1.GetJobTypeResponse{
		JobType: &jobv1.JobType{
			Id:   1,
			Name: "test",
		},
	}
	return res, nil
}

func (s *JobServiceHandler) ListJobTypes(
	_ context.Context,
	req *jobv1.ListJobTypesRequest,
) (*jobv1.ListJobTypesResponse, error) {
	res := &jobv1.ListJobTypesResponse{}
	return res, nil
}

func (s *JobServiceHandler) GetJob(
	_ context.Context,
	req *jobv1.GetJobRequest,
) (*jobv1.GetJobResponse, error) {
	res := &jobv1.GetJobResponse{}
	return res, nil
}

func (s *JobServiceHandler) ListJobs(
	_ context.Context,
	req *jobv1.ListJobsRequest,
) (*jobv1.ListJobsResponse, error) {
	res := &jobv1.ListJobsResponse{}
	return res, nil
}
