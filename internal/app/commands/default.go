package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) Default(updMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(updMsg.Chat.ID, updMsg.Text)
	msg.ReplyToMessageID = updMsg.MessageID

	_, _ = c.bot.Send(msg)
}
