package model

import pb "github.com/ozonmp/rcn-service-api/pkg/rcn-service-api"

type Service struct {
	Id          uint64
	Title       string
	Description string
	Rating      int
}

type ServiceList []*Service

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type ServiceEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Service
}

func (service *Service) ToPb() *pb.Service {
	return &pb.Service{
		Id:          service.Id,
		Title:       service.Title,
		Description: service.Description,
		Rating:      int32(service.Rating),
	}
}

func (ls ServiceList) ToPb() []*pb.Service {
	var result []*pb.Service
	for _, val := range ls {
		result = append(result, val.ToPb())
	}
	return result
}
