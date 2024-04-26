// Code generated by protoc-gen-dyn-gin. DO NOT EDIT.
// versions:
//   - protoc-gen-dyn-gin v1.0.0
//   - protoc             v4.25.3
// source: hello.proto

package examples

import (
	context "context"
	errors "errors"
	gin "github.com/gin-gonic/gin"
	http "github.com/things-go/dyn/transport/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible.
var _ = errors.New
var _ = context.TODO
var _ = gin.New
var _ = http.FromCarrier

// GreeterHTTPServer The greeting service definition.
type GreeterHTTPServer interface {
	// SayHello Sends a hello
	// I am a trailing comment
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// GetHello Sends a hello
	GetHello(context.Context, *GetHelloRequest) (*GetHelloReply, error)
}

func RegisterGreeterHTTPServer(g *gin.RouterGroup, srv GreeterHTTPServer) {
	r := g.Group("")
	{
		r.POST("/v1/hello", _Greeter_SayHello0_HTTP_Handler(srv))
		r.GET("/v1/hello/:id", _Greeter_GetHello0_HTTP_Handler(srv))
	}
}

func _Greeter_SayHello0_HTTP_Handler(srv GreeterHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		carrier := http.FromCarrier(c.Request.Context())
		shouldBind := func(req *HelloRequest) error {
			if err := carrier.Bind(c, req); err != nil {
				return err
			}
			return carrier.Validate(c.Request.Context(), req)
		}

		var err error
		var req HelloRequest
		var reply *HelloReply

		if err = shouldBind(&req); err != nil {
			carrier.Error(c, err)
			return
		}
		reply, err = srv.SayHello(c.Request.Context(), &req)
		if err != nil {
			carrier.Error(c, err)
			return
		}
		carrier.Render(c, reply)
	}
}

func _Greeter_GetHello0_HTTP_Handler(srv GreeterHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		carrier := http.FromCarrier(c.Request.Context())
		shouldBind := func(req *GetHelloRequest) error {
			if err := carrier.BindQuery(c, req); err != nil {
				return err
			}
			if err := carrier.BindUri(c, req); err != nil {
				return err
			}
			return carrier.Validate(c.Request.Context(), req)
		}

		var err error
		var req GetHelloRequest
		var reply *GetHelloReply

		if err = shouldBind(&req); err != nil {
			carrier.Error(c, err)
			return
		}
		reply, err = srv.GetHello(c.Request.Context(), &req)
		if err != nil {
			carrier.Error(c, err)
			return
		}
		carrier.Render(c, reply)
	}
}
