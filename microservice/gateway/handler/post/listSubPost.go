package post

//
// import (
// 	"context"
// 	. "forum-gateway/handler"
// 	"forum-gateway/service"
// 	"forum-gateway/util"
// 	pb "forum-post/proto"
// 	pbu "forum-user/proto"
// 	"forum/log"
// 	"forum/pkg/constvar"
// 	"forum/pkg/errno"
// 	"strconv"
//
// 	"github.com/gin-gonic/gin"
// 	"go.uber.org/zap"
// )
//
// // ListSubPost ... 获取从帖
// // @Summary list 从贴 api
// // @Description type_name = normal -> 团队外 (type_name暂时均填normal); 根据 main_post_id 获取主贴的从贴list
// // @Tags post
// // @Accept application/json
// // @Produce application/json
// // @Param limit query int false "limit"
// // @Param page query int false "page"
// // @Param last_id query int false "last_id"
// // @Param type_name path string true "type_name"
// // @Param main_post_id path int true "main_post_id"
// // @Param Authorization header string true "token 用户令牌"
// // @Success 200 {object} ListResponse
// // @Router /post/list/{type_name}/{main_post_id} [get]
// func (a *Api) ListSubPost(c *gin.Context) {
// 	log.Info("Post ListMainPost function called.", zap.String("X-Request-Id", util.GetReqID(c)))
//
// 	userId := c.MustGet("userId").(uint32)
//
// 	typeName := c.Param("type_name")
// 	if typeName != "normal" { //TODO
// 		SendError(c, errno.ErrPathParam, nil, "type_name not legal", GetLine())
// 		return
// 	}
// 	ok, err := a.Dao.Enforce(userId, typeName, constvar.Read)
// 	if err != nil {
// 		SendError(c, errno.InternalServerError, nil, err.Error(), GetLine())
// 		return
// 	}
//
// 	if !ok {
// 		SendError(c, errno.ErrPermissionDenied, nil, "权限不足", GetLine())
// 		return
// 	}
//
// 	id, err := strconv.Atoi(c.Param("main_post_id"))
// 	if err != nil {
// 		SendError(c, errno.ErrPathParam, nil, err.Error(), GetLine())
// 		return
// 	}
//
// 	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
// 	if err != nil {
// 		SendError(c, errno.ErrQuery, nil, err.Error(), GetLine())
// 		return
// 	}
//
// 	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
// 	if err != nil {
// 		SendError(c, errno.ErrQuery, nil, err.Error(), GetLine())
// 		return
// 	}
//
// 	lastId, err := strconv.Atoi(c.DefaultQuery("last_id", "0"))
// 	if err != nil {
// 		SendError(c, errno.ErrQuery, nil, err.Error(), GetLine())
// 		return
// 	}
//
// 	listReq := &pb.ListSubPostRequest{
// 		UserId:     userId,
// 		MainPostId: uint32(id),
// 		TypeName:   typeName,
// 		LastId:     uint32(lastId),
// 		Offset:     uint32(page * limit),
// 		Limit:      uint32(limit),
// 		Pagination: page != 0,
// 	}
//
// 	postResp, err := service.PostClient.ListSubPost(c, listReq)
// 	if err != nil {
// 		SendError(c, err, nil, "", GetLine())
// 		return
// 	}
//
// 	var ids []uint32
// 	for _, post := range postResp.List {
// 		ids = append(ids, post.CreatorId)
// 	}
//
// 	getReq := &pbu.GetInfoRequest{
// 		Ids: ids,
// 	}
// 	userResp, err := service.UserClient.GetInfo(context.TODO(), getReq)
// 	if err != nil {
// 		SendError(c, err, nil, "", GetLine())
// 		return
// 	}
//
// 	//TODO
// 	resp := ListResponse{}
// 	resp.Posts = make([]Post, len(userResp.List))
// 	for i, user := range userResp.List {
// 		resp.Posts[i] = Post{
// 			Content:       postResp.List[i].Content,
// 			Title:         postResp.List[i].Title,
// 			LastEditTime:  postResp.List[i].Time,
// 			CategoryId:    postResp.List[i].CategoryId,
// 			CreatorId:     postResp.List[i].CreatorId,
// 			IsLiked:       postResp.List[i].IsLiked,
// 			CreatorName:   user.Name,
// 			CreatorAvatar: user.AvatarUrl,
// 		}
// 	}
//
// 	SendResponse(c, errno.OK, resp)
// }
