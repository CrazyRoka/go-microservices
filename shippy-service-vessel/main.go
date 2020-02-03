package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/RostyslavToch/go-microservices/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro"
)

type Repository interface {
	FindAvailable(*vessel.Specification) (*vessel.Vessel, error)
}

type VesselRepository struct {
	vessels []*vessel.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *vessel.Specification) (*vessel.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *vessel.Specification, res *vessel.Response) error {
	vessels, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessels
	return nil
}

func main() {
	vessels := []*vessel.Vessel{
		&vessel.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(micro.Name("shippy.service.vessel"))
	srv.Init()

	vessel.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
