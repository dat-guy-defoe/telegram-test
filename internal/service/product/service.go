package product

import "fmt"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	if idx < 0 || idx > len(allProducts)-1 {
		return nil, fmt.Errorf("selected product %v does not exist", idx)
	}

	return &allProducts[idx], nil
}
