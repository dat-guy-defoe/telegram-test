package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (c *Commander) Get(updMsg *tgbotapi.Message) {
	args := updMsg.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(updMsg.Chat.ID, "Incorrect argument!")

		_, _ = c.bot.Send(msg)

		return
	}

	var msg tgbotapi.MessageConfig
	product, err := c.productService.Get(arg)
	if err != nil {
		msg = tgbotapi.NewMessage(updMsg.Chat.ID, fmt.Sprintf("Error: %s", err))
	} else {
		msg = tgbotapi.NewMessage(updMsg.Chat.ID, product.Title)
	}

	_, _ = c.bot.Send(msg)
}
