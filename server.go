package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main(){
	
	//ここではTCP通信(https://e-words.jp/w/TCP.html)でポート9000番で立ち上げる
	
	lis, err :=net.Listen("tcp",":9000")
	
	//ここでerrがある場合はエラーハンドリングしてエラーを出す
	if err !=nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//grpcサーバーを作成する
	gerpcServer :=grpc.NewServer()

	//ここでエラーがある場合はエラーハンドリングとしてエラーを出す
	if err:=gerpcServer.Serve(lis); err!=nil{
		log.Fatalf("failed to listen: %v", err)
	}
}