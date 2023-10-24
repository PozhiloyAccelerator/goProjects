package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	botUrl := "https://api.telegram.org/bot6369916806:AAEg0GREXZCTsLxblQCWUiGJ1YvQjIYgADY" //общий url бота, надо добавлять еще слеш и команду
	for {
		updates, err := getUpdates(botUrl)
		if err != nil {
			log.Println("что-то пошло не так", err.Error())
		}
		for _, update := range updates {
			err = respond{botUrl, update}
		}
		fmt.Println(update)
	}
}

// запрашивает обновления
func getUpdates(botUrl string) ([]Update, error) {
	respons, err := http.Get(botUrl + "/getUpdates")
	if err != nil {
		return nil, err
	}
	defer respons.Body.Close()
	body, err := ioutil.ReadAll(respons.Body) //т.к данные с сервера приходят в байтовом формате, необходимо привести их в читаемый вид
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

// отвечает на обноаления
func respond(botUrl string, update Update) {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
}

/*func sendMessage(chat_id int, text string) {
	//формировали адрес куда будем отправлять данные
	url := "https://api.telegram.org/bot6369916806:AAEg0GREXZCTsLxblQCWUiGJ1YvQjIYgADY/sendMessage?chat_id=" + strconv.Itoa(chat_id) + "&text=" + text
	//отправка запроса на указанный адрес
	http.Get(url)
}

func sendPhoto(chat_id int, photo string) {
	url := "https://api.telegram.org/bot6369916806:AAEg0GREXZCTsLxblQCWUiGJ1YvQjIYgADY/sendPhoto?chat_id=" + strconv.Itoa(chat_id) + "&photo=" + photo
	http.Get(url)
}      */
