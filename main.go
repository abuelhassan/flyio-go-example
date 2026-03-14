package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	ctx := context.Background()
	cl, err := s3Client(ctx)
	if err != nil {
		log.Println(err)
	}

	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		log.Println("BUCKET_NAME is not set")
	}

	text, err := readS3Object(ctx, cl, bucketName, "file.txt")
	if err != nil {
		log.Println(err)
	}
	body := fmt.Sprintf(`<html><body>
			<h1>Welcome to the Fly.io Go example!</h1>
			<p style="white-space: pre-wrap; word-wrap: break-word;">%s</p>
		</body></html>`, text)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = fmt.Fprint(w, body)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func s3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load Tigris config: %w", err)
	}

	fmt.Printf("cfg.Region: %s\n", cfg.Region)

	return s3.NewFromConfig(cfg), nil
}

func readS3Object(ctx context.Context, s3Client *s3.Client, bucketName, fileName string) (string, error) {
	out, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &fileName,
	})
	if err != nil {
		return "", fmt.Errorf("failed to download object: %w", err)
	}

	text, err := io.ReadAll(out.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read Tigris object: %w", err)
	}

	return string(text), nil
}
