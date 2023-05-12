package usecase

import (
	"context"
	"net/http"
	"self-payrol/model"
)

type transactionUsecase struct {
	transactionRepository model.TransactionRepository
}

func NewTransactionUsecase(transaction model.TransactionRepository) model.TransactionUsecase {
	return &transactionUsecase{transactionRepository: transaction}
}

func (t *transactionUsecase) Fetch(ctx context.Context, limit, offset int) ([]*model.Transaction, int, error) {
	transations, err := t.transactionRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return transations, http.StatusOK, err

}
