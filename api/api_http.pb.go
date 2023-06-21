// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.12
// source: api/api.proto

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

const OperationApiServiceHello = "/api.ApiService/Hello"

type ApiServiceHTTPServer interface {
	// Hello 发送消息
	Hello(context.Context, *Empty) (*HelloResp, error)
}

func RegisterApiServiceHTTPServer(s *http.Server, srv ApiServiceHTTPServer) {
	r := s.Route("/")
	r.POST("api/Hello", _ApiService_Hello0_HTTP_Handler(srv))
}

func _ApiService_Hello0_HTTP_Handler(srv ApiServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiServiceHello)
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

type ApiServiceHTTPClient interface {
	Hello(ctx context.Context, req *Empty, opts ...http.CallOption) (rsp *HelloResp, err error)
}

type ApiServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewApiServiceHTTPClient(client *http.Client) ApiServiceHTTPClient {
	return &ApiServiceHTTPClientImpl{client}
}

func (c *ApiServiceHTTPClientImpl) Hello(ctx context.Context, in *Empty, opts ...http.CallOption) (*HelloResp, error) {
	var out HelloResp
	pattern := "api/Hello"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApiServiceHello))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
