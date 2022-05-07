package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/youpps/ruler/internal/bot"
)

func init() {
	userDirectory, _ := os.UserHomeDir()
	os.Setenv("FILES_DIRECTORY", userDirectory+"\\jupiter\\")
}

func main() {
	fmt.Print("Enter your telegram bot token: ")
	var token string
	fmt.Scanln(&token)

	bot, err := bot.NewBot(token)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	go func() {
		bot.Run(func(msg string) {
			fmt.Println(msg)
		})
	}()
	
	<-ctx.Done()

	fmt.Println("Bot has started destroying...")

	bot.Destroy(func(msg string) {
		fmt.Println(msg)
	})
}
