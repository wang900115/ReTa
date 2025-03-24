package repositoryweb2

import (
	model "backend/internal/adapter/gorm"
	entitiesweb2 "backend/internal/domain/entities/web2"
	irepositoryweb2 "backend/internal/domain/irepository/web2"

	"gorm.io/gorm"
)

// 使用者儲存庫結構
type UserRepository struct {
	gorm *gorm.DB
}

// 新建使用者儲存庫
func NewUserRepository(gorm *gorm.DB) irepositoryweb2.IUserRepository {
	return &UserRepository{gorm: gorm}
}

// 使用者儲存庫方法: 新增
func (u *UserRepository) CreateUser(user entitiesweb2.User) (entitiesweb2.User, error) {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Create(&userModel).Error; err != nil {
		return entitiesweb2.User{}, err
	}
	createdUser := userModel.ToDomain()
	return createdUser, nil
}

// 使用者儲存庫方法: 刪除
func (u *UserRepository) DeleteUser(user entitiesweb2.User) (entitiesweb2.User, error) {
	usermodel := model.User{}.FromDomain(user)
	if err := u.gorm.Delete(&userModel).Error; err != nil {
		return entitiesweb2.User{}, err
	}
	deletedUser := userModel.ToDomain()
	return deletedUser, nil

}

// 使用者儲存庫方法: 更新
func (u *UserRepository) UpdateUser(user entitiesweb2.User) (entitiesweb2.User, error) {
	usermodel := model.User{}.FromDomain(user)
	if err := u.gorm.Update(&userModel).Error; err != nil {
		return entitiesweb2.User{}, err
	}
	updatedUser := userModel.ToDomain()
	return updatedUser, nil

}

// 使用者附加權限管理儲存庫結構
type UserWithAuthorityRepository struct {
	gorm *gorm.DB
}

// 新建使用者附加權限管理儲存庫
func NewUserWithAuthorityRepository(gorm *gorm.DB) irepositoryweb2.IUserWithAuthorityRepository {
	return &UserWithAuthorityRepository{
		gorm: gorm,
	}
}

// 使用者附加權限管理儲存庫方法: 新增權限
func (ua *NewUserWithAuthorityRepository) CreateAuthorityByManager(manager entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error) {

}

// 使用者附加權限管理儲存庫方法: 刪除權限
func (ua *NewUserWithAuthorityRepository) DeleteAuthorityByManager(manager entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error) {

}

// 使用者附加權限管理儲存庫方法: 更新權限
func (ua *NewUserWithAuthorityRepository) UpdateAuthorityByManager(manager entitiesweb2.UserWithAuthority) (entitiesweb2.UserWithAuthority, error) {

}

// 使用者附加好友關係儲存庫結構
type UserWithFriendRepository struct {
	gorm *gorm.DB
}

// 新建使用者附加好友關係儲存庫
func NewUserWihFriendRepository(gorm *gorm.DB) irepositoryweb2.IUserWithFriendRepository {
	return &UserWithFriendRepository{
		gorm: gorm,
	}
}

// 使用者附加好友關係儲存庫方法: 新增朋友
func (uf *UserWithFriendRepository) CreateFriendByUser(uwf entitiesweb2.UserWithFriend) (entitiesweb2.UserWithFriend, error) {

}

// 使用者附加好友關係儲存庫方法: 刪除朋友
func (uf *UserWithFriendRepository) DeleteFriendByUser(uwf entitiesweb2.UserWithFriend) (entitiesweb2.UserWithFriend, error) {

}

// 使用者附加好友關係儲存庫方法: 更新朋友
func (uf *UserWithFriendRepository) UpdateFriendByUser(uwf entitiesweb2.UserWithFriend) (entitiesweb2.UserWithFriend, error) {

}
