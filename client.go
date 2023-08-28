package main

import (
	"context"
	"fmt"
	"log"

	// hw "Grpc-demo/helloworld"
	// pb "Grpc-demo/task"
	tt "Grpc-demo/customer"
	"google.golang.org/grpc"
)

func main(){
	conn,err:=grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err!=nil{
		log.Fatal("failed1 %v",err)
	}
	defer conn.Close()
	// client1:=hw.NewGreeterClient(conn)
	// client2:=pb.NewTaskServiceClient(conn)
	client2:=tt.NewTaskServiceClient(conn)
	// name:="thash"
	// age:=18

	// task:=&pb.Task{
	// 	Title:"Buy groceries",
	// }
	detail:=&tt.Task{
		Name:"Thash",
		AccountId: 01,
		Balance: 100,
		BankId:100,
	}
	// response,err:=client1.SayHello(context.Background(),&hw.HelloRequest{Name:name,Age:int32(age)})
	// addres,err:=client2.AddTask(context.Background(),task)
	addres,err:=client2.InsertCustomer(context.Background(),detail)
	if err!=nil{
		log.Fatal("failed2 %v",err)
	}

	fmt.Printf("Response of add :%s\n",addres.Id)

	taskres,err:=client2.GetCustomer(context.Background(),&tt.Empty{})
	if err!=nil{
		log.Fatal("failed2 %v",err)
	}
	fmt.Printf("Response of get :")
	for _,det:=range taskres.Customer{
		fmt.Printf("Name:%s,AccountID:%v,Balance:%v,BankID:%v\n",det.Name,det.AccountId,det.Balance,det.BankId)
	}
}