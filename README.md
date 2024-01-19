# AWS Lambda Go Application :zap:

This repository contains a simple AWS Lambda function written in Go. The Lambda function receives an event, processes it, and returns a response. The function is triggered by the AWS Lambda service.

## Overview :mag_right:

The Lambda function takes an event of type `MyEvent` with fields "Name" and "Age" and responds with a message indicating the provided name and age.

### Input Event Structure

```json
{
  "What is your name?": "Anjas",
  "How old are you?": 20
}
```

### Output Response Structure

```json
{
  "Answer: ": "Anjas is 20 years old"
}
```

## Code Explanation :rocket:

### `main.go`

The `main.go` file serves as the entry point for the AWS Lambda Go application.

- **Imports:**
  The necessary packages are imported, including `fmt` for formatted I/O and `github.com/aws/aws-lambda-go/lambda` for AWS Lambda functionality.

- **Event Structures:**
  Two structures, `MyEvent` and `MyResponse`, are defined to represent the input and output data formats for the Lambda function.

- **Handler Function:**
  The `handleLambdaEvent` function processes the Lambda event and generates a response. It formats a message based on the provided name and age.

- **Main Function:**
  The `main` function initiates the Lambda function by calling `lambda.Start()`, passing the `handleLambdaEvent` function as the handler.

## Usage :wrench:

Invoke the Lambda function with an event similar to the provided `MyEvent` structure:

```bash
aws lambda invoke --function-name YourLambdaFunctionName --payload '{"What is your name?": "Anjas", "How old are you?": 20}' output.json
```

## Closing Notes 📝

Feel free to adjust the configuration, and if you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request.

Happy coding! 🤖👨‍💻
