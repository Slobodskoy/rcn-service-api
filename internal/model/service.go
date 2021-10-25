package model

type Service struct {
	Id          uint64
	Title       string
	Description string
	Rating      int
}

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
