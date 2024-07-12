package main

import (
	"bytes"
	"html/template"
	"log"
	"sync"
	"text/template/parse"
)

type Message struct {
	ClientID string `json:"clientID"`
	Text     string `json:"text"`
}

type Hub struct {
	clients map[*Client]bool

	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client

	message []*Message
}

func NewHub() *Hub {
	return &Hub{
		clients:    map[*Client]bool{},
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
	}
}

func getMessageTemplate(msg *Message) []byte {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("error while parsing templates", err)
	}
	var renderdMessage bytes.Buffer
	err = tmpl.Execute(&renderdMessage, msg)

	if err != nil {
		log.Fatal("temlate execution:", err)
	}

	return renderdMessage.Bytes()
}
