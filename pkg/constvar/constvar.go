package constvar

const (
	DefaultLimit = 50

	// 角色权限
	Nobody     = 0 // 无权限用户
	Normal     = 1 // 普通用户
	TeamNormal = 3 // 团队成员
	TeamAdmin  = 4 // 团队管理员

	// 权限限制等级
	AuthLevelNobody     = 0 // 无权限用户级别
	AuthLevelNormal     = 1 // 普通用户级别
	AuthLevelAdmin      = 2 // 管理员级别
	AuthLevelSuperAdmin = 4 // 超管级别

	// trashbin redis key
	Trashbin = "trashbin"

	// item.TypeName
	Post    = "post"
	Comment = "comment"

	// casbin
	Write = uint8(1)
	Read  = uint8(2)

	// comment
	FirstLevelComment  = "first-level"
	SecondLevelComment = "second-level"
)
