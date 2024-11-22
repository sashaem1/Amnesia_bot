package telegram

import (
	"context"
	"log"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       true,
	})

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	params := strings.Split(update.CallbackQuery.Data, ":")
	switch params[0] {
	case "button_test":
		filepath := "./configs/amnezia_for_wireguard.conf"

		SendDocumentInChat(ctx, b, update.CallbackQuery.Message.Message.Chat.ID, filepath)

		b.DeleteMessage(ctx, deletedMessageInfo)
	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "You selected the button: " + update.CallbackQuery.Data,
		})
		b.DeleteMessage(ctx, deletedMessageInfo)
	}

}
