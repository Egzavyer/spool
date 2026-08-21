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
	res, err := client.CreateJobType(
		context.Background(),
		&jobv1.CreateJobTypeRequest{Name: "test"},
	)
	if err != nil {
		log.Printf("Error occurred: %v", err)
		return
	}

	log.Println(res.JobType)
}
