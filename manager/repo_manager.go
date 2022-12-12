package manager

import "github.com/jutionck/go-laundry-app-core/repository"

type RepositoryManager interface {
	CustomerRepo() repository.CustomerRepository
	ProductRepo() repository.ProductRepository
	BillRepo() repository.BillRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.SqlDb())
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.SqlDb())
}

func (r *repositoryManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra InfraManager) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
