package test

import (
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/docker"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/stretchr/testify/assert"
)

func TestDockerGin(t *testing.T) {
	t.Parallel()
	// Configure the tag to use on the Docker image.
	tag := "gruntwork/docker-gin"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	// Build the Docker image.
	docker.Build(t, "../", buildOptions)

	//defer docker.Stop(t, nil, &docker.StopOptions{Time: 10})

	// Run the Docker image, read the text file from it, and make sure it contains the expected output.
	opts := &docker.RunOptions{
		Detach:       true,
		OtherOptions: []string{"-p", "8080:8080"},
		Remove:       true,
	}
	output := docker.Run(t, tag, opts)

	expectedStatus := 200
	expectedBody := `{"message":"pong"}`
	maxRetries := 3
	timeBetweenRetries := 3 * time.Second

	http_helper.HttpGetWithRetry(t, "http://localhost:8080", nil, expectedStatus, expectedBody, maxRetries, timeBetweenRetries)
	assert.Equal(t, expectedBody, output)
}
