package parser

import (
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
	"github.com/hackathon-22-digicon-team4/comiQ-server/gen/comiq/daocore"
)

func User(user *daocore.User) model.User {
	return model.User{
		ID:       user.ID,
		Password: user.Password,
	}
}

type UserModel struct {
	model.User
}

func (u UserModel) ToDao() []*daocore.User {
	user := []*daocore.User{}
	user = append(user, &daocore.User{
		ID:       u.ID,
		Password: u.Password,
	})
	return user
}
