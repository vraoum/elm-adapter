package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vraoum/elm-adapter/pkg/pid"
	"github.com/vraoum/elm-adapter/pkg/serial"
	"github.com/vraoum/elm-adapter/pkg/util"
	"time"
)

func main() {
	// Set the format for the logger
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05.000",
	})
	logrus.SetLevel(logrus.DebugLevel)

	// Load the environment file
	err := godotenv.Load()
	if err != nil {
		logrus.Error("Error loading .env file")
	}

	sc := serial.OpenSerial(util.GetEnv("SERIAL", ""))
	go sc.Read()

	vin := pid.GetVinInstance()
	speed := pid.GetSpeedInstance()
	rpm := pid.GetRpmInstance()

	_ = sc.AskPid(vin)
	for {
		_ = sc.AskPid(rpm)
		_ = sc.AskPid(speed)
		time.Sleep(100 * time.Millisecond)
	}

}
