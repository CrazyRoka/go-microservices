package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/RostyslavToch/go-microservices/shippy-service-vessel/proto/vessel"
	micro "github.com/micro/go-micro/v2"
)

const (
	defaultHost = "mongodb://datastore:27017"
)

func createDummyData(repo repository) {
	vessels := []*Vessel{
		{ID: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(context.Background(), v)
	}
}

func main() {
	srv := micro.NewService(micro.Name("shippy.service.vessel"))
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(context.Background(), uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("shippy").Collection("vessel")
	repository := &VesselRepository{vesselCollection}

	createDummyData(repository)

	vessel.RegisterVesselServiceHandler(srv.Server(), &handler{repository})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
