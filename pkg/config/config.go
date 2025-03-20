package config

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

// loads the environment variables from the .env file passed as a flag
// if no filename is given, looks for a file named .env in the current directory
func LoadEnvironment() (err error) {

	log.Println("Loading environment configuration...")

	var envFile string

	flag.StringVar(&envFile, "e", "", "environment configuration file")

	flag.Parse()

	if envFile == "" {
		envFile = ".env"
	}

	err = godotenv.Overload(envFile)

	return

}
