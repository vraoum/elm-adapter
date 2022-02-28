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

	vin, _ := sc.FindPid(&pid.Vin{})
	el, _ := sc.FindPid(&pid.EngineLoad{})
	ect, _ := sc.FindPid(&pid.EngineCoolantTemperature{})
	fp, _ := sc.FindPid(&pid.FuelPressure{})
	rpm, _ := sc.FindPid(&pid.Rpm{})
	speed, _ := sc.FindPid(&pid.Speed{})

	_ = sc.AskPid(vin)
	for {
		time.Sleep(1000 * time.Millisecond)
		_ = sc.AskPid(el)
		_ = sc.AskPid(ect)
		_ = sc.AskPid(fp)
		_ = sc.AskPid(rpm)
		_ = sc.AskPid(speed)
	}
}
