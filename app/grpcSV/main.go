package main
import (
    "log"
    "net"
	//
	"google.golang.org/grpc"
	//
	pb "test/pb/proto"
    "test/service"

)
func main() {
    listenPort, err := net.Listen("tcp", ":8000")
    if err != nil {
        log.Fatalln(err)
    }
    server := grpc.NewServer()
    catService := &service.MyCatService{}
    // 実行したい実処理をseverに登録する
    pb.RegisterCatServer(server, catService)
    server.Serve(listenPort)
}