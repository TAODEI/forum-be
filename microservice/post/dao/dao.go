package dao

import (
	"errors"
	pb "forum-post/proto"
	"forum/model"
	"forum/pkg/constvar"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	dao *Dao
)

// Dao .
type Dao struct {
	DB    *gorm.DB
	Redis *redis.Client
}

// Interface dao
type Interface interface {
	GetItem(Item) (GetDeleter, error)

	CreatePost(*PostModel) (uint32, error)
	ListPost(*PostModel, uint32, uint32, uint32, bool) ([]*PostInfo, error)
	GetPostInfo(uint32) (*PostInfo, error)
	GetPost(uint32) (*PostModel, error)
	IsUserFavoritePost(uint32, uint32) (bool, error)

	CreateComment(*CommentModel) (uint32, error)
	GetCommentInfo(uint32) (*CommentInfo, error)
	GetComment(uint32) (*CommentModel, error)
	ListCommentByPostId(uint32) ([]*pb.CommentInfo, error)
	GetCommentNumByPostId(uint32) uint32

	AddLike(uint32, Item) error
	RemoveLike(uint32, Item) error
	GetLikedNum(Item) (int64, error)
	IsUserHadLike(uint32, Item) (bool, error)
	ListUserLike(uint32) ([]*Item, error)

	CreatePost2Tag(Post2TagModel) error
	GetTagById(uint32) (*TagModel, error)
	GetTagByContent(string) (*TagModel, error)
	ListTagsByPostId(uint32) ([]string, error)
}

// Init init dao
func Init() {
	if dao != nil {
		return
	}

	// init db
	model.DB.Init()

	// init redis
	model.RedisDB.Init()

	// init casbin
	model.CB.Init()

	dao = &Dao{
		DB:    model.DB.Self,
		Redis: model.RedisDB.Self,
	}
}

func GetDao() *Dao {
	return dao
}

type GetDeleter interface {
	Get(uint32) error
	Delete() error
}

func (Dao) GetItem(i Item) (GetDeleter, error) {
	if i.TypeName == constvar.Post {
		item := &PostModel{}
		err := item.Get(i.Id)
		return item, err
	} else if i.TypeName == constvar.Comment {
		item := &CommentModel{}
		err := item.Get(i.Id)
		return item, err
	} else {
		return nil, errors.New("wrong TypeName")
	}
}
