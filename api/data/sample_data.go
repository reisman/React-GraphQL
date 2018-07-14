package data

import (
	"time"
)

var hans = &Author{ID: "1", Name: "Hans"}
var peter = &Author{ID: "2", Name: "Peter"}
var sam = &Author{ID: "3", Name: "Sam"}
var tom = &Author{ID: "4", Name: "Tom"}
var joe = &Author{ID: "5", Name: "Joe"}

var authors = map[string]*Author{
	"1": hans,
	"2": peter,
	"3": sam,
	"4": tom,
	"5": joe,
}

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

func getAuthor(id string) *Author {
	if author, err := authors[id]; err {
		return author
	}
	return nil
}

func getMessage(id string) *Message {
	if msg, err := messages[id]; err {
		return msg
	}
	return nil
}

func getAuthors() []*Author {
	values := []*Author{}
	for _, author := range authors {
		values = append(values, author)
	}
	return values
}

func getMessages() []*Message {
	values := []*Message{}
	for _, msg := range messages {
		values = append(values, msg)
	}
	return values
}

func toInterfaceSlice(a ...*Author) []interface{} {
	var interfaceSlice = make([]interface{}, len(a))
	for i, d := range a {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
