package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var deletedMessageInfo = &bot.DeleteMessageParams{}

type TgBot struct {
	bot *bot.Bot
}

func TgBotInit() []bot.Option {
	return []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
	}

}

// обычная отправка сообщений
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Я не хочу с тобой разговаривать",
	})
}

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	// answering callback query first to let Telegram know that we received the callback query,
	// and we're handling it. Otherwise, Telegram might retry sending the update repetitively
	// as it thinks the callback query doesn't reach to our application. learn more by
	// reading the footnote of the https://core.telegram.org/bots/api#callbackquery type.
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   "You selected the button: " + update.CallbackQuery.Data,
	})
	b.DeleteMessage(ctx, deletedMessageInfo)
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Button 1", CallbackData: "button_1"},
				{Text: "Button 2", CallbackData: "button_2"},
				{Text: "Button 3", CallbackData: "button_3"},
			},
		},
	}

	mess, _ := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Click by button!",
		ReplyMarkup: kb,
	})

	deletedMessageInfo = &bot.DeleteMessageParams{mess.Chat.ID, mess.ID}

}

func NewTgBot(tg_api_key string) (*TgBot, error) {
	opts := TgBotInit()

	b, err := bot.New(tg_api_key, opts...)
	if err != nil {
		return nil, err
	}
	return &TgBot{
		bot: b,
	}, nil
}

func (b *TgBot) Start(ctx context.Context) {
	b.bot.Start(ctx)

}
