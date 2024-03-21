package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var FileName string = "18chatlink.html"

func EditFile(id string) string {
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Println("打开文件失败！")
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("读取文件失败")
	}

	r := regexp.MustCompile(`\b\w{25,}\b`)
	newContent := r.ReplaceAllLiteralString(string(data), id)

	const randstr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	NewFileName := make([]byte, 6)
	for i := 0; i < 6; i++ {
		NewFileName[i] = randstr[rand.Intn(len(randstr))]
	}
	ioutil.WriteFile(string(NewFileName)+".html", []byte(newContent), 0644)

	newFileName := string(NewFileName) + ".html"

	return newFileName
}

func UploadFile(uri string, fileName string, id string) string {
	client := resty.New()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	defer file.Close()

	res, err := client.R().SetFile("fileToUpload", fileName).Post(uri)
	if err != nil {
		fmt.Println("上传成功")
	}
	log.Println(res)

	return "企业ID：" + id + "\n" + "链接： " + "https://106.53.179.226:8032/" + fileName
}

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
				fileName := EditFile(newContent[0])
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, UploadFile("https://106.53.179.226:8032/shjkadfcxsadasdsa4234dsae3421/upload.php", fileName, newContent[0]))
				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}

		}
	}
}
