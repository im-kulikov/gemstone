package main

import (
	"fmt"
	"os"
	"time"

	"github.com/im-kulikov/hermione"
	proto "github.com/im-kulikov/hermione/examples/proto"
	"golang.org/x/net/context"
)

func main() {
	// Create a new service. Optionally include some options here.
	service, err := hermione.NewService()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create new greeter client
	client, err := service.Client("greeter")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			fmt.Printf("Closing server listener error: %v\n", closeErr)
		}
	}()

	greeter := proto.NewGreeterClient(client)

	// Call the greeter
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*2)
	rsp, err := greeter.Hello(ctx, &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
