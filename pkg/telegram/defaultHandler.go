package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	roleVerification(ctx, b, update)
	testSendFile(ctx, b, update)
}

func testSendFile(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Test", CallbackData: "button_test"},
				{Text: "Button 2", CallbackData: "button_2"},
				{Text: "Button 3", CallbackData: "button_3"},
			},
		},
	}

	mess, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Click by button! " + update.Message.From.Username,
		ReplyMarkup: kb,
	})

	deletedMessageInfo = &bot.DeleteMessageParams{mess.Chat.ID, mess.ID}
}
