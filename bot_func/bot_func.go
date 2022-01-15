package bot_func

import (
	"log"
	"net/http"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

type TeleBot struct {
	BotAPI *tgbotapi.BotAPI
}

var TGBot TeleBot

var Category string
var Type string

var Bot_info struct {
	Token       string
	Username    string
	WebhookURL  string
	WebhookPort string
}

func Bot_init() (bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	err := godotenv.Load(".test.env")
	if err != nil {
		log.Fatal(err)
	}
	Bot_info.Token = os.Getenv("BotToken")
	Bot_info.Username = os.Getenv("BotUserName")
	Bot_info.WebhookURL = os.Getenv("WebhookURL")
	Bot_info.WebhookPort = ":" + os.Getenv("PORT")

	bot, err = tgbotapi.NewBotAPI(Bot_info.Token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	TGBot.BotAPI = bot

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert(Bot_info.WebhookURL, nil))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s\n\n", info.LastErrorMessage)
	}
	updates = bot.ListenForWebhook("/")
	go http.ListenAndServe(Bot_info.WebhookPort, nil)

	// log.Printf("%#V\n\n", updates)

	return bot, updates
}

func (t *TeleBot) SendMessage(ChatID int64, m string) {
	msg := tgbotapi.NewMessage(ChatID, m)
	t.BotAPI.Send(msg)
}

func (t *TeleBot) ReplyMessage(ChatID int64, MessageID int, m string) {
	msg := tgbotapi.NewMessage(ChatID, m)
	msg.ReplyToMessageID = MessageID
	t.BotAPI.Send(msg)
}

// func (t *TeleBot) CreateInlineKeyboardMarkup(ChatID int64) {
// 	msg := tgbotapi.NewMessage(ChatID, "Please choose a category")
// 	msg.ReplyMarkup = inline.CategoryKeyboard
// 	t.BotAPI.Send(msg)
// }

// func (t *TeleBot) InlineKeybaordHandler(ChatID int64, MessageID int, CallBackData string) {
// 	log.Println("CallBackData:", CallBackData)
// 	if CallBackData == "cancel" {
// 		edited_text := tgbotapi.NewEditMessageText(
// 			ChatID,
// 			MessageID,
// 			"cancel this operation",
// 		)
// 		t.BotAPI.Send(edited_text)
// 	} else if CallBackData == "back" {
// 		edited_str := "Please choose a category"
// 		edited_text := tgbotapi.NewEditMessageText(
// 			ChatID,
// 			MessageID,
// 			edited_str,
// 		)
// 		edited_text.ReplyMarkup = &inline.CategoryKeyboard
// 		t.BotAPI.Send(edited_text)
// 	} else if inline.Category[CallBackData] {
// 		Category = CallBackData
// 		edited_str := "Category: " + Category + "\n" +
// 			"Please choose the type"
// 		edited_text := tgbotapi.NewEditMessageText(
// 			ChatID,
// 			MessageID,
// 			edited_str,
// 		)
// 		edited_text.ReplyMarkup = &inline.TypeKeyboard
// 		t.BotAPI.Send(edited_text)
// 	} else if inline.Type[CallBackData] {
// 		Type = CallBackData
// 		edited_str := "Category: " + Category + "\n" +
// 			"Type: " + Type + "\n" +
// 			"Please enter a tag"
// 		edited_text := tgbotapi.NewEditMessageText(
// 			ChatID,
// 			MessageID,
// 			edited_str,
// 		)
// 		t.BotAPI.Send(edited_text)
// 	} else {
// 		edited_str := "The instruction is canceled.\nPlease enter again."
// 		edited_text := tgbotapi.NewEditMessageText(
// 			ChatID,
// 			MessageID,
// 			edited_str,
// 		)
// 		t.BotAPI.Send(edited_text)
// 	}
// }
