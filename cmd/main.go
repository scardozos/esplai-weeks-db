package main

import (
	"log"
	"net"
	"os"
	"time"

	pb "github.com/scardozos/esplai-weeks-db/api/weeksdb"
	"github.com/scardozos/esplai-weeks-db/pkg/grpcdates"
	"google.golang.org/grpc"
)

var (
	jsonFilename     = "database.json"
	grpcDbServerAddr = os.Getenv("GRPC_ADDR_SELF")
)

func main() {
	startTime := time.Now()
	lis, err := net.Listen("tcp", grpcDbServerAddr)
	if err != nil {
		// if there's an error when attempting to listen,
		// execute fatal error
		log.Fatalf("failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	// register DatesServer with data from our database
	pb.RegisterWeeksDatabaseServer(grpcServer,
		grpcdates.NewDatesDBServer(jsonFilename, startTime),
	)
	grpcServer.Serve(lis)
}
