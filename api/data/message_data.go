package data

import (
	"sort"
	"time"
)

var now = time.Now()
var message1 = &Message{ID: "1", Author: joe, PublishedAt: now.Add(-1 * time.Second), Text: "TextA"}
var message2 = &Message{ID: "2", Author: sam, PublishedAt: now.Add(-2 * time.Second), Text: "TextB"}
var message3 = &Message{ID: "3", Author: sam, PublishedAt: now.Add(-3 * time.Second), Text: "TextC"}
var message4 = &Message{ID: "4", Author: tom, PublishedAt: now.Add(-4 * time.Second), Text: "TextE"}
var message5 = &Message{ID: "5", Author: joe, PublishedAt: now.Add(-5 * time.Second), Text: "TextF"}

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

	sort.Sort(sortByDate(values))
	return values
}

func addMessage(msg *Message) {
	messages[msg.ID] = msg
}

type sortByDate []*Message

func (s sortByDate) Len() int {
	return len(s)
}
func (s sortByDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortByDate) Less(i, j int) bool {
	return s[i].PublishedAt.Before(s[j].PublishedAt)
}
