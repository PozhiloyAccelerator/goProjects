package main

type Update struct { //сам апдейт
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"chat"`
}

type RestResponse struct { //ответ на запрос
	Result []Update `json:"result"`
}

type BotMessage struct { //будем отсылать на сервер в ответ на какое-нибудь сообщение
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}
