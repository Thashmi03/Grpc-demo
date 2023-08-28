package main

import (
	"context"
	"fmt"
	"log"

	pb "Grpc-demo/helloworld"

	"google.golang.org/grpc"
)

func main(){
	conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err!=nil{
		log.Fatal("failed %v",err)
	}
	defer conn.Close()
	client:=pb.NewGreeterClient(conn)

	name:="thash"
	age:=18
	response,err:=client.SayHello(context.Background(),&pb.HelloRequest{Name:name,Age:int32(age)})

	if err!=nil{
		log.Fatal("failed %v",err)
	}

	fmt.Printf("Response :%s\n",response.Message)
}