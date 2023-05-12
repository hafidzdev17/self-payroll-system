package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"self-payrol/model"
	"self-payrol/request"
)

type userUsecase struct {
	userRepository model.UserRepository
	positionRepo   model.PositionRepository
	companyRepo    model.CompanyRepository
}

func NewUserUsecase(user model.UserRepository, post model.PositionRepository, company model.CompanyRepository) model.UserUsecase {
	return &userUsecase{userRepository: user, positionRepo: post, companyRepo: company}
}

func (p *userUsecase) WithdrawSalary(ctx context.Context, req *request.WithdrawRequest) error {
	user, err := p.userRepository.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if user.SecretID != req.SecretID {
		return errors.New("secret id not valid")
	}

	notes := user.Name + " withdraw salary "

	err = p.companyRepo.DebitBalance(ctx, user.Position.Salary, notes)
	if err != nil {
		return err
	}

	return nil
}

func (p *userUsecase) GetByID(ctx context.Context, id int) (*model.User, error) {
	user, err := p.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *userUsecase) FetchUser(ctx context.Context, limit, offset int) ([]*model.User, error) {

	users, err := p.userRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (p *userUsecase) DestroyUser(ctx context.Context, id int) error {
	err := p.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *userUsecase) EditUser(ctx context.Context, id int, req *request.UserRequest) (*model.User, error) {
	_, err := p.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user, err := p.userRepository.UpdateByID(ctx, id, &model.User{
		SecretID:   req.SecretID,
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Address:    req.Address,
		PositionID: req.PositionID,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *userUsecase) StoreUser(ctx context.Context, req *request.UserRequest) (*model.User, error) {
	newUser := &model.User{
		SecretID:   req.SecretID,
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Address:    req.Address,
		PositionID: req.PositionID,
	}

	_, err := p.positionRepo.FindByID(ctx, req.PositionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("position id not valid ")
		}

		return nil, err
	}

	user, err := p.userRepository.Create(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return user, nil
}
