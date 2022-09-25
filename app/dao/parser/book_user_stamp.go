package parser

import (
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

func BookUserStamp(bookUserStamp *daocore.BookUserStamp) model.BookUserStamp {
	return model.BookUserStamp{
		ID:           bookUserStamp.ID,
		BookID:       bookUserStamp.BookID,
		BookSeriesID: bookUserStamp.BookSeriesID,
		PageNum:      bookUserStamp.PageNum,
		X:            bookUserStamp.X,
		Y:            bookUserStamp.Y,
		UserID:       bookUserStamp.UserID,
		StampID:      bookUserStamp.StampID,
	}
}
