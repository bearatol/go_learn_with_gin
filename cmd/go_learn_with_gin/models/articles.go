package models

type article struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Date string `json:"date"`
}

var articlesList = []article{
	article{ID: 1, Text: "Test text. Test text. Test text. Test text. Test text. Test text. ", Date: "07.06.2021"},
	article{ID: 2, Text: "Test text. Test text.", Date: "08.06.2021"},
}

func GetAllArticles() []article {
	return articlesList
}
