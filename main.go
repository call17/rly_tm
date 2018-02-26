package main

import (
"gopkg.in/syfaro/telegram-bot-api.v4"
"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("525083131:AAG3PhZ35aAHt1qYo6UKa1dBWHzoaEgToj0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		var rs string;

		if update.Message == nil {
			continue
		}else {
			rs = templateResponseGenerator(update.Message.Text)
		}


		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, rs)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func templateResponseGenerator(rq string) string {

	var hello string = "Привет, я Галочка , и я тест манагер. Я расскажу тебе все о моей работе, а также много интересного."

	if "/start" == rq {
		return hello
	}else {
		return "Unexpected"
	}

}