package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/ansel1/merry/v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/getsentry/sentry-go"
	config "github.com/granitebps/puasa-sunnah-api/configs"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/spf13/viper"
)

func main() {
	log.Println("Starting backup process...")

	config.SetupConfig(".env")

	database := viper.GetString("DB_NAME")
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
	host := viper.GetString("DB_HOST")
	username := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASS")
	database := viper.GetString("DB_NAME")

	// Check if mysqldump is available in PATH
	_, err = exec.LookPath("mysqldump")
	if err != nil {
		log.Println("Error: mysqldump is not available in PATH.")
		err = merry.Wrap(err)
		return
	}

	// Set the command and arguments
	// cmd := exec.Command("mysqldump", "-h", host, "-u", username, "-p"+password, database)
	cmd := exec.Command("mysqldump", "--socket", "/tmp/mysql_3306.sock", "-h", host, "-u", username, "-p"+password, database)

	// Create an output file for the dump
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Println("Failed to create dump file")
		err = merry.Wrap(err)
		return
	}
	defer outputFile.Close()

	// Redirect command output to the file
	cmd.Stdout = outputFile

	// Run the command
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
	region := viper.GetString("AWS_DEFAULT_REGION")
	bucket := viper.GetString("AWS_BUCKET")
	accessKey := viper.GetString("AWS_ACCESS_KEY_ID")
	secretKey := viper.GetString("AWS_SECRET_ACCESS_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		err = merry.Wrap(err)
		log.Println("Error creating session:", err)
		return
	}

	svc := s3.New(sess)

	file, err := os.Open(fileName)
	if err != nil {
		err = merry.Wrap(err)
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the contents of the file into a buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, file)
	if err != nil {
		err = merry.Wrap(err)
		log.Println("Error reading file:", err)
		return
	}

	// This uploads the contents of the buffer to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	// Delete the file after it's uploaded
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
	secret := viper.GetString("MAILGUN_SECRET")
	domain := viper.GetString("MAILGUN_DOMAIN")
	appName := viper.GetString("APP_NAME")
	appEnv := viper.GetString("APP_ENV")

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(domain, secret)

	sender := "info@granitebps.com"
	subject := fmt.Sprintf("Successful new backup of %s (%s)", appName, appEnv)
	body := fmt.Sprintf("Hello! Great news, a new backup of %s (%s) was successfully created with filename %s!", appName, appEnv, fileName)
	recipient := "granitebagas28@gmail.com"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	_, _, err = mg.Send(ctx, message)
	if err != nil {
		err = merry.Wrap(err)
		log.Println("Error sending email:", err)
		return
	}

	log.Println("Email notification sent successfully")

	return
}
