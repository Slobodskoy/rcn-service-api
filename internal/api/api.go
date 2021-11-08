package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozonmp/rcn-service-api/internal/model"
	"github.com/ozonmp/rcn-service-api/internal/repo"

	pb "github.com/ozonmp/rcn-service-api/pkg/rcn-service-api"
)

type serviceAPI struct {
	pb.UnimplementedRcnServiceApiServiceServer
	repo repo.Repo
}

// NewServiceAPI returns api of rcn-service-api service
func NewServiceAPI(r repo.Repo) pb.RcnServiceApiServiceServer {
	return &serviceAPI{repo: r}
}

func (o *serviceAPI) CreateServiceV1(
	ctx context.Context,
	req *pb.CreateServiceV1Request,
) (*pb.CreateServiceV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateServiceV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	service := model.Service{
		Id:          req.GetValue().Id,
		Title:       req.GetValue().Title,
		Description: req.GetValue().Description,
		Rating:      (int)(req.GetValue().Rating),
	}
	serviceId, err := o.repo.CreateService(ctx, service)
	if err != nil {
		log.Error().Err(err).Msg("CreateServiceV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("CreateServiceV1 - success")

	return &pb.CreateServiceV1Response{
		ServiceId: serviceId,
	}, nil
}

func (o *serviceAPI) DescribeServiceV1(
	ctx context.Context,
	req *pb.DescribeServiceV1Request,
) (*pb.DescribeServiceV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeServiceV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	service, err := o.repo.DescribeService(ctx, req.ServiceId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeServiceV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("DescribeServiceV1 - success")

	return &pb.DescribeServiceV1Response{
		Value: service.ToPb(),
	}, nil
}

func (o *serviceAPI) ListServiceV1(
	ctx context.Context,
	req *pb.ListServiceV1Request,
) (*pb.ListServiceV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("ListServiceV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	listService, err := o.repo.ListService(ctx)
	if err != nil {
		log.Error().Err(err).Msg("ListServiceV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("ListServiceV1 - success")

	return &pb.ListServiceV1Response{
		Items: (model.ServiceList)(listService).ToPb(),
	}, nil
}

func (o *serviceAPI) RemoveServiceV1(
	ctx context.Context,
	req *pb.RemoveServiceV1Request,
) (*pb.RemoveServiceV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemoveServiceV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ok, err := o.repo.RemoveService(ctx, req.ServiceId)
	if err != nil {
		log.Error().Err(err).Msg("RemoveServiceV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("RemoveServiceV1 - success")

	return &pb.RemoveServiceV1Response{
		Found: ok,
	}, nil
}
