// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/go-micro/examples/booking/srv/geo/proto/geo.proto

/*
Package geo is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-micro/examples/booking/srv/geo/proto/geo.proto

It has these top-level messages:
	Request
	Result
*/
package geo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Geo service

type GeoService interface {
	// Finds the hotels contained nearby the current lat/lon.
	Nearby(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error)
}

type geoService struct {
	c           client.Client
	serviceName string
}

func NewGeoService(serviceName string, c client.Client) GeoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "geo"
	}
	return &geoService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *geoService) Nearby(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.serviceName, "Geo.Nearby", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Geo service

type GeoHandler interface {
	// Finds the hotels contained nearby the current lat/lon.
	Nearby(context.Context, *Request, *Result) error
}

func RegisterGeoHandler(s server.Server, hdlr GeoHandler, opts ...server.HandlerOption) {
	type geo interface {
		Nearby(ctx context.Context, in *Request, out *Result) error
	}
	type Geo struct {
		geo
	}
	h := &geoHandler{hdlr}
	s.Handle(s.NewHandler(&Geo{h}, opts...))
}

type geoHandler struct {
	GeoHandler
}

func (h *geoHandler) Nearby(ctx context.Context, in *Request, out *Result) error {
	return h.GeoHandler.Nearby(ctx, in, out)
}
