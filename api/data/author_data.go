package data

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
	return values
}
