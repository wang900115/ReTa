package repository

import (
	model "backend/internal/adapter/gorm/model"
	entities "backend/internal/domain/entities"
	irepository "backend/internal/domain/irepository"

	"gorm.io/gorm"
)

type UserRepository struct {
	gorm *gorm.DB
}

func NewUserRepository(gorm *gorm.DB) irepository.IUserRepository {
	return &UserRepository{gorm: gorm}
}

func (u *UserRepository) CreateUser(user entities.User) (entities.User, error) {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Create(&userModel).Error; err != nil {
		return entities.User{}, err
	}
	createdUser := userModel.ToDomain()
	return createdUser, nil
}

func (u *UserRepository) DeleteUser(user entities.User) (entities.User, error) {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Delete(&userModel).Error; err != nil {
		return entities.User{}, err
	}
	deletedUser := userModel.ToDomain()
	return deletedUser, nil

}

func (u *UserRepository) UpdateUser(user entities.User) (entities.User, error) {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Model(&userModel).Updates(user).Error; err != nil {
		return entities.User{}, err
	}
	updatedUser := userModel.ToDomain()
	return updatedUser, nil
}

func (u *UserRepository) ListAuthority(user entities.User) ([]entities.Authority, error) {
	var authoritiesModel []model.Authority
	if err := u.gorm.Table("user_authority").Joins("JOIN authority ON authority.UUID = user_authority.authority_uuid").Where("user_authority.user_uuid = ?", user.UUID).Find(&authoritiesModel).Error; err != nil {
		return []entities.Authority{}, err
	}
	var authorities []entities.Authority
	for _, authorityModel := range authoritiesModel {
		authorities = append(authorities, authorityModel.ToDomain())
	}
	return authorities, nil
}

func (u *UserRepository) ListFriend(user entities.User) ([]entities.Friend, error) {
	var friendsModel []model.Friend
	if err := u.gorm.Table("user_friend").Joins("JOIN friend ON friend.UUID = user_friend.friend_uuid").Where("user_friend.user_uuid = ?", user.UUID).Find(&friendsModel).Error; err != nil {
		return []entities.Friend{}, nil
	}
	var friends []entities.Friend
	for _, friendModel := range friendsModel {
		friends = append(friends, friendModel.ToDomain())
	}
	return friends, nil
}

func (u *UserRepository) ListChannel(user entities.User) ([]entities.Channel, error) {
	var channelsModel []model.Channel
	if err := u.gorm.Table("user_channel").Joins("JOIN channel ON channel.UUID = user_channel.channel_uuid").Where("user_channel.user_uuid = ?", user.UUID).Find(&channelsModel).Error; err != nil {
		return []entities.Channel{}, nil
	}

	var channels []entities.Channel
	for _, channelModel := range channelsModel {
		channels = append(channels, channelModel.ToDomain())
	}
	return channels, nil
}

func (u *UserRepository) ListPost(user entities.User) ([]entities.Post, error) {
	var postsModel []model.Post
	if err := u.gorm.Table("user_post").Joins("JOIN post ON post.UUID = user_post.post_uuid").Where("user_post.user_uuid = ?", user.UUID).Find(&postsModel).Error; err != nil {
		return []entities.Post{}, nil
	}

	var posts []entities.Post
	for _, postModel := range postsModel {
		posts = append(posts, postModel.ToDomain())
	}
	return posts, nil
}

func (u *UserRepository) UpdatePassword(user entities.User, hashedPassword string) error {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Model(&userModel).Update("password", hashedPassword).Error; err != nil {
		return err
	}
	return nil
}

type UserAuthorityRepository struct {
	gorm *gorm.DB
}

func NewUserAuthorityRepository(gorm *gorm.DB) irepository.IUserWithAuthorityRepository {
	return &UserAuthorityRepository{
		gorm: gorm,
	}
}

