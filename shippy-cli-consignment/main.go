package main

import (
	"context"
	"encoding/json"
	"os"

	pb "github.com/RostyslavToch/go-microservices/shippy-service-consignment/proto/consignment"
	"github.com/micro/go-micro"

	"io/ioutil"
	"log"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("shippy.consignment.cli"))
	service.Init()
	client := pb.NewShippingServiceClient("shippy.service.consignment", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[0]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}
	for _, c := range getAll.Consignments {
		log.Println(c)
	}
}
