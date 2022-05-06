package main

import (
	"fmt"
	"log"
	"os"

	"github.com/youpps/ruler/internal/bot"
)

func init() {
	userDirectory, _ := os.UserHomeDir()
	os.Setenv("FILES_DIRECTORY", userDirectory+"\\jupiter\\")
}
func main() {
	bot, err := bot.NewBot("5380916893:AAEfXWE_4BmAc7aS3GJFIEHRybr-rbXfCvs")
	if err != nil {
		log.Fatalln(err)
	}

	bot.Run(func(msg string) {
		fmt.Println(msg)
	})
}
