package main

// Библиотеки, нужные программе
import (
	"log"
)

func main() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	// инициализируем канал, куда будут прилетать обновления от API
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	// читаем обновления из канала
	for update := range updates {
		if update.Message == nil {
			continue
		}
		// Созадаем сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		//Ping-pong
		if update.Message.Text == "/ping" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "pong")
			// В остальных случаях
		} else if update.Message != nil {
			// Ответить на сообщение его копией
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		}

		// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		// Созадаем сообщение

		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
