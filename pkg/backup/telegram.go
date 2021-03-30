package backup

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"mime/multipart"
	"net/http"
	"bytes"
	"github.com/stefanprodan/mgob/pkg/config"
)

func telegramUpload(filename string, plan config.Plan) (string, error) {
	t1 := time.Now()
	
		file, err := os.Open(filename)
	
		if err != nil {
			return "", err
		}
		defer file.Close()
	
		url := "https://api.telegram.org/bot"+plan.Telegram.Token+"/sendDocument?caption="+filename+"&chat_id="+plan.Telegram.Channel
	
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("document", filepath.Base(file.Name()))
	
		if err != nil {
			return "", err
		}
	
		io.Copy(part, file)
		writer.Close()
		request, err := http.NewRequest("POST", url, body)
	
		if err != nil {
			return "", err
		}
	
		request.Header.Add("Content-Type", writer.FormDataContentType())
		client := &http.Client{}
	
		response, err := client.Do(request)
	
		if err != nil {
			return "", err
		}
		defer response.Body.Close()
	
		content, err := ioutil.ReadAll(response.Body)
	
		if err != nil {
			return "", err
		}
	

	t2 := time.Now()
	msg := fmt.Sprintf("Telegram upload finished `%v` -> `%v` Duration: %v",
		filename, "Chat ID: " + plan.Telegram.Channel, t2.Sub(t1))
	return msg, nil
}
