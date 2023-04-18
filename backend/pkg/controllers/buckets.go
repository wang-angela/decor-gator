package controllers

// Used code from this tutorial for bucket operations: https://www.youtube.com/watch?v=gzBnrBK1P5Q&t=710s

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
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
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})))
}

func ListBuckets() (resp *s3.ListBucketsOutput) {
	resp, err := s3sess.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}

	return resp
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

func UploadObject(filename string) (resp *s3.PutObjectOutput) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Uploading:", filename)
	resp, err = s3sess.PutObject(&s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(strings.Split(filename, "/")[1]),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})

	if err != nil {
		panic(err)
	}

	return resp
}

func ListObjects() (resp *s3.ListObjectsV2Output) {
	resp, err := s3sess.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(BUCKET_NAME),
	})

	if err != nil {
		panic(err)
	}

	return resp
}

func GetObject(filename string) {
	fmt.Println("Downloading: ", filename)

	resp, err := s3sess.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename),
	})

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		panic(err)
	}
}

func DeleteObject(filename string) (resp *s3.DeleteObjectOutput) {
	fmt.Println("Deleting: ", filename)
	resp, err := s3sess.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(filename),
	})

	if err != nil {
		panic(err)
	}

	return resp
}

func UploadObjectHelper(w http.ResponseWriter, r *http.Request) {
	UploadObject(mux.Vars(r)["filename"])
}

func ListBucketsHelper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	resp := ListBuckets()

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func DeleteObjectHelper(w http.ResponseWriter, r *http.Request) {
	DeleteObject(mux.Vars(r)["filename"])
}
