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

	vin, _ := sc.FindPid(&pid.Vin{})
	el, _ := sc.FindPid(&pid.EngineLoad{})
	rpm, _ := sc.FindPid(&pid.Rpm{})
	speed, _ := sc.FindPid(&pid.Speed{})

	_ = sc.AskPid(vin)
	for {
		_ = sc.AskPid(el)
		_ = sc.AskPid(rpm)
		_ = sc.AskPid(speed)
		time.Sleep(1000 * time.Millisecond)
	}
}
