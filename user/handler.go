package main

import (
	"context"
	api "github.com/joker-star-l/dousheng_idls/user/kitex_gen/api"
)

// UserImpl implements the last service interface defined in the IDL.
type UserImpl struct{}

// UserInfo implements the UserImpl interface.
func (s *UserImpl) UserInfo(ctx context.Context, userId int64) (resp *api.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}
