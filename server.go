package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	hw "Grpc-demo/helloworld"
)

type server struct{
	hw.UnimplementedGreeterServer
}

func (s*server)SayHello(ctx context.Context,req *hw.HelloRequest)(*hw.HelloResponse,error){
	return &hw.HelloResponse{
		Message:fmt.Sprintf("hello,%s! age:%v", req.Name,req.Age),

	},nil
}
func main(){
	lis,err:=net.Listen("tcp", ":50051")
	if err!=nil{
		fmt.Printf("Failed to lIsten:%v", err)
		return
	}
	s:=grpc.NewServer()
	hw.RegisterGreeterServer(s,&server{})

	fmt.Println("Server listening on:50051")
	if err:=s.Serve(lis);err!=nil{
		fmt.Printf("failed to server:%v",err)
	}
}

