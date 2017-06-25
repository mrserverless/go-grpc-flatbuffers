package main

import (
	"github.com/yunspace/go-grpc-flatbuffers/server/hello"
	"github.com/micro/go-micro/metadata"
	"context"
	"google.golang.org/grpc"
	"log"
	"fmt"
	"github.com/google/flatbuffers/go"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cl := hello.NewGreeterClient(conn)

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp, err := cl.Say(ctx, createHello())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}

func createHello() *flatbuffers.Builder {
	b := flatbuffers.NewBuilder(0)
	name := b.CreateString("")

	hello.RequestStart(b)
	hello.RequestAddName(b, name)
	end := hello.RequestEnd(b)

	b.Finish(end)
	return b
}
