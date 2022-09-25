package parser

import (
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

func Book(book *daocore.Book) model.Book {
	return model.Book{
		ID:           book.ID,
		Title:        book.Title,
		BookSeriesID: book.BookSeriesID,
		TotalPages:   book.TotalPage,
	}
}
