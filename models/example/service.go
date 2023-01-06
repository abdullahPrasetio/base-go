package example

/********************************************************************************
* Temancode Example Service Package                                             *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

type service struct {
	repository Repository
}
type Service interface {
	Create(input RequestEmployee) (Employee, error)
	GetAll() ([]Employee, error)
}

func NewService(repo Repository) *service {
	return &service{repository: repo}
}

func (s *service) Create(input RequestEmployee) (Employee, error) {
	var err error
	var result Employee
	result, err = s.repository.Create(input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *service) GetAll() ([]Employee, error) {
	var err error
	var results []Employee
	results, err = s.repository.GetAll()
	if err != nil {
		return results, err
	}
	return results, nil
}
