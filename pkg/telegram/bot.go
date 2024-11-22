package telegram

import (
	"context"

	"github.com/go-telegram/bot"
)

var deletedMessageInfo = &bot.DeleteMessageParams{}

type TgBot struct {
	bot *bot.Bot
}

func (b *TgBot) Start(ctx context.Context) {
	b.bot.Start(ctx)
}

func TgBotInit() []bot.Option {
	return []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, callbackHandler),
	}
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

func SendMessageForAdmins(ctx context.Context, b *bot.Bot, mess string) {
	users := []User{}
	database.Select(&users, "SELECT * FROM bot.users where role = 'admin'")

	for _, admin := range users {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: admin.Chat_id,
			Text:   mess,
		})
	}
}
