package impl_repository

import (
	"context"

	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/dao/parser"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/model"
	"github.com/hackathon-22-digicon-team4/comiQ-server/app/repository"
)

func (r *Repository) CreateUser(ctx context.Context, user model.User) error {
	txn, err := r.db.BeginRWTx(ctx)
	if err != nil {
		return err
	}
	defer txn.Rollback()
	// ユーザーの存在チェック
	_, err = dao.SelectOneUserByID(ctx, txn, user.ID)
	if err == nil {
		return repository.ErrAlreadyExists
	}
	if err := dao.InsertUser(ctx, txn, parser.UserModel{User: user}.ToDao()); err != nil {
		return err
	}
	return txn.Commit()
}

func (r *Repository) FindUserByID(ctx context.Context, id string) (model.User, error) {
	txn, err := r.db.BeginROTx(ctx)
	if err != nil {
		return model.User{}, err
	}
	defer txn.Rollback()
	user, err := dao.SelectOneUserByID(ctx, txn, id)
	if err != nil {
		return model.User{}, err
	}
	return parser.User(&user), nil
}
