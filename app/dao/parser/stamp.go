package parser

import (
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

func Stamp(stamp *daocore.Stamp) model.Stamp {
	return model.Stamp{
		ID:   stamp.ID,
		Name: stamp.Name,
	}
}
