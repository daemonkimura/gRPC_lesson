
package main
import (
    "context"
    "fmt"
    "log"
	pb "client/pb/proto"
	"google.golang.org/grpc"
)
func main() {
    //sampleなのでwithInsecure
    conn, err := grpc.Dial("server:8000", grpc.WithInsecure())
    if err != nil {
        log.Fatal("client connection error:", err)
    }
    defer conn.Close()
    client := pb.NewCatClient(conn)
    message := &pb.GetMyCatMessage{TargetCat: "mike"}
    res, err := client.GetMyCat(context.TODO(), message)
    fmt.Printf("result:%#v \n", res)
    fmt.Printf("error::%#v \n", err)
}