// Code generated by truss.
// Rerunning truss will overwrite this file.
// DO NOT EDIT!

// Package grpc provides a gRPC client for the Routines service.
package grpc

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/zaquestion/routines"
	"github.com/zaquestion/routines/routines-service/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.RoutinesServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var scraptrelloresetEndpoint endpoint.Endpoint
	{
		scraptrelloresetEndpoint = grpctransport.NewClient(
			conn,
			"routines.Routines",
			"ScrapTrelloReset",
			EncodeGRPCScrapTrelloResetRequest,
			DecodeGRPCScrapTrelloResetResponse,
			pb.ScrapTrelloResetReply{},
			clientOptions...,
		).Endpoint()
	}

	var getroutinesEndpoint endpoint.Endpoint
	{
		getroutinesEndpoint = grpctransport.NewClient(
			conn,
			"routines.Routines",
			"GetRoutines",
			EncodeGRPCGetRoutinesRequest,
			DecodeGRPCGetRoutinesResponse,
			pb.GetRoutinesReply{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		ScrapTrelloResetEndpoint: scraptrelloresetEndpoint,
		GetRoutinesEndpoint:      getroutinesEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCScrapTrelloResetResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC scraptrelloreset reply to a user-domain scraptrelloreset response. Primarily useful in a client.
func DecodeGRPCScrapTrelloResetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ScrapTrelloResetReply)
	return reply, nil
}

// DecodeGRPCGetRoutinesResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getroutines reply to a user-domain getroutines response. Primarily useful in a client.
func DecodeGRPCGetRoutinesResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.GetRoutinesReply)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCScrapTrelloResetRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain scraptrelloreset request to a gRPC scraptrelloreset request. Primarily useful in a client.
func EncodeGRPCScrapTrelloResetRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ScrapTrelloResetRequest)
	return req, nil
}

// EncodeGRPCGetRoutinesRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getroutines request to a gRPC getroutines request. Primarily useful in a client.
func EncodeGRPCGetRoutinesRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetRoutinesRequest)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.RequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
