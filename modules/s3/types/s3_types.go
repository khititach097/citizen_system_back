package s3_types

import "github.com/aws/aws-sdk-go-v2/service/s3"

// S3Client wraps the AWS S3 client.
type S3Client struct {
	Client *s3.Client
}
