package service

import (
	"context"
	"forum-post/dao"
	pb "forum-post/proto"
	logger "forum/log"
	"forum/model"
	"forum/pkg/constvar"
	"forum/pkg/errno"
	"forum/util"
	"go.uber.org/zap"
)

func (s *PostService) CreatePost(_ context.Context, req *pb.CreatePostRequest, _ *pb.Response) error {
	logger.Info("PostService CreatePost")

	if req.MainPostId != 0 {
		post, err := s.Dao.GetPost(req.MainPostId)
		if err != nil {
			return errno.ServerErr(errno.ErrItemNotFound, "main_post cant find")
		}

		if post.MainPostId != 0 {
			return errno.ServerErr(errno.ErrBadRequest, "the main_post_id is not a main_post id")
		}
	}

	data := &dao.PostModel{
		TypeName:     req.TypeName,
		Content:      req.Content,
		Title:        req.Title,
		CreateTime:   util.GetCurrentTime(),
		CategoryId:   req.CategoryId,
		MainPostId:   req.MainPostId,
		Re:           false,
		CreatorId:    req.UserId,
		LastEditTime: util.GetCurrentTime(),
	}

	postId, err := s.Dao.CreatePost(data)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	if err := model.AddPolicy(req.UserId, constvar.Post, postId, constvar.Write); err != nil {
		return errno.ServerErr(errno.ErrCasbin, err.Error())
	}

	if err := model.AddRole(constvar.Post, postId, req.TypeName); err != nil {
		return errno.ServerErr(errno.ErrCasbin, err.Error())
	}

	for _, content := range req.Tags {
		tag, err := s.Dao.GetTagByContent(content)
		if err != nil {
			return errno.ServerErr(errno.ErrDatabase, err.Error())
		}

		item := dao.Post2TagModel{
			PostId: postId,
			TagId:  tag.Id,
		}
		if err := s.Dao.CreatePost2Tag(item); err != nil {
			return errno.ServerErr(errno.ErrDatabase, err.Error())
		}

		if err := s.Dao.AddTagToSortedSet(tag.Id); err != nil {
			logger.Error(err.Error(), zap.Error(errno.ErrRedis))
		}
	}

	return nil
}
