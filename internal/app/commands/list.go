package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) List(updMsg *tgbotapi.Message) {
	outMsgText := "Here all the products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outMsgText += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(updMsg.Chat.ID, outMsgText)

	_, _ = c.bot.Send(msg)
}
