package envloader

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Settings has config values
type Settings struct {
	ENV         string
	TOKEN       string
	LEAVE_LIMIT int
	BOT_ID      string
}

// LoadConfig will load settings
// returns Setting struct
func LoadConfig() *Settings {
	loadEnv()
	confs := setConfig()
	return &confs
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func setConfig() Settings {
	env := getMode()
	token := getToken()
	limit := getLimit()
	id := getID()
	return Settings{env, token, limit, id}
}

func getID() string {
	const BotID = "CLIENT_ID"
	id := os.Getenv(BotID)
	if id == "" {
		panic(errors.New(BotID + " is empty"))
	}
	return id
}

func getLimit() int {
	const LeaveMaxCount = "LEAVE_MAX_COUNT"
	limit := os.Getenv(LeaveMaxCount)
	num, err := strconv.Atoi(limit)
	if err != nil {
		panic(err)
	}
	return num
}

func getToken() string {
	const Token = "TOKEN"

	token := os.Getenv(Token)
	if token == "" {
		panic(errors.New(Token + " is empty"))
	}
	return token
}

func getMode() string {
	const ENV = "ENV"
	envTypes := make(map[string]bool)
	envTypes["production"] = true
	envTypes["development"] = true
	envTypes["test"] = true

	env := os.Getenv(ENV)
	if !envTypes[env] {
		panic(errors.New(env + " is invalid value of " + ENV))
	}
	return env
}
