package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"gorpc/rpc_grpc_server_client/pb"
	"log"
	"net"
	"os"
)

// 算术运算结构体
type Arith struct {
}

//context.Context, *ArithRequest
// 乘法运算方法
func (this *Arith) Multiply(ctx context.Context, req *pb.ArithRequest) (*pb.ArithResponse, error) {
	res := new(pb.ArithResponse)
	res.Pro = req.GetA() * req.GetB()
	return res, nil
}

// 除法运算方法
func (this *Arith) Divide(ctx context.Context, req *pb.ArithRequest) (res *pb.ArithResponse, err error) {
	if req.GetB() == 0 {
		err=errors.New("divide by zero")
		return
	}
	res =new(pb.ArithResponse)
	res.Quo = req.GetA() / req.GetB()
	res.Rem = req.GetA() % req.GetB()
	return
}

//(context.Context, *ArithRequest) (*ArithResponse, error)
func main() {
	lis,err := net.Listen("tcp", ":5000")
	if err!=nil{
		log.Fatalln("fatal error:",err)

	}
	fmt.Fprintf(os.Stdout,"%s","start connection\n")
	s := grpc.NewServer()
	pb.RegisterArithServiceServer(s, &Arith{})
	s.Serve(lis)

}
