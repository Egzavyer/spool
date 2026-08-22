package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"spool/config"
	jobv1 "spool/gen/job/v1"
	"spool/gen/job/v1/jobv1connect"
)

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	client := jobv1connect.NewJobServiceClient(
		http.DefaultClient,
		fmt.Sprintf("http://localhost%v", cfg.Server.Addr),
	)
	res, err := client.CreateJob(
		context.Background(),
		&jobv1.CreateJobRequest{JobTypeId: 1},
	)
	// res, err := client.CreateJobType(
	// 	context.Background(),
	// 	&jobv1.CreateJobTypeRequest{Name: "test"},
	// )
	if err != nil {
		log.Printf("Error occurred: %v", err)
		return
	}

	log.Println(res.Job)
}
