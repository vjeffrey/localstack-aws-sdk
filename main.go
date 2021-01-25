package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"

	"fmt"
)

// pip3 install localstack
// localstack start

func main() {
	fmt.Println("\n**PLEASE ENSURE LOCALSTACK IS RUNNING (localstack start)**")

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Print(err)
	}
	localResolverFn := func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               "http://localhost:4566",
			SigningRegion:     region,
			HostnameImmutable: true,
		}, nil
	}
	cfg.EndpointResolver = aws.EndpointResolverFunc(localResolverFn)

	cfg.Region = "us-east-1"
	svc := kms.NewFromConfig(cfg)

	keyList, err := svc.ListKeys(ctx, &kms.ListKeysInput{})
	if err != nil {
		fmt.Print(err)
	}
	keyIds := []string{}
	for _, key := range keyList.Keys {
		keyIds = append(keyIds, *key.KeyArn)
	}
	fmt.Println(keyIds)
}
