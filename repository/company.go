package repository

import (
	"context"
	"errors"
	"self-payrol/config"
	"self-payrol/model"

	"gorm.io/gorm"
)

type companyRepository struct {
	Cfg config.Config
}

func NewCompanyRepository(cfg config.Config) model.CompanyRepository {
	return &companyRepository{Cfg: cfg}
}

func (c *companyRepository) Get(ctx context.Context) (*model.Company, error) {
	company := new(model.Company)

	if err := c.Cfg.Database().WithContext(ctx).First(company).Error; err != nil {
		return nil, err
	}

	return company, nil
}

func (c *companyRepository) CreateOrUpdate(ctx context.Context, company *model.Company) (*model.Company, error) {

	companyModel := new(model.Company)

	if err := c.Cfg.Database().WithContext(ctx).Debug().
		First(&companyModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := c.Cfg.Database().WithContext(ctx).Create(&company).Find(companyModel).Error; err != nil {
				return nil, err
			}

			return companyModel, nil
		}
		return nil, err
	}

	// TODO: tuliskan baris code untuk update data company

	return companyModel, nil
}

func (c *companyRepository) DebitBalance(ctx context.Context, amount int, note string) error {
	company, err := c.Get(ctx)
	if err != nil {
		return errors.New("company data not found")
	}

	// TODO: tuliskan baris code untuk mengurangi balance

	if err := c.Cfg.Database().WithContext(ctx).Model(company).Updates(company).Find(company).Error; err != nil {
		return err

	}

	if err := c.Cfg.Database().WithContext(ctx).Create(&model.Transaction{
		Amount: amount,
		Note:   note,
		Type:   model.TransactionTypeDebit,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (c *companyRepository) AddBalance(ctx context.Context, balance int) (*model.Company, error) {
	company, err := c.Get(ctx)
	if err != nil {
		return nil, errors.New("company data not found")
	}

	// TODO: tuliskan baris code untuk topup balance

	if err := c.Cfg.Database().WithContext(ctx).Model(company).Updates(company).Find(company).Error; err != nil {
		return nil, err

	}

	if err := c.Cfg.Database().WithContext(ctx).Create(&model.Transaction{
		Amount: balance,
		Note:   "Topup balance company",
		Type:   model.TransactionsTypeCredit,
	}).Error; err != nil {
		return nil, err
	}

	return company, nil
}
