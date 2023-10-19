package grpc

import (
	"context"

	"github.com/lavinas/keel/internal/client/adapter/hdlr/grpc/pb"
	"github.com/lavinas/keel/internal/client/core/port"
	"github.com/lavinas/keel/internal/client/core/dto"

)

// ClientGRPCService is a handler for grpc framework
type ClientGRPCService struct {
	Service port.Service
	pb.UnimplementedClientServiceServer
}

// NewClientGRPCService creates a new ClientGRPCService
func NewClientGRPCService(service port.Service) *ClientGRPCService {
	return &ClientGRPCService{
		Service: service,
	}
}

// Insert responds for call of creates a new client
func (c *ClientGRPCService) Insert(ctx context.Context, in *pb.ClientInsert) (*pb.ClientInsertResult, error) {
	var input dto.InsertInputDto
	var output dto.InsertOutputDto
	input.Name = in.Name
	input.Nickname = in.Nickname
	input.Document = in.Document
	input.Phone = in.Phone
	input.Email = in.Email
	if err := c.Service.Insert(&input, &output); err != nil {
		return &pb.ClientInsertResult{Status: err.Error()}, err
	}
	return &pb.ClientInsertResult{
		Status: "ok", 
		Id: output.Id, 
		Name: output.Name,
		Nickname: output.Nickname,
		Document: output.Document,
		Phone: output.Phone,
		Email: output.Email,
		}, nil
}