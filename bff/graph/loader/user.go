package loader

import (
	"context"
	"fmt"
	"go-micro-sample/bff/client"
	"go-micro-sample/bff/graph/model"
	"log"
	"strconv"

	"github.com/graph-gophers/dataloader"
)

type UserLoader struct {
	userService client.IUserService
}

func (u *UserLoader) BatchGetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	userIDs := make([]int64, len(keys))
	for i, key := range keys {
		convertedNum, _ := strconv.ParseInt(key.String(), 10, 64)
		userIDs[i] = convertedNum
	}

	log.Printf("BatchGetUsers(id = %v)\n", userIDs)
	userByID, err := u.userService.GetUsers(ctx, userIDs)
	if err != nil {
		err := fmt.Errorf("fail get users, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	output := make([]*dataloader.Result, len(userIDs))
	for i, userKey := range userIDs {
		user, ok := userByID[userKey]
		if ok {
			output[i] = &dataloader.Result{Data: user, Error: nil}
		} else {
			err := fmt.Errorf("user not found %d", userKey)
			output[i] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

func LoadUser(ctx context.Context, userID int64) (*model.User, error) {
	log.Printf("LoadUser(id = %d)\n", userID)
	loaders := GetLoaders(ctx)
	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(fmt.Sprintf("%d", userID)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	user := result.(*model.User)
	log.Printf("return LoadUser(id = %d, name = %s)\n", user.ID, user.Name)
	return user, nil
}
