package telegram

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
)

func EditFile(id string) (string, string) {
	file, err := os.Open("E:\\blog\\blogBackend\\18chatlink.html")
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
	fileName := make([]byte, 6)
	for i := 0; i < 6; i++ {
		fileName[i] = randstr[rand.Intn(len(randstr))]
	}
	ioutil.WriteFile("E:\\blog\\blogBackend\\"+string(fileName)+".html", []byte(newContent), 0644)

	newFileName := string(fileName) + ".html"
	fileNamePath := "E:\\blog\\blogBackend\\"

	return fileNamePath, newFileName
}

func UploadFile(uri string, filePath string, fileName string, id string) string {
	client := resty.New()

	file, err := os.Open(filePath + fileName)
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
