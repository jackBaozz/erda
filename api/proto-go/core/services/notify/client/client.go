// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: notify.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/services/notify/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// NotifyService notify.proto
	NotifyService() pb.NotifyServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		notifyService: pb.NewNotifyServiceClient(cc),
	}
}

type serviceClients struct {
	notifyService pb.NotifyServiceClient
}

func (c *serviceClients) NotifyService() pb.NotifyServiceClient {
	return c.notifyService
}

type notifyServiceWrapper struct {
	client pb.NotifyServiceClient
	opts   []grpc1.CallOption
}

func (s *notifyServiceWrapper) GetAllNotifyTemplates(ctx context.Context, req *pb.GetAllNotifyTemplatesRequest) (*pb.GetAllNotifyTemplatesResponse, error) {
	return s.client.GetAllNotifyTemplates(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyServiceWrapper) GetNotifyTemplate(ctx context.Context, req *pb.GetNotifyTemplateRequest) (*pb.GetNotifyTemplateResponse, error) {
	return s.client.GetNotifyTemplate(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyServiceWrapper) CreateNotify(ctx context.Context, req *pb.CreateNotifyRequest) (*pb.CreateNotifyResponse, error) {
	return s.client.CreateNotify(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyServiceWrapper) DeleteNotify(ctx context.Context, req *pb.DeleteNotifyRequest) (*pb.DeleteNotifyResponse, error) {
	return s.client.DeleteNotify(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyServiceWrapper) UpdateNotify(ctx context.Context, req *pb.UpdateNotifyRequest) (*pb.UpdateNotifyResponse, error) {
	return s.client.UpdateNotify(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyServiceWrapper) GetUserNotifyList(ctx context.Context, req *pb.GetUserNotifyListRequest) (*pb.GetUserNotifyListResponse, error) {
	return s.client.GetUserNotifyList(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyServiceWrapper) NotifyEnable(ctx context.Context, req *pb.NotifyEnableRequest) (*pb.NotifyEnableResponse, error) {
	return s.client.NotifyEnable(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyServiceWrapper) CreateUserDefineNotifyTemplate(ctx context.Context, req *pb.CreateUserDefineNotifyTemplateRequest) (*pb.CreateUserDefineNotifyTemplateResponse, error) {
	return s.client.CreateUserDefineNotifyTemplate(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyServiceWrapper) GetNotifyDetail(ctx context.Context, req *pb.GetNotifyDetailRequest) (*pb.GetNotifyDetailResponse, error) {
	return s.client.GetNotifyDetail(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *notifyServiceWrapper) GetAllGroups(ctx context.Context, req *pb.GetAllGroupsRequest) (*pb.GetAllGroupsResponse, error) {
	return s.client.GetAllGroups(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}