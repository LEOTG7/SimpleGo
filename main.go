package main

import (
       "fmt"
       "net/http"
       "runtime/debug"
       "time"

       "github.com/LEOTG7/SimpleGo/plugins"

)

func main() {
	err := godotenv.Load(".env") // config.env is supported bcuz other repos use it for some reason
	if err != nil {
		fmt.Println("ERROR: load variables from .env file failed", err)
	}
	defer func() {
		if r := recover(); r != nil {
			// Print reason for panic + stack for some sort of helpful log output
			fmt.Println(r)
			fmt.Println(string(debug.Stack()))
		}
	}()

	
       go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(w, "Hi Hi")
		})

		http.ListenAndServe(":"+config.Port, nil)
	}()
	
	botToken := opts.BotToken
	if s := os.Getenv("BOT_TOKEN"); s != "" {
		botToken = s
	}
	bot, err := gotgbot.NewBot(botToken, &gotgbot.BotOpts{})
	if err != nil {
		logger.Fatal("create bot failed", zap.Error(err))
	}

