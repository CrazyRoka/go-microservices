package main

import (
	"context"
	"errors"
	pb "github.com/RostyslavToch/go-microservices/shippy-service-consignment/proto/consignment"
	"github.com/RostyslavToch/go-microservices/shippy-service-vessel/proto/vessel"
)

type service struct {
	repository
	vesselClient vessel.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vessel.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	if vesselResponse == nil {
		return errors.New("error fetching vessel, returning nil")
	}

	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err := s.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repository.GetAll()
	if err != nil {
		return err
	}

	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
