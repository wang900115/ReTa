package model

import (
	entities "backend/internal/domain/entities"
	"time"

	"gorm.io/gorm"
)

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

func (u User) ToDomain() entities.User {
	return entities.User{
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

func (u User) FromDomain(user entities.User) User {
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

type UserWithAuthority struct {
	User
	Authorities []Authority `gorm:"many2many:user_authority;foreignKey:UUID;joinForeignKey:user_uuid;References:UUID;JoinReferences:authority_uuid" json:"authorities"`
}

func (UserWithAuthority) TableName() string {
	return "user_authority"
}

func (ua UserWithAuthority) ToDomain() entities.UserWithAuthority {
	authorities := make([]entities.Authority, len(ua.Authorities))
	for i, a := range ua.Authorities {
		authorities[i] = a.ToDomain()
	}
	return entities.UserWithAuthority{
		User:        ua.User.ToDomain(),
		Authorities: authorities,
	}
}

func (ua UserWithAuthority) FromDomain(userwithauthority entities.UserWithAuthority) UserWithAuthority {
	authorities := make([]Authority, len(userwithauthority.Authorities))
	for i, a := range userwithauthority.Authorities {
		authorities[i] = Authority{}.FromDomain(a)
	}
	return UserWithAuthority{
		User:        User{}.FromDomain(userwithauthority.User),
		Authorities: authorities,
	}
}

type UserWithFriend struct {
	User
	Friends []Friend `gorm:"many2many:user_friend;foreignKey:UUID;joinForeignKey:user_uuid;References:UUID;JoinReferences:friend_uuid" json:"friends"`
}

func (UserWithFriend) TableName() string {
	return "user_friend"
}
func (uf UserWithFriend) ToDomain() entities.UserWithFriend {
	friends := make([]entities.Friend, len(uf.Friends))
	for i, a := range uf.Friends {
		friends[i] = a.ToDomain()
	}
	return entities.UserWithFriend{
		User:    uf.User.ToDomain(),
		Friends: friends,
		Total:   len(uf.Friends),
	}
}
func (uf UserWithFriend) FromDomain(userwithfriend entities.UserWithFriend) UserWithFriend {
	friends := make([]Friend, len(userwithfriend.Friends))
	for i, a := range userwithfriend.Friends {
		friends[i] = Friend{}.FromDomain(a)
	}
	return UserWithFriend{
		User:    User{}.FromDomain(userwithfriend.User),
		Friends: friends,
	}
}

type UserWithChannel struct {
	User
	Channels []Channel `gorm:"many2many:user_channel;foreignKey:UUID;joinForeignKey:user_uuid;Reference:UUID;JoinReferences:channel_uuid" json:"channels"`
}

func (UserWithChannel) TableName() string {
	return "user_channel"
}

func (uc UserWithChannel) ToDomain() entities.UserWithChannel {
	channels := make([]entities.Channel, len(uc.Channels))
	for i, a := range uc.Channels {
		channels[i] = a.ToDomain()
	}
	return entities.UserWithChannel{
		User:     uc.User.ToDomain(),
		Channels: channels,
	}
}

func (uc UserWithChannel) FromDomain(userwithchannel entities.UserWithChannel) UserWithChannel {
	channels := make([]Channel, len(userwithchannel.Channels))
	for i, a := range userwithchannel.Channels {
		channels[i] = Channel{}.FromDomain(a)
	}
	return UserWithChannel{
		User:     User{}.FromDomain(userwithchannel.User),
		Channels: channels,
	}
}

type UserWithPost struct {
	User
	Posts []Post `gorm:"many2many: user_post;foreignKey:UUID;joinForeignKey:user_uuid;Reference:Author;JoinReferences:post_uuid" json:"posts"`
}

func (UserWithPost) TableName() string {
	return "user_post"
}

func (up UserWithPost) ToDomain() entities.UserWithPost {
	posts := make([]entities.Post, len(up.Posts))
	for i, a := range up.Posts {
		posts[i] = a.ToDomain()
	}
	return entities.UserWithPost{
		User:  up.User.ToDomain(),
		Posts: posts,
	}
}

func (up UserWithPost) FromDomain(userwithpost entities.UserWithPost) UserWithPost {
	posts := make([]Post, len(userwithpost.Posts))
	for i, a := range userwithpost.Posts {
		posts[i] = Post{}.FromDomain(a)
	}
	return UserWithPost{
		User:  User{}.FromDomain(userwithpost.User),
		Posts: posts,
	}
}
