package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

const (
	chatID   = 123456
	youToken = "123456:AABBBCC"
)

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI(youToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = false

	link := fmt.Sprintf(`<a href="%s">[google]</a>`, "https://www.google.com.tw/")
	sendMsg(link)
}

func sendMsg(msg string) {
	NewMsg := tgbotapi.NewMessage(chatID, msg)
	NewMsg.ParseMode = tgbotapi.ModeHTML
	_, err := bot.Send(NewMsg)
	if err == nil {
		log.Printf("Send telegram message success")
	} else {
		log.Printf("Send telegram message error")
	}
}
