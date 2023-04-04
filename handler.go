package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/momentohq/client-sdk-go/momento"
)

type handler struct {
	momentoClient      momento.TopicClient
	awsLambdaClient    *lambda.Lambda
	cacheName          string
	topicName          string
	functionTargetName string
}

func (h handler) handler(ctx context.Context) {

	// Init momento topic subscription
	sub, err := h.momentoClient.Subscribe(ctx, &momento.TopicSubscribeRequest{
		CacheName: h.cacheName,
		TopicName: h.topicName,
	})
	if err != nil {
		log.Fatalf("fatal error occurred subscribing to momento topic "+
			"err=%+v", err,
		)
		return
	}

	log.Println("Successfully created subscription listening for new messages on topic. " +
		"topic=" + h.topicName)

	for {
		item, err := sub.Item(ctx)
		if err != nil {
			log.Printf("error occurred reading item from momento topic. "+
				"topic=%s "+
				"err=%+v",
				h.topicName, err,
			)
		}

		var invokeErr error
		switch msg := item.(type) {
		case momento.Bytes:
			_, invokeErr = h.awsLambdaClient.InvokeWithContext(ctx, &lambda.InvokeInput{
				FunctionName:   aws.String(h.functionTargetName),
				InvocationType: aws.String(lambda.InvocationTypeEvent),
				Payload:        msg,
			})
			if err != nil {
				return
			}
		case momento.String:
			_, invokeErr = h.awsLambdaClient.InvokeWithContext(ctx, &lambda.InvokeInput{
				FunctionName:   aws.String(h.functionTargetName),
				InvocationType: aws.String(lambda.InvocationTypeEvent),
				Payload:        []byte(msg),
			})
		}
		if invokeErr != nil {
			log.Printf("error occurred triggering target lambda. "+
				"cache=%s "+
				"topic=%s "+
				"lambda_target=%s "+
				"err=%+v",
				h.cacheName, h.topicName, h.functionTargetName, invokeErr,
			)
		}
	}
}
