package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

// MyEvent represents the input data structure for the Lambda function.
type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

// MyResponse represents the output data structure for the Lambda function.
type MyResponse struct {
	Message string `json:"Answer: "`
}

// handleLambdaEvent is the function that will be executed when the Lambda function is invoked.
func handleLambdaEvent(event MyEvent) (MyResponse, error) {
	// Formulate the response message based on the input event.
	return MyResponse{Message: fmt.Sprintf("%s is %d years old", event.Name, event.Age)}, nil
}

// main is the entry point for the Lambda function.
func main() {
	// Start the Lambda function, invoking handleLambdaEvent on each invocation.
	lambda.Start(handleLambdaEvent)
}