func (ua *UserAuthorityRepository) CreateAuthorityByManager(manager entities.UserWithAuthority, authority entities.Authority) (entities.UserWithAuthority, error) {
	var userModel model.User
	if err := ua.gorm.First(&userModel, "uuid = ?", manager.UUID).Error; err != nil {
		return entities.UserWithAuthority{}, err
	}

	var authModel model.Authority
	if err := ua.gorm.First(&authModel, "uuid = ?", authority.UUID).Error; err != nil {
		return entities.UserWithAuthority{}, err
	}

	var exists int64
	if err := ua.gorm.Table("user_authority").Where("user_uuid = ? AND authority_uuid = ?", userModel.UUID, authModel.UUID).Count(&exists).Error; err != nil || (exists > 0) {
		return entities.UserWithAuthority{}, err
	}

	tx := ua.gorm.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&userModel).Association("Authorities").Append(&authModel); err != nil {
		tx.Rollback()
		return entities.UserWithAuthority{}, err
	}

	var authoritiesModel []model.Authority
	if err := tx.Table("user_authority").Joins("JOIN authority ON authority.UUID = user_authority.authority_uuid").Where("user_authority.user_uuid = ?", userModel.UUID).Find(&authoritiesModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithAuthority{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.UserWithAuthority{}, err
	}

	userAuthority := model.UserWithAuthority{
		User:        userModel,
		Authorities: authoritiesModel,
	}

	return userAuthority.ToDomain(), nil
}

func (ua *UserAuthorityRepository) DeleteAuthorityByManager(manager entities.UserWithAuthority, authority entities.Authority) (entities.UserWithAuthority, error) {
	var userModel model.User
	if err := ua.gorm.First(&userModel, "uuid = ?", manager.UUID).Error; err != nil {
		return entities.UserWithAuthority{}, err
	}

	var authModel model.Authority
	if err := ua.gorm.First(&authModel, "uuid = ?", authority.UUID).Error; err != nil {
		return entities.UserWithAuthority{}, err
	}

	var exists int64
	if err := ua.gorm.Table("user_authority").Where("user_uuid = ? AND authority_uuid = ?", userModel.UUID, authModel.UUID).Count(&exists).Error; err != nil || (exists == 0) {
		return entities.UserWithAuthority{}, err
	}

	tx := ua.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&userModel).Association("Authorities").Delete(&authModel); err != nil {
		tx.Rollback()
		return entities.UserWithAuthority{}, err
	}

	var authoritiesModel []model.Authority
	if err := tx.Table("user_authority").Joins("JOIN authority ON authority.UUID = user_authority.authority_uuid").Where("user_authority.user_uuid = ?", userModel.UUID).Find(&authoritiesModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithAuthority{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.UserWithAuthority{}, err
	}
	userAuthority := model.UserWithAuthority{
		User:        userModel,
		Authorities: authoritiesModel,
	}

	return userAuthority.ToDomain(), nil
}

type UserFriendRepository struct {
	gorm *gorm.DB
}

func NewUserFriendRepository(gorm *gorm.DB) irepository.IUserWithFriendRepository {
	return &UserFriendRepository{
		gorm: gorm,
	}
}

