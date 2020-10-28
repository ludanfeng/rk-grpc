// Copyright (c) 2020 rookie-ninja
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/rookie-ninja/rk-interceptor/context"
	"github.com/rookie-ninja/rk-interceptor/example/proto"
	"github.com/rookie-ninja/rk-interceptor/logging/zap"
	"github.com/rookie-ninja/rk-interceptor/panic"
	"github.com/rookie-ninja/rk-logger"
	"github.com/rookie-ninja/rk-query"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var (
	bytes = []byte(`{
     "level": "info",
     "encoding": "console",
     "outputPaths": ["stdout"],
     "errorOutputPaths": ["stderr"],
     "initialFields": {},
     "encoderConfig": {
       "messageKey": "msg",
       "levelKey": "",
       "nameKey": "",
       "timeKey": "",
       "callerKey": "",
       "stacktraceKey": "",
       "callstackKey": "",
       "errorKey": "",
       "timeEncoder": "iso8601",
       "fileKey": "",
       "levelEncoder": "capital",
       "durationEncoder": "second",
       "callerEncoder": "full",
       "nameEncoder": "full"
     },
    "maxsize": 1,
    "maxage": 7,
    "maxbackups": 3,
    "localtime": true,
    "compress": true
   }`)

	logger, _, _ = rk_logger.NewZapLoggerWithBytes(bytes, rk_logger.JSON)
)

func main() {
	// create listener
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create event factory
	factory := rk_query.NewEventFactory(
		rk_query.WithAppName("my-app"),
		rk_query.WithLogger(logger),
		rk_query.WithFormat(rk_query.RK))

	// create server interceptor
	opt := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			rk_inter_logging.UnaryServerInterceptor(
				rk_inter_logging.WithEventFactory(factory),
				rk_inter_logging.WithLogger(rk_logger.StdoutLogger)),
			rk_inter_panic.UnaryServerInterceptor(rk_inter_panic.PanicToStderr)),
	}

	// create server
	s := grpc.NewServer(opt...)
	proto.RegisterGreeterServer(s, &GreeterServer{})

	// serving
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type GreeterServer struct{}

func (server *GreeterServer) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	event := rk_inter_context.GetEvent(ctx)
	// add fields
	event.AddFields(zap.String("key", "value"))
	// add error
	event.AddErr(errors.New(""))
	// add pair
	event.AddPair("key", "value")
	// set counter
	event.SetCounter("ctr", 1)
	// timer
	event.StartTimer("sleep")
	time.Sleep(1 * time.Second)
	event.EndTimer("sleep")
	// add to metadata
	rk_inter_context.AddToOutgoingMD(ctx, "key", "1", "2")
	// add request id
	rk_inter_context.AddRequestIdToOutgoingMD(ctx)

	// print incoming metadata
	bytes, _ := json.Marshal(rk_inter_context.GetIncomingMD(ctx))
	println(string(bytes))

	rk_inter_context.GetLogger(ctx).Info("this is info message")

	return &proto.HelloResponse{
		Message: "hello",
	}, nil
}