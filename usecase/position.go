package usecase

import (
	"context"
	"self-payrol/model"
	"self-payrol/request"
)

type positionUsecase struct {
	positionRepository model.PositionRepository
}

func NewPositionUsecase(position model.PositionRepository) model.PositionUsecase {
	return &positionUsecase{positionRepository: position}
}

func (p *positionUsecase) GetByID(ctx context.Context, id int) (*model.Position, error) {
	position, err := p.positionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return position, nil
}

func (p *positionUsecase) FetchPosition(ctx context.Context, limit, offset int) ([]*model.Position, error) {

	positions, err := p.positionRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return positions, nil

}

func (p *positionUsecase) DestroyPosition(ctx context.Context, id int) error {
	err := p.positionRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *positionUsecase) EditPosition(ctx context.Context, id int, req *request.PositionRequest) (*model.Position, error) {
	_, err := p.positionRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	position, err := p.positionRepository.UpdateByID(ctx, id, &model.Position{
		Name:   req.Name,
		Salary: req.Salary,
	})

	if err != nil {
		return nil, err
	}

	return position, nil
}

func (p *positionUsecase) StorePosition(ctx context.Context, req *request.PositionRequest) (*model.Position, error) {
	newPosition := &model.Position{
		Name:   req.Name,
		Salary: req.Salary,
	}

	position, err := p.positionRepository.Create(ctx, newPosition)

	if err != nil {
		return nil, err
	}

	return position, nil
}
