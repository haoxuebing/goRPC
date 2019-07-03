package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorpc/rpc_grpc_server_client/pb"
	"log"
)

func main() {

	cc,err:=grpc.Dial("127.0.0.1:5000",grpc.WithInsecure())
	if err!=nil{
		log.Fatalln("cc error:",err)
	}

	client := pb.NewArithServiceClient(cc)


	req := &pb.ArithRequest{A: 9, B: 2}

	ctx:=context.Background()
	res, err := client.Multiply(ctx,req)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.GetA(), req.GetB(), res.GetPro())

	res, err = client.Divide(ctx,req)
	if err != nil {
		log.Fatalln("arith error ", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
