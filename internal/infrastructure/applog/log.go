package applog

import (
	"fmt"
	"io"
	"time"

	"github.com/gofiber/fiber/v2/log"

	"os"
)

func InitLog() {
	makeLogFile()
}

func makeLogFile() {
	err := os.MkdirAll("log", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating log directory:", err)
		return
	}
	file, _ := os.OpenFile("log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
}

func backupLog() {
	err := os.MkdirAll("backup_log/app_log", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	content, err := os.ReadFile("app.log")

	date := time.Now().Format("02-01-2006_15-04-05")
	backupPath := "backup_log/app_log/" + date

	backupFile, err := os.Create(backupPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer backupFile.Close()

	// Write the content to the file
	_, err = backupFile.WriteString(string(content))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func backupLogger() {
	err := os.MkdirAll("backup_log/logger", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	content, err := os.ReadFile("logger.log")

	date := time.Now().Format("02-01-2006_15-04-05")
	backupPath := "backup_log/logger/" + date

	backupFile, err := os.Create(backupPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer backupFile.Close()

	// Write the content to the file
	_, err = backupFile.WriteString(string(content))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
