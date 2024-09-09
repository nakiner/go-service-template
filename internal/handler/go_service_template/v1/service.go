package v1

import (
	"context"

	pb "github.com/nakiner/go-service-template/pkg/pb/go_service_template/v1"
)

type Service struct {
}

func (s *Service) TestHandler(ctx context.Context, request *pb.TestHandlerRequest) (*pb.TestHandlerResponse, error) {
	return &pb.TestHandlerResponse{}, nil
}

func NewService() *Service {
	return &Service{}
}
