package seeder

import (
	"clean-code-structure/delivery/httpserver"
	"clean-code-structure/param"
	"clean-code-structure/param/productparam"
	"context"
)

type Seeder struct {
	services httpserver.RequiredServices
}

func New(services httpserver.RequiredServices) Seeder {
	return Seeder{
		services: services,
	}
}

func (s Seeder) Seed() {
	s.productSeeder()
}
func (s Seeder) productSeeder() {
	_, err := s.services.ProductService.Store(productparam.StoreRequest{
		BaseRequest: param.BaseRequest{
			Ctx: context.Background(),
		},
		Name: "product-test",
	})
	if err != nil {
		panic(err)
	}
}
