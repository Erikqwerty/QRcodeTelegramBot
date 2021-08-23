package main

import (
	"fmt"
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("your token")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		qrcodeDeiler(update.Message.Text)
		messegeQRforUser(bot, update)
	}

}

func messegeQRforUser(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	photoBytes, err := ioutil.ReadFile("./qr.png")
	if err != nil {
		panic(err)
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}
	_, err = bot.Send(tgbotapi.NewPhotoUpload(int64(update.Message.Chat.ID), photoFileBytes))
}
func qrcodeDeiler(messege string) {
	msg := messege
	err := qrcode.WriteFile(msg, qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("write error")
	}
}
