package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) Help(updMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(updMsg.Chat.ID,
		"/help - help\n"+
			"/list - list")

	_, _ = c.bot.Send(msg)
}
