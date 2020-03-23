package main

import (
	"github.com/andregumieri/music-for-all/tree/develop/api/router"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(router.Router)
}
