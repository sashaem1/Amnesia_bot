package telegram

import (
	"context"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func roleVerification(ctx context.Context, b *bot.Bot, update *models.Update) {
	mess := "Я вас не знаю"
	users := []User{}
	database.Select(&users, fmt.Sprintf("SELECT * FROM bot.users where user_id = %d", update.Message.From.ID))
	if len(users) != 0 {
		if users[0].Chat_id == 0 {
			UpdateUserChatId(users[0].User_id, update.Message.Chat.ID)
		}
		mess = fmt.Sprintf("Здравствуйте, %s", users[0].Username)
	} else {
		AddUnregiseredUser(update.Message.From.ID, update.Message.From.Username, update.Message.Chat.ID)
		SendMessageForAdmins(ctx, b,
			fmt.Sprintf("Появился новый незарегестрированный пользователь\nНик:%s Имя: %s %s",
				update.Message.From.Username,
				update.Message.From.FirstName,
				update.Message.From.LastName,
			))
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   mess,
	})
}

func UpdateUserChatId(user_id int64, chat_id int64) {
	query := `
	UPDATE bot.users
	SET  chat_id=%d
	WHERE user_id=%d;`

	database.MustExec(fmt.Sprintf(query, chat_id, user_id))
}

func AddUnregiseredUser(user_id int64, username string, chat_id int64) {
	query := `
	INSERT INTO bot.users
	(user_id, "role", username, chat_id)
	VALUES(%d, 'none', '%s', %d);`

	//fmt.Println(fmt.Sprintf(query, user_id, username, chat_id), user_id, username, chat_id)
	database.MustExec(fmt.Sprintf(query, user_id, username, chat_id))

}
