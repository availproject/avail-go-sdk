package main

import (
	"os"

	"github.com/availproject/avail-go-sdk/examples"
	"github.com/sirupsen/logrus"
)

func main() {
	// Set log level based on the environment variable
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel // Default to INFO if parsing fails
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	examples.Run_data_submission()
}
