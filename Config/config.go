package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	stringTrue        = "true"
	defaultPort       = "10000"
)

var (
	BotToken    string  // bot token obtained from @botfather
	Port        string  // port on which the webserver should run
)

func init() {
	err := godotenv.Load()
	if err == nil {
		fmt.Println("configs loaded from .env file")
	}

	BotToken = os.Getenv("BOT_TOKEN")
	Port = stringEnviron("PORT", defaultPort)
}

// int64Environ gets a environment variable as an int64.
func int64Environ(envVar string, defaultVal ...int64) int64 {
	s := os.Getenv(envVar)
	if s == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}

		return 0
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("config.int64environ: failed to parse %s: %v\n", envVar, err)
	}

	return i
}

// stringEnviron gets a environment variable.
func stringEnviron(envVar string, defaultVal ...string) string {
	s := os.Getenv(envVar)
	if s == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}

		return ""
	}

	return s
}

// int64ListEnviron returns a slice of int64 for an environment variable.
func int64ListEnviron(envVar string) (result []int64) {
	s := os.Getenv(envVar)
	if s == "" {
		return result
	}

	for _, str := range strings.Split(s, " ") {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			fmt.Printf("config.int64listenviron: failed to parse %s: %v\n", envVar, err)
		}

		result = append(result, i)
	}

	return result
}
