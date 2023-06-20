package main

import (
	"log"
	"os"
	"telegram-test/internal/service/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, updMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(updMsg.Chat.ID,
		"/help - help\n"+
			"/list - list")

	_, _ = bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, updMsg *tgbotapi.Message, productService *product.Service) {
	outMsgText := "Here all the products: \n\n"
	products := productService.List()
	for _, p := range products {
		outMsgText += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(updMsg.Chat.ID, outMsgText)

	_, _ = bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, updMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(updMsg.Chat.ID, updMsg.Text)
	msg.ReplyToMessageID = updMsg.MessageID

	_, _ = bot.Send(msg)
}
