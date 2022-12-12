package manager

import "github.com/jutionck/go-laundry-app-core/usecase"

type UseCaseManager interface {
	CustomerUseCase() usecase.CustomerUseCase
	ProductUseCase() usecase.ProductUseCase
	BillUseCase() usecase.BillUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CustomerUseCase() usecase.CustomerUseCase {
	return usecase.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

func (u *useCaseManager) ProductUseCase() usecase.ProductUseCase {
	return usecase.NewProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) BillUseCase() usecase.BillUseCase {
	return usecase.NewBillUseCase(
		u.repoManager.BillRepo(),
		u.repoManager.CustomerRepo(),
		u.repoManager.ProductRepo(),
	)
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
