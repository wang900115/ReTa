package repository

import (
	model "backend/internal/adapter/gorm/model"
	entities "backend/internal/domain/entities"
	irepository "backend/internal/domain/irepository"

	"gorm.io/gorm"
)

// controll user
type UserRepository struct {
	gorm *gorm.DB
}

// new user
func NewUserRepository(gorm *gorm.DB) irepository.IUserRepository {
	return &UserRepository{gorm: gorm}
}

// create user (return created user)
func (u *UserRepository) CreateUser(user entities.User) (entities.User, error) {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Create(&userModel).Error; err != nil {
		return entities.User{}, err
	}
	createdUser := userModel.ToDomain()
	return createdUser, nil
}

// delete user (return deleted user)
func (u *UserRepository) DeleteUser(user entities.User) (entities.User, error) {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Delete(&userModel).Error; err != nil {
		return entities.User{}, err
	}
	deletedUser := userModel.ToDomain()
	return deletedUser, nil

}

// update user (return updated user)
func (u *UserRepository) UpdateUser(user entities.User) (entities.User, error) {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Model(&userModel).Updates(user).Error; err != nil {
		return entities.User{}, err
	}
	updatedUser := userModel.ToDomain()
	return updatedUser, nil
}

// update user's password (return update'error)
func (u *UserRepository) UpdatePassword(user entities.User, hashedPassword string) error {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Model(&userModel).Update("password", hashedPassword).Error; err != nil {
		return err
	}
	return nil
}

// list user's authorities
func (u *UserRepository) ListAuthority(user entities.User) ([]entities.Authority, error) {
	// todo
	return []entities.Authority{}, nil
}

// list user's friends
func (u *UserRepository) ListFriend(user entities.User) ([]entities.Friend, error) {
	// todo
	return []entities.Friend{}, nil
}

// controll user_autority
type UserAuthorityRepository struct {
	gorm *gorm.DB
}

// new user and user's authority
func NewUserAuthorityRepository(gorm *gorm.DB) irepository.IUserWithAuthorityRepository {
	return &UserAuthorityRepository{
		gorm: gorm,
	}
}

// add authority into user_authority relation map (return user and user's authorities)
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

	if err := ua.gorm.Model(&userModel).Association("Authorities").Append(&authModel); err != nil {
		return entities.UserWithAuthority{}, err
	}

	var updatedUserModel model.User
	if err := ua.gorm.Preload("Authorities").First(&updatedUserModel, "uuid = ?", userModel.UUID).Error; err != nil {
		return entities.UserWithAuthority{}, err
	}

	var authoritiesModel []model.Authority
	if err := ua.gorm.Model(&updatedUserModel).Association("Authorities").Find(&authoritiesModel); err != nil {
		return entities.UserWithAuthority{}, err
	}

	userAuthority := model.UserWithAuthority{
		User:        updatedUserModel,
		Authorities: authoritiesModel,
	}

	return userAuthority.ToDomain(), nil
}

// delete authority from user_authority relation map (return user and user's authorities)
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

	if err := ua.gorm.Model(&userModel).Association("Authorities").Delete(&authModel); err != nil {
		return entities.UserWithAuthority{}, err
	}

	var authoritiesModel []model.Authority
	if err := ua.gorm.Table("user_authority").Joins("JOIN authority ON authority.UUID = user_authority.authority_uuid").Where("user_authority.user_uuid = ?", userModel.UUID).Find(&authoritiesModel).Error; err != nil {
		return entities.UserWithAuthority{}, err
	}

	userAuthority := model.UserWithAuthority{
		User:        userModel,
		Authorities: authoritiesModel,
	}

	return userAuthority.ToDomain(), nil
}

// controll user_friend
type UserFriendRepository struct {
	gorm *gorm.DB
}

// new user and user's friend
func NewUserFriendRepository(gorm *gorm.DB) irepository.IUserWithFriendRepository {
	return &UserFriendRepository{
		gorm: gorm,
	}
}

// add friend into user_friend relation map (return user and user's friends)
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

	if err := uf.gorm.Model(&userModel).Association("Friends").Append(&friendModel); err != nil {
		return entities.UserWithFriend{}, err
	}

	var friendsModel []model.Friend
	if err := uf.gorm.Table("user_friend").Joins("JOIN friend ON friend.UUID = user_friend.friend_uuid").Where("user_friend.user_uuid = ?", userModel.UUID).Find(&friendsModel).Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	userFriend := model.UserWithFriend{
		User:    userModel,
		Friends: friendsModel,
	}
	return userFriend.ToDomain(), nil
}

// delete friend from user_friend relation map (return user and user's friends)
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

	if err := uf.gorm.Model(&userModel).Association("Friends").Delete(&friendModel); err != nil {
		return entities.UserWithFriend{}, err
	}

	var friendsModel []model.Friend
	if err := uf.gorm.Table("user_friend").Joins("JOIN friend ON friend.UUID = user_friend.friend_uuid").Where("user_friend.user_uuid = ?", userModel.UUID).Find(&friendsModel).Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	userFriend := model.UserWithFriend{
		User:    userModel,
		Friends: friendsModel,
	}
	return userFriend.ToDomain(), nil
}

// update friend
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
	if err := uf.gorm.Save(&friendModel).Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	var friendsModel []model.Friend
	if err := uf.gorm.Table("user_friend").Joins("JOIN friend ON friend.UUID = user_friend.friend_uuid").Where("user_friend.user_uuid = ?", userModel.UUID).Find(&friendsModel).Error; err != nil {
		return entities.UserWithFriend{}, err
	}

	userFriend := model.UserWithFriend{
		User:    userModel,
		Friends: friendsModel,
	}
	return userFriend.ToDomain(), nil
}
