package irepository

import entities "backend/internal/domain/entities"

type IUserRepository interface {

	// 創建使用者
	CreateUser(entities.User) (entities.User, error)
	// 刪除使用者
	DeleteUser(entities.User) (entities.User, error)
	// 更新使用者
	UpdateUser(entities.User) (entities.User, error)

	// 列出該使用者的權限
	ListAuthority(entities.User) ([]entities.Authority, error)
	// 列出該使用者的朋友
	ListFriend(entities.User) ([]entities.Friend, error)
	// 列出該使用者的頻道
	ListChannel(entities.User) ([]entities.Channel, error)
	// 列出該使用者的貼文
	ListPost(entities.User) ([]entities.Post, error)

	// 更新密碼
	UpdatePassword(entities.User, string) error
}

type IUserWithAuthorityRepository interface {
	// 新增權限
	CreateAuthorityByManager(entities.UserWithAuthority, entities.Authority) (entities.UserWithAuthority, error)
	// 刪除權限
	DeleteAuthorityByManager(entities.UserWithAuthority, entities.Authority) (entities.UserWithAuthority, error)
}

type IUserWithFriendRepository interface {
	// 新增朋友
	CreateFriendBySelf(entities.UserWithFriend, entities.Friend) (entities.UserWithFriend, error)
	// 刪除朋友
	DeleteFriendBySelf(entities.UserWithFriend, entities.Friend) (entities.UserWithFriend, error)
	// 更新朋友狀態
	UpdateFriendInitiative(entities.UserWithFriend, entities.Friend) (entities.UserWithFriend, error)
}

type IUserWithChannelRepository interface {
	// 新增頻道
	CreateChannel(entities.UserWithChannel, entities.Channel) (entities.UserWithChannel, error)
	// 刪除
	DeleteChannel(entities.UserWithChannel, entities.Channel) (entities.UserWithChannel, error)
	// 加入公開頻道
	JoinPublicChannel(entities.UserWithChannel, string) (entities.UserWithChannel, error)
	// 退出頻道
	QuitChannel(entities.UserWithChannel, string) (entities.UserWithChannel, error)
}

type IUserWithPostRepository interface {
	// 新增貼文
	CreatePost(entities.UserWithPost, entities.Post) (entities.UserWithPost, error)
	// 刪除貼文
	DeletePost(entities.UserWithPost, string) (entities.UserWithPost, error)
	// 更新貼文
	UpdatePost(entities.UserWithPost, string, string, []string) (entities.UserWithPost, error)
}
