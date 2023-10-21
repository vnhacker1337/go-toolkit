package helpers

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getBot(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	bot.Debug = true

	if err != nil {
		panic("Telegram Token error: " + err.Error())
	}

	return bot, err
}

func SendWithDocument(bot *tgbotapi.BotAPI, chatID int64, filepath string) error {

	msg := tgbotapi.NewDocument(chatID, tgbotapi.FilePath(filepath))
	_, err := bot.Send(msg)

	if err != nil {
		fmt.Println("sendMessage:", err.Error())
		return err
	}

	return nil
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, msg string) error {

	_msg := tgbotapi.NewMessage(chatID, msg)
	_msg.ParseMode = "Markdown"
	_, err := bot.Send(_msg)

	if err != nil {
		fmt.Println("sendMessage:", err.Error())
		return err
	}

	return nil
}

func sendBatchMessageFromFile(bot *tgbotapi.BotAPI, chatID int64, filemsg string) {
	chunkFileContent := ChunkFileByPart(filemsg, 5)
	for _, chunk := range chunkFileContent {

		data := ""

		for _, line := range chunk {
			data = data + "\n" + line
		}

		_msg := tgbotapi.NewMessage(chatID, data)
		_msg.ParseMode = "Markdown"
		_, err := bot.Send(_msg)

		time.Sleep(time.Second * 2)
		if err != nil {
			fmt.Println("sendMessage:", err.Error())
		}
	}
}
