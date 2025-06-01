package plugins

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello, @%s! I'm alive ✅", ctx.EffectiveUser.Username), nil)
	return err
}
