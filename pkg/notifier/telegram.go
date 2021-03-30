package notifier

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/stefanprodan/mgob/pkg/config"

	"encoding/json"

	"bytes"
)

func sendTelegramNotification(subject string, body string, config *config.Telegram) error {

	msg := "<b>Subject:</b>\n\r" + subject + "\r\n\r\n\r" +
		"<i>" + body + "</i>\r\n"

		reqBody := struct {
			ChatID string  `json:"chat_id"`
			Text   string `json:"text"`
			ParseMode string `json:"parse_mode"`
		}{
			config.Channel,
			msg,
			"HTML",
		}
	
		reqBytes, err := json.Marshal(reqBody)
	
		if err != nil {
			return err
		}
	
		resp, err := http.Post(
			"https://api.telegram.org/bot"+config.Token+"/"+"sendMessage",
			"application/json",
			bytes.NewBuffer(reqBytes),
		)
	
		if err != nil {
			return err
		}
	
		defer resp.Body.Close()
	
		if resp.StatusCode != http.StatusOK {
			return errors.New("unexpected status" + resp.Status)
		}
	
		return err

}
