package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ansel1/merry/v2"
	"github.com/getsentry/sentry-go"
	appconfig "github.com/granitebps/puasa-sunnah-api/configs"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/resend/resend-go/v3"
	"github.com/spf13/viper"
)

func main() {
	log.Println("Starting backup process...")

	appconfig.SetupConfig(".env")

	database := viper.GetString(constants.DB_NAME)
	dateTime := time.Now().Format("2006-01-02_15-04-05")
	sqlFileName := fmt.Sprintf("%s_%s.sql", dateTime, database)

	err := dumpSqlFile(sqlFileName)
	if err != nil {
		err = merry.Wrap(err)
		sentry.CaptureException(err)
		os.Exit(1)
	}

	err = sendSqlToS3(sqlFileName)
	if err != nil {
		err = merry.Wrap(err)
		sentry.CaptureException(err)
		os.Exit(1)
	}

	err = sendEmailNotification(sqlFileName)
	if err != nil {
		err = merry.Wrap(err)
		sentry.CaptureException(err)
		os.Exit(1)
	}

	log.Println("Backup process finished.")
}

func dumpSqlFile(fileName string) (err error) {
	host := viper.GetString(constants.DB_HOST)
	username := viper.GetString(constants.DB_USER)
	password := viper.GetString(constants.DB_PASS)
	database := viper.GetString(constants.DB_NAME)

	_, err = exec.LookPath("mysqldump")
	if err != nil {
		log.Println("Error: mysqldump is not available in PATH.")
		err = merry.Wrap(err)
		return
	}

	cmd := exec.Command("mysqldump", "-h", host, "-u", username, "-p"+password, database) // #nosec G204

	outputFile, err := os.Create(filepath.Clean(fileName))
	if err != nil {
		log.Println("Failed to create dump file")
		err = merry.Wrap(err)
		return
	}
	defer outputFile.Close()

	cmd.Stdout = outputFile

	err = cmd.Run()
	if err != nil {
		log.Printf("Failed to execute mysqldump: %s\n", err)
		err = merry.Wrap(err)
		return
	}

	log.Printf("Database dumped successfully to dump %s\n", fileName)
	return nil
}

func sendSqlToS3(fileName string) (err error) {
	region := viper.GetString(constants.AWS_DEFAULT_REGION)
	bucket := viper.GetString(constants.AWS_BUCKET)
	accessKey := viper.GetString(constants.AWS_ACCESS_KEY_ID)
	secretKey := viper.GetString(constants.AWS_SECRET_ACCESS_KEY)
	endpoint := viper.GetString(constants.AWS_ENDPOINT)
	appEnv := viper.GetString(constants.APP_ENV)

	ctx := context.Background()

	optFns := []func(*config.LoadOptions) error{
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	}

	cfg, err := config.LoadDefaultConfig(ctx, optFns...)
	if err != nil {
		err = merry.Wrap(err)
		log.Println("Error creating config:", err)
		return
	}

	clientOptFns := []func(*s3.Options){}
	if endpoint != "" {
		clientOptFns = append(clientOptFns, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(endpoint)
		})
	}

	svc := s3.NewFromConfig(cfg, clientOptFns...)

	file, err := os.Open(filepath.Clean(fileName))
	if err != nil {
		err = merry.Wrap(err)
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		err = merry.Wrap(err)
		log.Println("Error reading file:", err)
		return
	}

	_, err = svc.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fmt.Sprintf("PSN_%s/%s", appEnv, fileName)),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		err = merry.Wrap(err)
		fmt.Println("Error uploading file:", err)
		return
	}

	err = os.Remove(fileName)
	if err != nil {
		err = merry.Wrap(err)
		log.Println("Error deleting file:", err)
		return
	}

	log.Printf("File uploaded to AWS S3 bucket: %s\n", bucket)
	return
}

func sendEmailNotification(fileName string) (err error) {
	apiKey := viper.GetString(constants.RESEND_API_KEY)
	from := viper.GetString(constants.RESEND_FROM)
	to := viper.GetString(constants.RESEND_TO)
	appName := viper.GetString(constants.APP_NAME)
	appEnv := viper.GetString(constants.APP_ENV)

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    from,
		To:      []string{to},
		Subject: fmt.Sprintf("Successful new backup of %s (%s)", appName, appEnv),
		Html:    fmt.Sprintf("<p>Hello! Great news, a new backup of <strong>%s</strong> (%s) was successfully created with filename <code>%s</code>!</p>", appName, appEnv, fileName),
	}

	_, err = client.Emails.Send(params)
	if err != nil {
		err = merry.Wrap(err)
		log.Println("Error sending email:", err)
		return
	}

	log.Println("Email notification sent successfully")
	return
}
