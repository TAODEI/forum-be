package chat

import (
	. "forum-gateway/handler"
	"forum-gateway/pkg/errno"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"

	m "forum/model"
)

type id struct {
	Id string `json:"id"`
}

func GetId(c *gin.Context) {
	userId := c.MustGet("userId")

	u4 := uuid.NewV4().String()

	if err := m.SetStringInRedis(u4, userId, time.Hour*24); err != nil {
		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
	}

	SendResponse(c, nil, &id{Id: u4})
}
