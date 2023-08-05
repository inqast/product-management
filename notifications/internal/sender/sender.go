package sender

import (
	"context"
	"route256/libs/mw/tracing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/opentracing/opentracing-go"
)

type Sender struct {
	bot        *tgbotapi.BotAPI
	chatID     int64
	retryCount int
}

func New(
	token string,
	chatID int64,
	retryCount int,
) (*Sender, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Sender{
		bot:        bot,
		chatID:     chatID,
		retryCount: retryCount,
	}, nil
}

func (s *Sender) Send(ctx context.Context, notification string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "sender/Send")
	defer span.Finish()

	msg := tgbotapi.NewMessage(s.chatID, notification)

	var sendErr error
	for i := 0; i <= s.retryCount; i++ {
		_, sendErr = s.bot.Send(msg)
		if sendErr == nil {
			break
		}
	}

	if sendErr != nil {
		sendErr = tracing.MarkSpanWithError(ctx, sendErr)
	}

	return sendErr
}
