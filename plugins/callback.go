package plugins

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func HandleEditTextCallback(bot *gotgbot.Bot, ctx *gotgbot.Context) error {
	_, err := bot.EditMessageText(
		ctx.CallbackQuery.Message.Chat.Id,
		ctx.CallbackQuery.Message.MessageId,
		"😂😂 Nothing Here",
		&gotgbot.EditMessageTextOpts{
			ParseMode: gotgbot.ParseModeHTML,
		},
	)
	if err != nil {
		fmt.Println("Failed to edit message text:", err)
		return err
	}
package plugins
	_, err = ctx.CallbackQuery.Answer(bot, nil)
	if err != nil {
		fmt.Println("Failed to answer callback query:", err)
	}
	return err
}
