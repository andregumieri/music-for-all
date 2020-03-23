package main

import (
	"github.com/andregumieri/music-for-all/playlist/router"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(router.Router)
}
