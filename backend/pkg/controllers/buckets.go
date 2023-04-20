package controllers

// Used code from this tutorial for bucket operations: https://www.youtube.com/watch?v=gzBnrBK1P5Q&t=710s

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/decor-gator/backend/pkg/models"
)

var (
	s3sess *s3.S3
)

const (
	BUCKET_NAME = "decorgatorbucket"
	REGION      = "us-east-1"
)

func InitAWSSession() {
	s3sess = s3.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(REGION),
		Credentials: credentials.NewStaticCredentials("AKIAWYOMFPS7DV5H73C6", "ihgJi8CeqWOtL9IR8rTFDuEnojs+3jtzEabdnXd2", ""),
	})))
}

func CreateBucket() (resp *s3.CreateBucketOutput) {
	resp, err := s3sess.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(BUCKET_NAME),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println("Bucket name already in use!")
				panic(err)
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println("Bucket exists and is owned by you!")
			default:
				panic(err)
			}
		}
	}

	return resp
}

func UploadObject(post models.Post) (resp *s3.PutObjectOutput) {
	p, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	resp, err = s3sess.PutObject(&s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(bytes.NewReader(p)),
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(string(post.ID.Hex())),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	if err != nil {
		panic(err)
	}

	return resp
}

func GetObject(id string, post models.Post) error {
	resp, err := s3sess.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(id),
	})

	if err != nil {
		return err
	}

	value := json.NewDecoder(resp.Body).Decode(post)

	return value
}

func DeleteObject(key string) (resp *s3.DeleteObjectOutput) {
	resp, err := s3sess.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(key),
	})

	if err != nil {
		panic(err)
	}

	return resp
}
