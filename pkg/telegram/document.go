package telegram

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func SendDocumentInChat(ctx context.Context, b *bot.Bot, chatID int64, filepath string) {
	fileData, errReadFile := os.ReadFile(filepath)
	if errReadFile != nil {
		fmt.Printf("error read file, %v\n", errReadFile)
		return
	}

	params := &bot.SendDocumentParams{
		ChatID:   chatID,
		Document: &models.InputFileUpload{Filename: "demo.txt", Data: bytes.NewReader(fileData)},
		Caption:  "Document",
	}

	b.SendDocument(ctx, params)
}
