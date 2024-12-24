package main

import (
	database "api-channel/pkg/db"
	server "api-channel/pkg/server"
	"context"
	"flag"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/policy"
)

func main() {

	port := flag.Int("p", 9090, "The port of server")
	flag.Parse()

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	bucketName := "uploads"
	// location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}) // minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	err = minioClient.SetBucketPolicy(ctx, bucketName, string(policy.BucketPolicyReadOnly))
	if err != nil {
		log.Fatalln("Failed to set bucket policy:", err)
	}
	log.Println("Bucket is now public:", bucketName)

	s := server.NewServer(server.Config{
		FileServer: minioClient,
		Database:   db,
		Port:       *port,
		Options: server.ServerOptions{
			MaxTimeout: time.Second * 10,
		},
	})

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
