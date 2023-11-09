package factory

import (
	"github.com/ineBallardin/imersaofullstackfullcycle/codepix/application/usecase"
	"github.com/ineBallardin/imersaofullstackfullcycle/codepix/infrastructure/repository"

	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase
}
