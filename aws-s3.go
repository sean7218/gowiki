package main

import (

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)


func setupAWS() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("V17498-ICND-0001", "test-account"),
	})

	if err != nil {
		panic(err.Error())
	}


	_, err = sess.Config.Credentials.Get()

	if err != nil {
		panic(err.Error())
	}

	s3Svc := s3.New(session.New())
	s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("bucketName"),
		Key:    aws.String("keyName"),
	})

	result, err := s3Svc.GetObject(&s3.GetObjectInput{...})
	// result is a *s3.GetObjectOutput struct pointer
	// err is a error which can be cast to awserr.Error.

}

