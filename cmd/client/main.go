package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	kvstorev1 "github.com/rayzhao/grpc-kv-store/pkg/kvstore/v1"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: client <command> [args]")
		fmt.Println("  put <id> <value>")
		fmt.Println("  get <id>")
		fmt.Println("  delete <id>")
		fmt.Println("  list")
		os.Exit(1)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := kvstorev1.NewKVStoreClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch os.Args[1] {
	case "put":
		if len(os.Args) < 4 {
			log.Fatal("usage: client put <id> <value>")
		}
		resp, err := client.Put(ctx, &kvstorev1.PutRequest{Id: os.Args[2], Value: os.Args[3]})
		if err != nil {
			log.Fatalf("put failed: %v", err)
		}
		fmt.Printf("OK: %s\n", resp.Value)

	case "get":
		if len(os.Args) < 3 {
			log.Fatal("usage: client get <id>")
		}
		resp, err := client.Get(ctx, &kvstorev1.GetRequest{Id: os.Args[2]})
		if err != nil {
			log.Fatalf("get failed: %v", err)
		}
		fmt.Printf("%s\n", resp.Value)

	case "delete":
		if len(os.Args) < 3 {
			log.Fatal("usage: client delete <id>")
		}
		resp, err := client.Delete(ctx, &kvstorev1.DeleteRequest{Id: os.Args[2]})
		if err != nil {
			log.Fatalf("delete failed: %v", err)
		}
		fmt.Printf("deleted: %v\n", resp.Deleted)

	case "list":
		resp, err := client.List(ctx, &kvstorev1.ListRequest{})
		if err != nil {
			log.Fatalf("list failed: %v", err)
		}
		if len(resp.Items) == 0 {
			fmt.Println("(empty)")
			return
		}
		for _, item := range resp.Items {
			fmt.Printf("%s = %s\n", item.Id, item.Value)
		}

	case "watch":
		// TODO: 你来实现
		//   1. 检查参数：需要一个 id
		//   2. 调用 client.Watch(ctx, &kvstorev1.WatchRequest{...})
		//      注意：watch 是长连接，ctx 不能用 5 秒 timeout
		//      用 context.Background() 或者一个很长的 timeout
		//   3. 拿到 stream 后，for 循环调用 stream.Recv()
		//   4. 每次收到 resp，打印出来
		//   5. err == io.EOF 时退出
		log.Fatal("watch: not implemented yet")

	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}
}
