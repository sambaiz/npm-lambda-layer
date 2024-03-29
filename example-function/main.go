package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if err := os.Setenv("NPM_CONFIG_USERCONFIG", "/opt/nodejs/.npmrc"); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}
	cmd := exec.Command("npm", "init", "-y")
	cmd.Dir = "/tmp"
	output, err := cmd.CombinedOutput()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%s %s", string(output), err.Error()),
			StatusCode: 500,
		}, nil
	}
	cmd = exec.Command("npm", "install", "express")
	cmd.Dir = "/tmp"
	output, err = cmd.CombinedOutput()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%s %s", string(output), err.Error()),
			StatusCode: 500,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       string(output),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
