package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) List(updMsg *tgbotapi.Message) {
	outMsgText := "Here all the products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outMsgText += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(updMsg.Chat.ID, outMsgText)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", "some data"),
			tgbotapi.NewInlineKeyboardButtonData("Prev page", "some prev data"),
		),
	)

	_, _ = c.bot.Send(msg)
}
