package repository

import (
	"context"
	"self-payrol/config"
	"self-payrol/model"
)

type userRepository struct {
	Cfg config.Config
}

func NewUserRepository(cfg config.Config) model.UserRepository {
	return &userRepository{Cfg: cfg}
}

func (p *userRepository) FindByID(ctx context.Context, id int) (*model.User, error) {

	// TODO: buat fungsi untuk mencari user berdasarkan ID pada parameter

	user := new(model.User)

	if err := p.Cfg.Database().WithContext(ctx).
		Where("id = ?", id).
		Preload("Position").
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

func (p *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	// TODO: buat fungsi untuk membuat user berdasarkan struct parameter
	if err := p.Cfg.Database().WithContext(ctx).
		Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

func (p *userRepository) UpdateByID(ctx context.Context, id int, user *model.User) (*model.User, error) {
	// TODO: buat fungsi untuk update user berdasarkan struct parameter

	if err := p.Cfg.Database().WithContext(ctx).
		Model(&model.User{ID: id}).
		Updates(user).
		Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil

}

func (p *userRepository) Delete(ctx context.Context, id int) error {

	_, err := p.FindByID(ctx, id)

	if err != nil {
		return err
	}

	res := p.Cfg.Database().WithContext(ctx).
		Delete(&model.User{}, id)
	if res.Error != nil {

		return res.Error
	}
	return nil
}

func (p *userRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.User, error) {
	var data []*model.User

	if err := p.Cfg.Database().WithContext(ctx).Preload("Position").
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
