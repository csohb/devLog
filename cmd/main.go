package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
)

const Env = ".env"
const Config = "DEV_LOG_CONFIG"

var configFile string

func init() {
	if len(os.Args) == 2 {
		configFile = os.Args[1]
	} else {
		err := godotenv.Load(Env)
		if err != nil {
			logrus.Fatalf("env file does not exists : %+v", err)
			panic(err)
		}
		if confPath := os.Getenv(Config); confPath == "" {
			logrus.Fatal("confPath is not found")
		} else {
			configFile = confPath
		}
	}
	fmt.Printf("configFile path : %+v\n", configFile)
}

func main() {
	e := echo.New()
	e.Start("localhost:8080")

}
