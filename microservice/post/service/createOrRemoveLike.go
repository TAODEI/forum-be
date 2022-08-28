package service

import (
	"context"
	"forum-post/dao"
	pb "forum-post/proto"
	logger "forum/log"
	"forum/pkg/errno"
)

func (s *PostService) CreateOrRemoveLike(_ context.Context, req *pb.LikeRequest, _ *pb.Response) error {
	logger.Info("PostService CreateLike")

	item := dao.Item{
		Id:       req.Item.TargetId,
		TypeName: req.Item.TypeName,
	}

	ok, err := s.Dao.IsUserHadLike(req.UserId, item)
	if err != nil {
		return errno.ServerErr(errno.ErrRedis, err.Error())
	}

	if ok {
		err = s.Dao.RemoveLike(req.UserId, item)
	} else {
		err = s.Dao.AddLike(req.UserId, item)
	}
	if err != nil {
		return errno.ServerErr(errno.ErrRedis, err.Error())
	}

	return nil
}