package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/ansel1/merry/v2"
	"github.com/getsentry/sentry-go"
	config "github.com/granitebps/puasa-sunnah-api/configs"
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
