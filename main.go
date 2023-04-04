package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/momentohq/client-sdk-go/auth"
	"github.com/momentohq/client-sdk-go/config"
	"github.com/momentohq/client-sdk-go/momento"
	"os"
)

func main() {
	var (
		cacheName          = os.Getenv("CACHE_NAME")
		topicName          = os.Getenv("TOPIC_NAME")
		functionTargetName = os.Getenv("FUNCTION_TARGET_NAME")
	)

	if cacheName == "" {
		panic("CACHE_NAME undefined")
	}
	if topicName == "" {
		panic("TOPIC_NAME undefined")
	}
	if functionTargetName == "" {
		panic("FUNCTION_TARGET_NAME undefined")
	}

	credProvider, err := auth.NewEnvMomentoTokenProvider("MOMENTO_AUTH_TOKEN")
	if err != nil {
		panic(err)
	}

	topicClient, err := momento.NewTopicClient(config.InRegionLatest(), credProvider)
	if err != nil {
		panic(err)
	}

	h := handler{
		momentoClient:      topicClient,
		awsLambdaClient:    lambda.New(session.Must(session.NewSession())),
		cacheName:          cacheName,
		topicName:          topicName,
		functionTargetName: functionTargetName,
	}
	h.handler(context.Background())
}
