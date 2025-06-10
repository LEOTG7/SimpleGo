package plugins

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(bot *gotgbot.Bot, ctx *gotgbot.Context) error {
	_, err := bot.SendMessage(
		ctx.EffectiveChat.Id,
		fmt.Sprintf("Hello, @%s! I'm alive ✅", ctx.EffectiveUser.Username),
		&gotgbot.SendMessageOpts{
			ParseMode: gotgbot.ParseModeHTML,
			ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
					{
						{
							Text:         "Click Me!",
							CallbackData: "clicked",
						},
					},
				},
			},
		},
	)

	if err != nil {
		fmt.Println("Failed to send start message:", err)
	}

	return err
}
