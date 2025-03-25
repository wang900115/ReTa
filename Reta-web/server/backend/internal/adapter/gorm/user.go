package model

import (
	entitiesweb2 "backend/internal/domain/entities/web2"
	"time"

	"gorm.io/gorm"
)

// 使用者
type User struct {
	UUID        string `gorm:"column:uuid;primary_key;type:char(36)"`
	IsEnable    bool   `gorm:"column:is_enable;not null;default:true;type:tinyint(1)" json:"is_enable"`
	Username    string `gorm:"column:username;not null;unique;type:varchar(128)" json:"username"`
	Password    string `gorm:"column:password;not null;type:varchar(128)" json:"password"`
	Fullname    string `gorm:"column:full_name;not null;type:varchar(100)" json:"full_name"`
	Nickname    string `gorm:"column:nick_name;not null;type:varchar(100)" json:"nick_name"`
	Phone       string `gorm:"column:phone;type:varchar(10)" json:"phone"`
	Email       string `gorm:"column:email;type:varchar(200)" json:"email"`
	Status      string `gorm:"column:status;not null;type:varchar(20)" json:"status"`
	Description string `gorm:"column:description;type:text" json:"description"`

	CreatedAt time.Time      `gorm:"autoCreateTime;column:created_at;not null;type:datetime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (User) TableName() string {
	return "user"
}
func (u User) ToDomain() entitiesweb2.User {
	return entitiesweb2.User{
		UUID:     u.UUID,
		IsEnable: u.IsEnable,
		Username: u.Username,
		Password: u.Password,
		Fullname: u.Fullname,
		Nickname: u.Nickname,
		Phone:    u.Phone,
		Email:    u.Email,

		Status:      u.Status,
		Description: u.Description,
	}
}
func (u User) FromDomain(user entitiesweb2.User) User {
	return User{
		UUID:     user.UUID,
		IsEnable: user.IsEnable,
		Username: user.Username,
		Password: user.Password,
		Fullname: user.Fullname,
		Nickname: user.Nickname,
		Phone:    user.Phone,
		Email:    user.Email,

		Status:      user.Status,
		Description: user.Description,
	}
}

// 使用者加上權限管理
type UserWithAuthority struct {
	User
	Authorities []Authority `gorm:"many2many:user_authority;foreignKey:UUID;joinForeignKey:user_uuid;References:UUID;JoinReferences:authority_uuid" json:"authorities"`
}

func (UserWithAuthority) TableName() string {
	return "user_authority"
}

func (ua UserWithAuthority) ToDomain() entitiesweb2.UserWithAuthority {
	authorities := make([]entitiesweb2.Authority, len(ua.Authorities))
	for i, a := range ua.Authorities {
		authorities[i] = a.ToDomain()
	}
	return entitiesweb2.UserWithAuthority{
		User:        ua.User.ToDomain(),
		Authorities: authorities,
	}
}
func (ua UserWithAuthority) FromDomain(userwithauthority entitiesweb2.UserWithAuthority) UserWithAuthority {
	authorities := make([]Authority, len(userwithauthority.Authorities))
	for i, a := range userwithauthority.Authorities {
		authorities[i] = Authority{}.FromDomain(a)
	}
	return UserWithAuthority{
		User:        User{}.FromDomain(userwithauthority.User),
		Authorities: authorities,
	}
}

// 使用者加上好友關係
type UserWithFriend struct {
	User
	Friends []Friend `gorm:"many2many:user_friend;foreignKey:UUID;joinForeignKey:user_uuid;References:UUID;JoinReferences:friend_uuid" json:"friends"`
}

func (UserWithFriend) TableName() string {
	return "user_friend"
}
func (uf UserWithFriend) ToDomain() entitiesweb2.UserWithFriend {
	friends := make([]entitiesweb2.Friend, len(uf.Friends))
	for i, a := range uf.Friends {
		friends[i] = a.ToDomain()
	}
	return entitiesweb2.UserWithFriend{
		User:    uf.User.ToDomain(),
		Friends: friends,
		Total:   len(uf.Friends),
	}
}
func (uf UserWithFriend) FromDomain(userwithfriend entitiesweb2.UserWithFriend) UserWithFriend {
	friends := make([]Friend, len(userwithfriend.Friends))
	for i, a := range userwithfriend.Friends {
		friends[i] = Friend{}.FromDomain(a)
	}
	return UserWithFriend{
		User:    User{}.FromDomain(userwithfriend.User),
		Friends: friends,
	}
}
