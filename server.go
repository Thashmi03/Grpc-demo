package main

import (
	"context"
	"fmt"
	"net"
	"sync"

	hw "Grpc-demo/helloworld"
	pb "Grpc-demo/task"
	tt "Grpc-demo/customer"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type server struct{
	hw.UnimplementedGreeterServer
}
type taskServiceserver struct{
	mu sync.Mutex
	tasks map[string]*pb.Task
	pb.UnimplementedTaskServiceServer
}
type customerserver struct{
	mu sync.Mutex
	detail map[string]*tt.Task
	tt.UnimplementedTaskServiceServer
}
func (s*server)SayHello(ctx context.Context,req *hw.HelloRequest)(*hw.HelloResponse,error){
	return &hw.HelloResponse{
		Message:fmt.Sprintf("hello,%s! age:%v", req.Name,req.Age),

	},nil
}
func (s*taskServiceserver)AddTask(ctx context.Context,req*pb.Task)(*pb.TaskResponse,error){
	s.mu.Lock()
	defer s.mu.Unlock()
	taskID:=generateID()
	req.Id =taskID
	s.tasks[taskID]=req

	return &pb.TaskResponse{
		Id:taskID},nil
}

func(s*taskServiceserver)GetTasks(ctx context.Context,req *pb.Empty)(*pb.TaskList,error){
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks:=make([]*pb.Task,0,len(s.tasks))
	for _,task:=range s.tasks{
		tasks =append(tasks, task)
	}
	return &pb.TaskList{Tasks: tasks},nil
}
func (s*customerserver)InsertCustomer(ctx context.Context,req*tt.Task)(*tt.TaskResponse,error){
	s.mu.Lock()
	defer s.mu.Unlock()
	customerID:=generateID()
	req.Customerid =customerID
	s.detail[customerID]=req

	return &tt.TaskResponse{
		Id:customerID},nil
}
func(s*customerserver)GetCustomer(ctx context.Context,req *tt.Empty)(*tt.TaskList,error){
	s.mu.Lock()
	defer s.mu.Unlock()

	customer:=make([]*tt.Task,0,len(s.detail))
	for _,cd:=range s.detail{
		customer =append(customer, cd)
	}
	return &tt.TaskList{Customer: customer},nil
}


func generateID() string{
	 
	id:=primitive.NewObjectID()
	taskID:=id.Hex()
	
	return taskID
}
func main(){
	lis,err:=net.Listen("tcp", ":50051")
	if err!=nil{
		fmt.Printf("Failed to lIsten:%v", err)
		return
	}
	s:=grpc.NewServer()
	// hw.RegisterGreeterServer(s,&server{})
	// pb.RegisterTaskServiceServer(s,&taskServiceserver{
	// 	tasks: make(map[string]*pb.Task),
	// })
	tt.RegisterTaskServiceServer(s,&customerserver{
		detail: make(map[string]*tt.Task),
	})
	fmt.Println("Server listening on:50051")
	if err:=s.Serve(lis);err!=nil{
		fmt.Printf("failed to server:%v",err)
	}
}

