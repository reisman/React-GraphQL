package data

import (
	"time"
)

var message1 = &Message{ID: "1", Author: joe, PublishedAt: time.Now(), Text: "TextA"}
var message2 = &Message{ID: "2", Author: sam, PublishedAt: time.Now(), Text: "TextB"}
var message3 = &Message{ID: "3", Author: sam, PublishedAt: time.Now(), Text: "TextC"}
var message4 = &Message{ID: "4", Author: tom, PublishedAt: time.Now(), Text: "TextE"}
var message5 = &Message{ID: "5", Author: joe, PublishedAt: time.Now(), Text: "TextF"}

var messages = map[string]*Message{
	"1": message1,
	"2": message2,
	"3": message3,
	"4": message4,
	"5": message5,
}

func getMessage(id string) *Message {
	if msg, err := messages[id]; err {
		return msg
	}
	return nil
}

func getMessages() []*Message {
	values := []*Message{}
	for _, msg := range messages {
		values = append(values, msg)
	}
	return values
}

func addMessage(msg *Message) {
	messages[msg.ID] = msg
}
