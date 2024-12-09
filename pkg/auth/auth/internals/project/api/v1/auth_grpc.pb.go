package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

var Auth_ServiceDesc = grpc.ServiceDesc{ServiceName: "auth.Auth", HandlerType: (*AuthServer)(nil), Methods: []grpc.MethodDesc{{MethodName: "RegisterUser", Handler: _Auth_RegisterUser_Handler}, {MethodName: "LoginUser", Handler: _Auth_LoginUser_Handler}, {MethodName: "RefreshUserAuth", Handler: _Auth_RefreshUserAuth_Handler}, {MethodName: "VerifyUserAuth", Handler: _Auth_VerifyUserAuth_Handler}}, Streams: []grpc.StreamDesc{}, Metadata: "auth.proto"}

const _ = grpc.SupportPackageIsVersion7

type authClient struct{ cc grpc.ClientConnInterface }
type UnimplementedAuthServer struct{}
type AuthClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	RefreshUserAuth(ctx context.Context, in *RefreshUserAuthRequest, opts ...grpc.CallOption) (*RefreshUserAuthResponse, error)
	VerifyUserAuth(ctx context.Context, in *VerifyUserAuthRequest, opts ...grpc.CallOption) (*VerifyUserAuthResponse, error)
}
type AuthServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	RefreshUserAuth(context.Context, *RefreshUserAuthRequest) (*RefreshUserAuthResponse, error)
	VerifyUserAuth(context.Context, *VerifyUserAuthRequest) (*VerifyUserAuthResponse, error)
}
type UnsafeAuthServer interface{ mustEmbedUnimplementedAuthServer() }

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}
func (c *authClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) RefreshUserAuth(ctx context.Context, in *RefreshUserAuthRequest, opts ...grpc.CallOption) (*RefreshUserAuthResponse, error) {
	out := new(RefreshUserAuthResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/RefreshUserAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) VerifyUserAuth(ctx context.Context, in *VerifyUserAuthRequest, opts ...grpc.CallOption) (*VerifyUserAuthResponse, error) {
	out := new(VerifyUserAuthResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/VerifyUserAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (UnimplementedAuthServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAuthServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAuthServer) RefreshUserAuth(context.Context, *RefreshUserAuthRequest) (*RefreshUserAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshUserAuth not implemented")
}
func (UnimplementedAuthServer) VerifyUserAuth(context.Context, *VerifyUserAuthRequest) (*VerifyUserAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserAuth not implemented")
}
func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}
func _Auth_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/auth.Auth/RegisterUser"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/auth.Auth/LoginUser"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_RefreshUserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshUserAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshUserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/auth.Auth/RefreshUserAuth"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshUserAuth(ctx, req.(*RefreshUserAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_VerifyUserAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyUserAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/auth.Auth/VerifyUserAuth"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyUserAuth(ctx, req.(*VerifyUserAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}
