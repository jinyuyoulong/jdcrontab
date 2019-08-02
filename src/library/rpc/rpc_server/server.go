package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/jinyuyoulong/jdcrontab/src/library/rpc"
	"log"
	"net"
)

type GRPCServer struct {
	code int
	msg string
	data interface{}
}

// 实现protobuf接口
func (g GRPCServer)GRPCResponse(ctx context.Context, r *pb.GRPCRequest)(*pb.GRPCReply, error){
	name := r.Name
	//code:
	//result := getData(requestID)
	//name := result.NameZh
	//id := result.Id

	s := pb.GRPCReply{
		Code: 0,
		Message:   "i am a msg",
		Data:name,
	}
	return &s, nil

}

func main() {
	fmt.Println("rpc 监听 tcp localhost:8000 start...")
	// 建立 tcp 监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	//创建gRPC 服务器，将我们实现的Greeter服务绑定到一个端口
	s := grpc.NewServer()

	pb.RegisterGrpcServiceServer(s, &GRPCServer{})

	reflection.Register(s)

	fmt.Println("rpc 启动服务")

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
