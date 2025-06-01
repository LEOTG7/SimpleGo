package plugins

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

const basicCommandsGroup = 1

var Dispatcher = ext.NewDispatcher(&ext.DispatcherOpts{
	Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
		fmt.Println("An error occurred while handling update:", err.Error())
		return ext.DispatcherActionNoop
	},
})


func init() {
	Dispatcher.AddHandlerToGroup(handlers.NewCommand("start", Start), basicCommandsGroup)
}
