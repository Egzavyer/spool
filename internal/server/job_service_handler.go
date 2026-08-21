package server

import (
	"context"
	jobv1 "spool/gen/job/v1"

	"github.com/jackc/pgx/v5/pgxpool"
)

type JobServiceHandler struct {
	dbpool *pgxpool.Pool
}

func (s *JobServiceHandler) CreateJobType(
	_ context.Context,
	req *jobv1.CreateJobTypeRequest,
) (*jobv1.CreateJobTypeResponse, error) {

	query := `
		INSERT INTO Job_Types(job_type_name)
		VALUES ($1) 
		RETURNING job_type_id, job_type_name`
	var jobType jobv1.JobType
	if err := s.dbpool.QueryRow(context.Background(), query, req.Name).Scan(&jobType.Id, &jobType.Name); err != nil {
		return nil, err
	}
	res := &jobv1.CreateJobTypeResponse{JobType: &jobType}
	return res, nil
}

func (s *JobServiceHandler) CreateJob(
	_ context.Context,
	req *jobv1.CreateJobRequest,
) (*jobv1.CreateJobResponse, error) {
	res := &jobv1.CreateJobResponse{}
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
