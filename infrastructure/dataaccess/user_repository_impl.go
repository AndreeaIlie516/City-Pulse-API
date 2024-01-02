package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"gorm.io/gorm"
	"errors"
)

type GormUserRepository struct {
	Db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
    return &GormUserRepository{Db: db}
}

func (r *GormUserRepository) AllUsers() ([]entities.User, error) {
	var users []entities.User
	result := r.Db.Find(&users)
	return users, result.Error
}

func (r *GormUserRepository) AllUserIDs() ([]uint, error) {
	var users []entities.User
	var userIDs []uint

	if err := r.Db.Model(&entities.User{}).Select("ID").Find(&users).Error; err != nil {
		return nil, err
	}

	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}

	return userIDs, nil
}

func (r *GormUserRepository) UserByID(id uint) (*entities.User, error) {
	var user entities.User
	if err:=r.Db.First(&user, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *GormUserRepository) CreateUser(user entities.User) (entities.User, error) {
	if err := r.Db.Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *GormUserRepository) DeleteUser(id uint) (entities.User, error) {
	var user entities.User

    if err := r.Db.First(&user, "ID = ?", id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return entities.User{}, errors.New("user not found")
        }
        return entities.User{}, err
    }

    if err := r.Db.Delete(&user).Error; err != nil {
        return entities.User{}, err
    }

    return user, nil
}

func (r *GormUserRepository) UpdateUser(id uint, updatedUser entities.User) (entities.User, error) {
	var user entities.User

    if err := r.Db.First(&user, "ID = ?", id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return entities.User{}, errors.New("user not found")
        }
        return entities.User{}, err
    }

    if err := r.Db.Model(&user).Updates(updatedUser).Error; err != nil {
        return entities.User{}, err
    }

    return user, nil
}
