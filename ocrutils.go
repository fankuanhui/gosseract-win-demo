package gosseract

import (
	"fmt"
	"log"
)

func OcrImage(imagePath string) {
	log.Println(Version())
	client := NewClient()
	defer client.Close()
	client.SetImage(imagePath)
	client.SetLanguage("eng")
	client.SetPageSegMode(PSM_SINGLE_BLOCK)
	text, err := client.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}
