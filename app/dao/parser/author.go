package parser

import (
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

func Author(author *daocore.Author) model.Author {
	return model.Author{
		ID:   author.ID,
		Name: author.Name,
	}
}
