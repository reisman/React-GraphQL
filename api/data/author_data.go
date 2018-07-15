package data

import (
	"sort"
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

func getAuthor(id string) *Author {
	if author, err := authors[id]; err {
		return author
	}
	return nil
}

func getAuthors() []*Author {
	values := []*Author{}
	for _, author := range authors {
		values = append(values, author)
	}

	sort.Sort(sortByName(values))
	return values
}

func addAuthor(author *Author) {
	authors[author.ID] = author
}

type sortByName []*Author

func (s sortByName) Len() int {
	return len(s)
}
func (s sortByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}
