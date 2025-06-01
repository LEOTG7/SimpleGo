package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/LEOTG7/SimpleGo/plugins"
	"github.com/LEOTG7/SimpleGo/Config"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func main() {
	// Recover from panics and log the error + stack trace
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println(string(debug.Stack()))
		}
	}()

	// Web server for uptime pings (e.g., Koyeb, Render)
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprint(w, "Bot is running!")
		})
		_ = http.ListenAndServe(":"+config.Port, nil)
	}()

	if config.BotToken == "" {
		panic("No BOT_TOKEN provided")
	}

	bot, err := gotgbot.NewBot(config.BotToken, nil)
	if err != nil {
		panic("Failed to create bot: " + err.Error())
	}

	// To make sure no other instance of the bot is running
	_, err = bot.GetUpdates(&gotgbot.GetUpdatesOpts{})
	if err != nil {
		fmt.Println("Waiting 10s because: " + err.Error())
		time.Sleep(10 * time.Second)
	}

	updater := ext.NewUpdater(plugins.Dispatcher, nil)

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
	})
	if err != nil {
		panic("Failed to start polling: " + err.Error())
	}

	fmt.Printf("@%s started.\n", bot.User.Username)

	updater.Idle()
}
