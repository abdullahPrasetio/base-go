package example

type service struct {
	repository Repository
}
type Service interface {
	Get(input RequestGetEmploye) (Employee, error)
}

func NewService(repo Repository) *service {
	return &service{repository: repo}
}

func (s *service) Get(input RequestGetEmploye) (Employee, error) {
	response := Employee{}

	return response, nil
}
