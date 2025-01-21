package s3_services

import (
	"citizen_system_back/database"
	"citizen_system_back/models"
	s3_types "citizen_system_back/modules/s3/types"
	"citizen_system_back/utils/response"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NewS3Client creates and returns a new S3 client instance
func NewS3Client() (*s3_types.S3Client, error) {
	// Load environment variables
	region := os.Getenv("AWS_S3_REGION")
	accessKeyID := os.Getenv("AWS_S3_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_S3_SECRET_ACCESS_KEY")

	// Create custom credentials provider
	credProvider := credentials.NewStaticCredentialsProvider(
		accessKeyID,
		secretAccessKey,
		"", // Session token (optional)
	)

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credProvider),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS SDK configuration: %v", err)
	}

	// Create and return new S3 client
	return &s3_types.S3Client{
		Client: s3.NewFromConfig(cfg),
	}, nil
}

// HeadObject retrieves metadata of an object in the specified S3 bucket.
func HeadObject(client *s3_types.S3Client, bucketName, key string) (*s3.HeadObjectOutput, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	response, err := client.Client.HeadObject(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("error getting metadata for %s in bucket %s: %v", key, bucketName, err)
	}
	return response, nil
}

// GetObjectStream retrieves an object from S3 and returns its body stream, content type, and metadata.
func GetObjectStream(client *s3.Client, bucketName, key string) (io.ReadCloser, string, map[string]string, error) {
	// Create the input for the GetObject API call
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	// Call the GetObject API
	response, err := client.GetObject(context.TODO(), input)
	if err != nil {
		return nil, "", nil, fmt.Errorf("error getting object stream for %s in bucket %s: %v", key, bucketName, err)
	}

	// Return the body (readable stream), content type, and metadata
	return response.Body, aws.ToString(response.ContentType), response.Metadata, nil
}

func PreviewImageByID(c *gin.Context, imgID uuid.UUID) {
	var db = database.GetDB()
	s3Client, err := NewS3Client()

	var resImageData models.AssetAttachment
	result := db.Where("img_id = ?", imgID).First(&resImageData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, response.NewResponse(
				"Image not found",
				nil,
			))
			return
		}
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"Database error",
			result.Error,
		))
		return
	}

	// Check if Key exists
	if resImageData.Key == nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(
			"Image key not found",
			fmt.Errorf("no storage key found for image ID: %s", imgID),
		))
		return
	}

	var bucket, key string
	// Check if ImageFrom exists and equals 2 or 3
	if resImageData.ImageFrom != nil && (*resImageData.ImageFrom == 2 || *resImageData.ImageFrom == 3) {
		// Handle S3 path format "s3://bucket/path"
		path := strings.TrimPrefix(*resImageData.Key, "s3://")
		parts := strings.SplitN(path, "/", 2)
		if len(parts) != 2 {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse(
				"Invalid S3 path format",
				fmt.Errorf("invalid S3 path format: %s", *resImageData.Key),
			))
			return
		}
		bucket = parts[0]
		key = parts[1]
	} else {
		// Use default bucket
		bucket = os.Getenv("AWS_BUCKET_NAME")
		key = *resImageData.Key
	}

	// Check if object exists
	_, err = HeadObject(s3Client, bucket, key)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewErrorResponse(
			"S3 object not found",
			err,
		))
		return
	}

	// Get object stream
	body, contentType, _, err := GetObjectStream(s3Client.Client, bucket, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse(
			"Failed to get object stream",
			err,
		))
		return
	}
	defer body.Close()

	// If FileType exists, use it as content type
	if resImageData.FileType != nil {
		contentType = *resImageData.FileType
	}

	// Set content disposition header if OriginalFileName exists
	if resImageData.OriginalFileName != nil {
		c.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", *resImageData.OriginalFileName))
	}

	// Set the appropriate content type and stream the file
	c.Header("Content-Type", contentType)
	c.DataFromReader(http.StatusOK, -1, contentType, body, nil)
}
