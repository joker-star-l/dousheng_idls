package main

import (
	"context"
	api "github.com/joker-star-l/dousheng_idls/message/kitex_gen/api"
)

// MessageImpl implements the last service interface defined in the IDL.
type MessageImpl struct{}

// LatestMessage implements the MessageImpl interface.
func (s *MessageImpl) LatestMessage(ctx context.Context, userId int64, friendId int64) (resp *api.LatestMessageResponse, err error) {
	// TODO: Your code here...
	return
}
