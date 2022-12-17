package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	userEntity := entity.User{}

	if err := r.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Where("id =?", id).
		Find(&userEntity).Error; err != nil {
		return entity.User{}, err
	}

	return userEntity, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	userEntity := entity.User{}

	if err := r.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Where("email =?", email).
		Find(&userEntity).Error; err != nil {
		return entity.User{}, err
	}

	return userEntity, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.db.
		WithContext(ctx).
		Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Where("user = ?", user).
		Updates(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	userEntity := entity.User{}

	if err := r.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", id).
		Delete(&userEntity).Error; err != nil {
		return err
	}

	return nil // TODO: replace this
}
