package parser

import (
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

func BookSeries(bookSeries *daocore.BookSery, authorName string) model.BookSeries {
	return model.BookSeries{
		ID:          bookSeries.ID,
		Title:       bookSeries.Title,
		AuthorID:    bookSeries.AuthorID,
		AuthorName:  authorName,
		Description: bookSeries.Description,
	}
}
