// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.12
// source: api/api_test.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationTestServiceHello = "/api.TestService/Hello"

type TestServiceHTTPServer interface {
	// Hello 发送消息
	Hello(context.Context, *Empty) (*HelloResp, error)
}

func RegisterTestServiceHTTPServer(s *http.Server, srv TestServiceHTTPServer) {
	r := s.Route("/")
	r.POST("api/test/Hello", _TestService_Hello0_HTTP_Handler(srv))
}

func _TestService_Hello0_HTTP_Handler(srv TestServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationTestServiceHello)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Hello(ctx, req.(*Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloResp)
		return ctx.Result(200, reply)
	}
}

type TestServiceHTTPClient interface {
	Hello(ctx context.Context, req *Empty, opts ...http.CallOption) (rsp *HelloResp, err error)
}

type TestServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewTestServiceHTTPClient(client *http.Client) TestServiceHTTPClient {
	return &TestServiceHTTPClientImpl{client}
}

func (c *TestServiceHTTPClientImpl) Hello(ctx context.Context, in *Empty, opts ...http.CallOption) (*HelloResp, error) {
	var out HelloResp
	pattern := "api/test/Hello"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationTestServiceHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