func (uf *UserFriendRepository) CreateFriendBySelf(self entities.UserWithFriend, friend entities.Friend) (entities.UserWithFriend, error) {
	var userModel model.User
	if err := uf.gorm.First(&userModel, "uuid = ?", self.UUID).Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	var friendModel model.Friend
	if err := uf.gorm.First(&friendModel, "uuid = ?", friend.UUID).Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	var exists int64
	if err := uf.gorm.Table("user_friend").Where("user_uuid = ? AND friend_uuid = ?", userModel.UUID, friendModel.UUID).Count(&exists).Error; err != nil || (exists > 0) {
		return entities.UserWithFriend{}, err
	}

	tx := uf.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&userModel).Association("Friends").Append(&friendModel); err != nil {
		tx.Rollback()
		return entities.UserWithFriend{}, err
	}

	var friendsModel []model.Friend
	if err := tx.Table("user_friend").Joins("JOIN friend ON friend.UUID = user_friend.friend_uuid").Where("user_friend.user_uuid = ?", userModel.UUID).Find(&friendsModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithFriend{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	userFriend := model.UserWithFriend{
		User:    userModel,
		Friends: friendsModel,
	}
	return userFriend.ToDomain(), nil
}

func (uf *UserFriendRepository) DeleteFriendBySelf(self entities.UserWithFriend, friend entities.Friend) (entities.UserWithFriend, error) {
	var userModel model.User
	if err := uf.gorm.First(&userModel, "uuid = ?", self.UUID).Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	var friendModel model.Friend
	if err := uf.gorm.First(&friendModel, "uuid = ?", friend.UUID).Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	var exists int64
	if err := uf.gorm.Table("user_friend").Where("user_uuid = ? AND friend_uuid = ?", userModel.UUID, friend.UUID).Count(&exists).Error; err != nil || (exists == 0) {
		return entities.UserWithFriend{}, err
	}

	tx := uf.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&userModel).Association("Friends").Delete(&friendModel); err != nil {
		tx.Rollback()
		return entities.UserWithFriend{}, err
	}

	var friendsModel []model.Friend
	if err := tx.Table("user_friend").Joins("JOIN friend ON friend.UUID = user_friend.friend_uuid").Where("user_friend.user_uuid = ?", userModel.UUID).Find(&friendsModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithFriend{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	userFriend := model.UserWithFriend{
		User:    userModel,
		Friends: friendsModel,
	}
	return userFriend.ToDomain(), nil
}

func (uf *UserFriendRepository) UpdateFriendInitiative(self entities.UserWithFriend, friend entities.Friend) (entities.UserWithFriend, error) {
	var userModel model.User
	if err := uf.gorm.First(&userModel, "uuid = ?", self.UUID).Error; err != nil {
		return entities.UserWithFriend{}, err
	}
	var friendModel model.Friend
	if err := uf.gorm.First(&friendModel, "uuid = ?", friend.UUID).Error; err != nil {
		return entities.UserWithFriend{}, err
	}
	var exists int64
	if err := uf.gorm.Table("user_friend").Where("user_uuid = ? AND friend_uuid = ?", userModel.UUID, friend.UUID).Count(&exists).Error; err != nil || (exists == 0) {
		return entities.UserWithFriend{}, err
	}

	friendModel.Username = friend.Username
	friendModel.Fullname = friend.Fullname
	friendModel.Nickname = friend.Nickname
	friendModel.Status = friend.Status
	friendModel.Banned = friend.Banned

	tx := uf.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Save(&friendModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithFriend{}, err
	}

	var friendsModel []model.Friend
	if err := tx.Table("user_friend").Joins("JOIN friend ON friend.UUID = user_friend.friend_uuid").Where("user_friend.user_uuid = ?", userModel.UUID).Find(&friendsModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithFriend{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	userFriend := model.UserWithFriend{
		User:    userModel,
		Friends: friendsModel,
	}
	return userFriend.ToDomain(), nil
}

type UserChannelRepository struct {
	gorm *gorm.DB
}

func NewUserChannelRepository(gorm *gorm.DB) irepository.IUserWithChannelRepository {
	return &UserChannelRepository{gorm: gorm}
}

func (uc *UserChannelRepository) CreateChannel(userwithChannel entities.UserWithChannel, channel entities.Channel) (entities.UserWithChannel, error) {
	var userModel model.User
	if err := uc.gorm.First(&userModel, "uuid = ?", userwithChannel.UUID).Error; err != nil {
		return entities.UserWithChannel{}, err
	}
	// ! error
	channelModel := model.Channel{}.FromDomain(channel)
	if channelModel.Host != userModel.UUID {
		return entities.UserWithChannel{}, nil
	}

	if err := uc.gorm.Create(&channelModel).Error; err != nil {
		return entities.UserWithChannel{}, err
	}

	tx := uc.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&userModel).Association("Channels").Append(&channelModel); err != nil {
		tx.Rollback()
		return entities.UserWithChannel{}, err
	}

	var channelsModel []model.Channel
	if err := tx.Table("user_channel").Joins("JOIN channel ON channel.UUID = user_channel.channel_uuid").Where("user_channel.user_uuid = ?", userwithChannel.UUID).Find(&channelsModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithChannel{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.UserWithChannel{}, err
	}

	userChannel := model.UserWithChannel{
		User:     userModel,
		Channels: channelsModel,
	}
	return userChannel.ToDomain(), nil
}

func (uc *UserChannelRepository) DeleteChannel(userwithChannel entities.UserWithChannel, channel entities.Channel) (entities.UserWithChannel, error) {
	var userModel model.User
	if err := uc.gorm.First(&userModel, "uuid = ?", userwithChannel.UUID).Error; err != nil {
		return entities.UserWithChannel{}, err
	}

	channelModel := model.Channel{}.FromDomain(channel)
	if channelModel.Host != userModel.UUID {
		return entities.UserWithChannel{}, nil
	}

	if err := uc.gorm.Delete(&channelModel).Error; err != nil {
		return entities.UserWithChannel{}, nil
	}

	tx := uc.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Model(&userModel).Association("Channels").Delete(&channelModel); err != nil {
		tx.Rollback()
		return entities.UserWithChannel{}, err
	}

	var channelsModel []model.Channel
	if err := tx.Table("user_channel").Joins("JOIN channel ON channel.UUID = user_channel.channel_uuid").Where("user_channel.user_uuid = ?", userwithChannel.UUID).Find(&channelsModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithChannel{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.UserWithChannel{}, err
	}

	userChannel := model.UserWithChannel{
		User:     userModel,
		Channels: channelsModel,
	}
	return userChannel.ToDomain(), nil

}

func (uc *UserChannelRepository) JoinPublicChannel(userwithChannel entities.UserWithChannel, uuid string) (entities.UserWithChannel, error) {
	var userModel model.User
	if err := uc.gorm.First(&userModel, "uuid = ?", userwithChannel.UUID).Error; err != nil {
		return entities.UserWithChannel{}, err
	}

	var channelModel model.Channel
	if err := uc.gorm.First(&channelModel, "uuid = ?", uuid).Error; err != nil {
		return entities.UserWithChannel{}, err
	}
	// ! error
	if !channelModel.Public {
		return entities.UserWithChannel{}, nil
	}

	tx := uc.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&userModel).Association("Channels").Append(&channelModel); err != nil {
		tx.Rollback()
		return entities.UserWithChannel{}, err
	}

	var channelsModel []model.Channel
	if err := tx.Table("user_channel").Joins("JOIN channel ON channel.UUID = user_channel.channel_uuid").Where("user_channel.user_uuid = ?", userwithChannel.UUID).Find(&channelsModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithChannel{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.UserWithChannel{}, err
	}

	userChannel := model.UserWithChannel{
		User:     userModel,
		Channels: channelsModel,
	}

	return userChannel.ToDomain(), nil
}

func (uc *UserChannelRepository) QuitChannel(userwithChannel entities.UserWithChannel, uuid string) (entities.UserWithChannel, error) {
	var userModel model.User
	if err := uc.gorm.First(&userModel, "uuid = ?", userwithChannel.UUID).Error; err != nil {
		return entities.UserWithChannel{}, err
	}

	var channelModel model.Channel
	if err := uc.gorm.First(&channelModel, "uuid = ?", uuid).Error; err != nil {
		return entities.UserWithChannel{}, err
	}

	tx := uc.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&userModel).Association("Channels").Delete(&channelModel); err != nil {
		tx.Rollback()
		return entities.UserWithChannel{}, err
	}

	var channelsModel []model.Channel
	if err := uc.gorm.Table("user_channel").Joins("JOIN channel ON channel.UUID = user_channel.channel_uuid").Where("user_channel.user_uuid = ?", userwithChannel.UUID).Error; err != nil {
		tx.Rollback()
		return entities.UserWithChannel{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return entities.UserWithChannel{}, err
	}

	userChannel := model.UserWithChannel{
		User:     userModel,
		Channels: channelsModel,
	}

	return userChannel.ToDomain(), nil
}

type UserPostRepository struct {
	gorm *gorm.DB
}

func NewUserPostRepository(gorm *gorm.DB) irepository.IUserWithPostRepository {
	return &UserPostRepository{gorm: gorm}
}

func (up *UserPostRepository) CreatePost(userwithPost entities.UserWithPost, post entities.Post) (entities.UserWithPost, error) {
	var userModel model.User
	if err := up.gorm.First(&userModel, "uuid = ?", userwithPost.UUID).Error; err != nil {
		return entities.UserWithPost{}, err
	}
	// ! error
	postModel := model.Post{}.FromDomain(post)
	if err := up.gorm.Create(&postModel).Error; err != nil {
		return entities.UserWithPost{}, err
	}

	tx := up.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&userModel).Association("Posts").Append(&postModel); err != nil {
		tx.Rollback()
		return entities.UserWithPost{}, err
	}

	var postsModel []model.Post
	if err := tx.Table("user_post").Joins("JOIN post ON post.UUID = user_post.post_uuid").Where("user_post.user_uuid = ?", userwithPost.UUID).Find(&postsModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithPost{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return entities.UserWithPost{}, err
	}

	userPost := model.UserWithPost{
		User:  userModel,
		Posts: postsModel,
	}

	return userPost.ToDomain(), nil
}

func (up *UserPostRepository) DeletePost(userwithPost entities.UserWithPost, uuid string) (entities.UserWithPost, error) {
	var userModel model.User
	if err := up.gorm.First(&userModel, "uuid = ?", userwithPost.UUID).Error; err != nil {
		return entities.UserWithPost{}, err
	}
	var postModel model.Post
	if err := up.gorm.First(&postModel, "uuid = ?", uuid).Error; err != nil {
		return entities.UserWithPost{}, err
	}

	tx := up.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&userModel).Association("Posts").Delete(&postModel); err != nil {
		tx.Rollback()
		return entities.UserWithPost{}, err
	}

	var postsModel []model.Post
	if err := tx.Table("user_post").Joins("JOIN post ON post.UUID = user_post.post_uuid").Where("user_post.user_uuid = ?", userwithPost.UUID).Find(&postsModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithPost{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return entities.UserWithPost{}, err
	}

	userPost := model.UserWithPost{
		User:  userModel,
		Posts: postsModel,
	}

	return userPost.ToDomain(), nil
}

func (up *UserPostRepository) UpdatePost(userwithPost entities.UserWithPost, uuid string, msg string, url []string) (entities.UserWithPost, error) {
	var userModel model.User
	if err := up.gorm.First(&userModel, "uuid = ?", userwithPost.UUID).Error; err != nil {
		return entities.UserWithPost{}, err
	}
	var postModel model.Post
	if err := up.gorm.First(&postModel, "uuid = ?", uuid).Error; err != nil {
		return entities.UserWithPost{}, err
	}

	tx := up.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	postModel.Content = msg
	postModel.MediaURL = url

	if err := tx.Save(&postModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithPost{}, err
	}

	var postsModel []model.Post
	if err := tx.Table("user_post").Joins("JOIN post ON post.UUID = user_post.post_uuid").Where("user_post.user_uuid = ?", userwithPost.UUID).Find(&postsModel).Error; err != nil {
		tx.Rollback()
		return entities.UserWithPost{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return entities.UserWithPost{}, err
	}

	userPost := model.UserWithPost{
		User:  userModel,
		Posts: postsModel,
	}

	return userPost.ToDomain(), nil
}
