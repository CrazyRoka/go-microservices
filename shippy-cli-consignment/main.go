package main

import (
	"encoding/json"
	pb "github.com/RostyslavToch/go-microservices/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"io/ioutil"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	json.Unmarshal(data, consignment)
	return consignment, err
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
}