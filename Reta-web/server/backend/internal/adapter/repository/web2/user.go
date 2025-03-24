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
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Delete(&userModel).Error; err != nil {
		return entitiesweb2.User{}, err
	}
	deletedUser := userModel.ToDomain()
	return deletedUser, nil

}

// 使用者儲存庫方法: 更新
func (u *UserRepository) UpdateUser(user entitiesweb2.User) (entitiesweb2.User, error) {
	userModel := model.User{}.FromDomain(user)
	if err := u.gorm.Update(&userModel).Error; err != nil {
		return entitiesweb2.User{}, err
	}
	updatedUser := userModel.ToDomain()
	return updatedUser, nil

}

// 使用者附加權限管理儲存庫結構
type UserAuthorityRepository struct {
	gorm *gorm.DB
}

// 新建使用者附加權限管理儲存庫
func NewUserAuthorityRepository(gorm *gorm.DB) irepositoryweb2.IUserWithAuthorityRepository {
	return &UserAuthorityRepository{
		gorm: gorm,
	}
}

// 使用者附加權限管理儲存庫方法: 新增權限
func (ua *UserAuthorityRepository) CreateAuthorityByManager(manager entitiesweb2.UserWithAuthority, authority entitiesweb2.Authority) (entitiesweb2.UserWithAuthority, error) {
	// 檢查管理者存不存在
	var userModel model.User
	if err := ua.gorm.First(&userModel, "uuid = ?", manager.UUID).Error; err != nil {
		return entitiesweb2.UserWithAuthority{}, err
	}

	// 檢查權限存不存在
	var authModel model.Authority
	if err := ua.gorm.Fist(&authModel, "uuid = ?", authority.UUID).Error; err != nil {
		return entitiesweb2.UserWithAuthority{}, err
	}

	//判斷就有該權限
	var exists int64
	if err := ua.gorm.Table("user_authority").Where("user_uuid = ? AND authority_uuid = ?", userModel.UUID, authModel.UUID).Count(&exists).Error; err | (exists > 0) {
		return entitiesweb2.UserWithAuthority{}, err
	}

	// 在原有權限陣列中新增關聯
	if err := ua.gorm.Model(&userModel).Association("Authorities").Append(&authModel); err != nil {
		return entitiesweb2.UserWithAuthority{}, err
	}

	// 查詢更新後的使用者
	var updatedUserModel model.User
	if err := ua.gorm.Preload("Authorities").First(&updatedUserModel, "uuid = ?", userModel.UUID).Error; err != nil {
		return entitiesweb2.UserWithAuthority{}, err
	}

	// 查詢該使用者更新後的所有權限
	var authoritiesModel []model.Authority
	if err := ua.gorm.Model(&updatedUserModel).Association("Authorities").Find(&authoritiesModel); err != nil {
		return entitiesweb2.UserWithAuthority{}, err
	}

	userAuthority := model.UserWithAuthority{
		User:        updatedUserModel,
		Authorities: authoritiesModel,
	}

	return userAuthority.ToDomain(), nil
}

// 使用者附加權限管理儲存庫方法: 刪除權限
func (ua *UserAuthorityRepository) DeleteAuthorityByManager(manager entitiesweb2.UserWithAuthority, authority entitiesweb2.Authority) (entitiesweb2.UserWithAuthority, error) {
	// 檢查管理者存不存在
	var userModel model.User
	if err := ua.gorm.First(&userModel, "uuid = ?", manager.UUID).Error; err != nil {
		return entitiesweb2.UserWithAuthority{}, err
	}

	// 檢查權限存不存在
	var authModel model.Authority
	if err := ua.gorm.Fist(&authModel, "uuid = ?", authority.UUID).Error; err != nil {
		return entitiesweb2.UserWithAuthority{}, err
	}

	// 判斷本來就沒有權限
	var exists int64
	if err := ua.gorm.Table("user_authority").Where("user_uuid = ? AND authority_uuid = ?", userModel.UUID, authModel.UUID).Count(&exists).Error; err|(exists) == 0 {
		return entitiesweb2.UserWithAuthority{}, err
	}

	// 從原有權限陣列中移除關聯
	if err := ua.gorm.Model(&userModel).Association("Authorities").Delete(&authModel); err != nil {
		return entitiesweb2.UserWithAuthority{}, err
	}

	// 查詢該使用者刪除後的所有權限
	var authoritiesModel []model.Authority
	if err := ua.gorm.Table("user_authority").Joins("JOIN authority ON authority.UUID = user_authority.authority_uuid").Where("user_authority.user_uuid = ?", userModel.UUID).FIND(&authoritiesModel).Error; err != nil {
		return entitiesweb2.UserWithAuthority{}, err
	}

	userAuthority := model.UserWithAuthority{
		User:        userModel,
		Authorities: authoritiesModel,
	}

	return userAuthority.ToDomain(), nil

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
func (uf *UserWithFriendRepository) CreateFriendBySelf(self entitiesweb2.UserWithFriend, friend entitiesweb2.Friend) (entitiesweb2.UserWithFriend, error) {

}

// 使用者附加好友關係儲存庫方法: 刪除朋友
func (uf *UserWithFriendRepository) DeleteFriendBySelf(self entitiesweb2.UserWithFriend, friend entitiesweb2.Friend) (entitiesweb2.UserWithFriend, error) {

}

// 使用者附加好友關係儲存庫方法: 更新朋友
func (uf *UserWithFriendRepository) UpdateFriendBySelf(self entitiesweb2.UserWithFriend, friend entitiesweb2.Friend) (entitiesweb2.UserWithFriend, error) {

}
