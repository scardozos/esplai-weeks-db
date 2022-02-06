package grpcdates

import (
	"context"
	"fmt"
	"log"
	"sync"

	weeks_pb "github.com/scardozos/esplai-weeks-db/api/weeksdb"
	local "github.com/scardozos/esplai-weeks-db/cmd/jsondates"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DatesServer handles all gRPC logic to local business logic
type DatesServer struct {
	weeks_pb.UnimplementedWeeksDatabaseServer
	WeeksList    local.WeeksJsonList // WeekJsonModel
	Mu           sync.RWMutex        // Read and write Mutex for safe concurrent reading
	JsonFilename string              // Database JSON name
}

// `GetStaticWeeks` RPC takes an empty request (`GetStaticWeeksRequest`), and returns
// a list of static weeks present in the database. (`GetStaticWeeksResponse`)
func (s *DatesServer) GetStaticWeeks(ctx context.Context, req *weeks_pb.GetStaticWeeksRequest) (*weeks_pb.GetStaticWeeksResponse, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	// Returns all weeks found in database in proto format
	log.Printf("Called GetStaticWeeks")

	return &weeks_pb.GetStaticWeeksResponse{
		StaticWeeks: s.WeeksList.ToGrpcWeeksList(), // []*pb.Date
	}, nil
}

// `IsWeekStatic` RPC takes a static week in the request parameter (`IsWeekStaticRequest`), and
// returns a bool depending on whether the week is present in the static week database
// and the requested week itself (`IsWeekStaticResponse`).
func (s *DatesServer) IsWeekStatic(ctx context.Context, req *weeks_pb.IsWeekStaticRequest) (*weeks_pb.IsWeekStaticResponse, error) {
	s.Mu.RLock()
	defer s.Mu.RUnlock()

	// as wl.IsInList returns the index (int) of the week if present, and
	// if it is present (bool), we don't want to return the index in this case
	_, isPresent := s.WeeksList.IsInList(ProtoToWeekJsonModel(req.Week))

	// return whether the requested week (Week `Date` ) is static or not (IsStatic `Date`),
	// as well as the requested week (RequestedWeek `Date`)
	return &weeks_pb.IsWeekStaticResponse{
		IsStatic:      isPresent, // bool
		RequestedWeek: req.Week,  // pb.Date
	}, nil
}

// `SetStaticWeek` RPC takes a static weej in the request parameter (`SetStaticWeekRequest`), and
// returns the requested week as a response (`SetStaticWeekResponse`).
func (s *DatesServer) SetStaticWeek(ctx context.Context, req *weeks_pb.SetStaticWeekRequest) (*weeks_pb.SetStaticWeekResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	err := s.WeeksList.AddToList(ProtoToWeekJsonModel(req.StaticWeek))
	// if error adding week to weeks list (in memory), return grpc error with Internal status code
	if err != nil {
		if err == local.ErrorDateAlreadyPresent {
			return nil, status.Errorf(codes.FailedPrecondition, "could not add week '%v' to list:", ProtoWeekToString(req.StaticWeek), err)
		}
		return nil, status.Errorf(codes.Internal, "could not add week '%v' to list: %v", ProtoWeekToString(req.StaticWeek), err)
	}

	err = s.WeeksList.ToJsonFile(s.JsonFilename)
	// if error persisting weeks list (to json file), return grpc error with Internal status code
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not add week '%v' to database: %v", ProtoWeekToString(req.StaticWeek), err)
	}

	// if no error, return set Static Week to grpc client
	return &weeks_pb.SetStaticWeekResponse{
		SetWeek: req.StaticWeek, // pb.Date
	}, nil
}

// `UnsetStaticWeek` RPC removes a static week from the database, as changes may occur throughout
// the year.
func (s *DatesServer) UnsetStaticWeek(ctx context.Context, req *weeks_pb.UnsetStaticWeekRequest) (*weeks_pb.UnsetStaticWeekResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	wl := &s.WeeksList // wl is shorthand for the DatesServer weeksList

	err := wl.RemoveFromList(ProtoToWeekJsonModel(req.StaticWeek))
	// if error is encountered when removing week from week list (in mem), return invalid argument
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "week '%v' could not be deleted: %v", ProtoWeekToString(req.StaticWeek), err)
	}

	err = wl.ToJsonFile(s.JsonFilename)
	// if error is encountered when persisting changes to json file, return internal grpc status code
	if err != nil {
		return nil, status.Errorf(codes.Internal, "week '%v' could not be deleted from database: %v", ProtoWeekToString(req.StaticWeek), err)
	}

	// if no error, return unset static week to grpc client
	return &weeks_pb.UnsetStaticWeekResponse{
		UnsetWeek: req.StaticWeek, // pb.Date
	}, nil

}

// Returns local Week Json Model from proto `Date` messages
// to allow for local business logic
func ProtoToWeekJsonModel(proto *weeks_pb.Date) local.WeekJsonModel {
	dateJsonModel := &local.DateJsonModel{Year: proto.Year, Month: proto.Month, Day: proto.Day}
	return *dateJsonModel.ToWeekJsonModel()
}

// Returns string from proto `Date` message for logging
// purposes
func ProtoWeekToString(proto *weeks_pb.Date) string {
	return fmt.Sprintf("%02d-%02d-%d", proto.Day, proto.Month, proto.Year)
}
