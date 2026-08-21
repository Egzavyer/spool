package main

import (
	"context"
	"log"
	"net/http"
	jobv1 "spool/gen/job/v1"
	"spool/gen/job/v1/jobv1connect"
)

func main() {
	client := jobv1connect.NewJobServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.GetJobType(
		context.Background(),
		&jobv1.GetJobTypeRequest{Id: 1},
	)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.JobType)
}
