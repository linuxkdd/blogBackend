package main

import (
	"blog/blogBackend/telegram"
	"log"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6888730162:AAHJOoi10TgmLMFDXKbh2R3L3MM-t5dBKP0")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.ID != 0 {
			log.Printf("Received message in group: %s", update.Message.Text)
			r := regexp.MustCompile(`\b\w{25,}\b`)
			if r.MatchString(update.Message.Text) {
				newContent := r.FindStringSubmatch(update.Message.Text)
				filePath, fileName := telegram.EditFile(newContent[0])
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, telegram.UploadFile("https://106.53.179.226:8032/shjkadfcxsadasdsa4234dsae3421/upload.php", filePath, fileName, newContent[0]))
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}

		}
	}
}
