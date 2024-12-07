package schedulerpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

var Scheduler_ServiceDesc = grpc.ServiceDesc{ServiceName: "schedulerpb.Scheduler", HandlerType: (*SchedulerServer)(nil), Methods: []grpc.MethodDesc{{MethodName: "Register", Handler: _Scheduler_Register_Handler}}, Streams: []grpc.StreamDesc{}, Metadata: "scheduler.proto"}

const _ = grpc.SupportPackageIsVersion7

type SchedulerClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}
type SchedulerServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}
type UnsafeSchedulerServer interface{ mustEmbedUnimplementedSchedulerServer() }
type schedulerClient struct{ cc grpc.ClientConnInterface }
type UnimplementedSchedulerServer struct{}

func RegisterSchedulerServer(s grpc.ServiceRegistrar, srv SchedulerServer) {
	s.RegisterService(&Scheduler_ServiceDesc, srv)
}
func _Scheduler_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/schedulerpb.Scheduler/Register"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (UnimplementedSchedulerServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}
func (c *schedulerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/schedulerpb.Scheduler/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
