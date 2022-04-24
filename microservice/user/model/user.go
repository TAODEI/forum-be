package model

import (
	m "forum/model"
	"forum/pkg/constvar"
	"github.com/ShiinaOrez/GoSecurity/security"

	"github.com/jinzhu/gorm"
)

type UserModel struct {
	ID           uint32 `json:"id" gorm:"column:id;not null" binding:"required"`
	Name         string `json:"name" gorm:"column:name;" binding:"required"`
	Email        string `json:"email" gorm:"column:email;" binding:"required"`
	Avatar       string `json:"avatar" gorm:"column:avatar;" binding:"required"`
	Role         uint32 `json:"role" gorm:"column:role;" binding:"required"`
	Message      uint32 `json:"message" gorm:"column:message;" binding:"required"`
	PasswordHash string `json:"password_hash" gorm:"column:password_hash;" binding:"required"`
	StudentID    string `json:"student_id" gorm:"column:student_id;"`
}

func (u *UserModel) TableName() string {
	return "users"
}

// Create ... create user
func (u *UserModel) Create() error {
	return m.DB.Self.Create(u).Error
}

// Save ... save user.
func (u *UserModel) Save() error {
	return m.DB.Self.Save(u).Error
}

// GetUser get a single user by id
func GetUser(id uint32) (*UserModel, error) {
	user := &UserModel{}
	d := m.DB.Self.Where("id = ?", id).First(user)
	if d.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, d.Error
}

// GetUserByIds get user by id array
func GetUserByIds(ids []uint32) ([]*UserModel, error) {
	list := make([]*UserModel, 0)
	if err := m.DB.Self.Where("id IN (?)", ids).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

// GetUserByEmail get a user by email.
func GetUserByEmail(email string) (*UserModel, error) {
	u := &UserModel{}
	err := m.DB.Self.Where("email = ?", email).First(u).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return u, err
}

// GetUserByName get a user by name.
func GetUserByName(name string) (*UserModel, error) {
	u := &UserModel{}
	err := m.DB.Self.Where("name = ?", name).First(u).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}
	return u, err
}

// ListUser list users
func ListUser(offset, limit, lastId uint32, filter *UserModel) ([]*UserModel, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	list := make([]*UserModel, 0)

	query := m.DB.Self.Model(&UserModel{}).Where(filter).Offset(offset).Limit(limit)

	if lastId != 0 {
		query = query.Where("id < ?", lastId).Order("id desc")
	}

	if err := query.Scan(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

//
func GeneratePasswordHash(password string) string {
	return security.GeneratePasswordHash(password)
}

func (u *UserModel) CheckPassword(password string) bool {
	return security.CheckPasswordHash(password, u.PasswordHash)
}
