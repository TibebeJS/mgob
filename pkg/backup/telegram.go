package backup

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"net/http"
	"encoding/json"
	"bytes"
	"github.com/stefanprodan/mgob/pkg/config"
)

func telegramUpload(filename string, plan config.Plan) (string, error) {
	t1 := time.Now()
	
		file, err := os.Open(filename)
	
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		url := "https://api.telegram.org/bot"+plan.Token+"/sendDocument?caption="+filename+"&chat_id="+plan.Channel
	
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile('document', filepath.Base(file.Name()))
	
		if err != nil {
			log.Fatal(err)
		}
	
		io.Copy(part, file)
		writer.Close()
		request, err := http.NewRequest("POST", url, body)
	
		if err != nil {
			log.Fatal(err)
		}
	
		request.Header.Add("Content-Type", writer.FormDataContentType())
		client := &http.Client{}
	
		response, err := client.Do(request)
	
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
	
		content, err := ioutil.ReadAll(response.Body)
	
		if err != nil {
			log.Fatal(err)
		}
	

	t2 := time.Now()
	msg := fmt.Sprintf("Telegram upload finished `%v` -> `%v` Duration: %v",
		filename, "Chat ID: " + plan.Channel, t2.Sub(t1))
	return msg, nil
}
