package goocr

import (
	"fmt"
	"github.com/fankuanhui/gosseract-win-demo"
	"log"
)

func OcrImage(imagePath string) {
	log.Println(gosseract.Version())
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(imagePath)
	client.SetLanguage("eng")
	client.SetPageSegMode(gosseract.PSM_SINGLE_BLOCK)
	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}
