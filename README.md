![Hermione Granger](./github/image.jpg)

# Hermione
> Go microservice framework 

## Installation
```bash
go get -u github.com/im-kulikov/hermione
```

## Usage

### Server
```go
package main

import (
	"fmt"
	"os"

	"github.com/im-kulikov/hermione"
	proto "github.com/im-kulikov/hermione/examples/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/reflection"
)

type Greeter struct {
	service gemstone.Service
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest) (res *proto.HelloResponse, err error) {
	g.service.Logger().Infof("We have new greeting request from %s", req.Name)

	return &proto.HelloResponse{Greeting: "Hello, " + req.Name}, nil
}

func main() {
	service, err := gemstone.NewService(
		gemstone.Version("1.0.0"),
		gemstone.Name("greeter"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	greeter := &Greeter{
		service: service,
	}
	proto.RegisterGreeterServer(service.Server(), greeter)
	reflection.Register(service.Server())

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

```

### Client
```go
package main

import (
	"fmt"
	"os"

	"github.com/im-kulikov/hermione"
	proto "github.com/im-kulikov/hermione/examples/proto"
	"golang.org/x/net/context"
)

func main() {
	// Create a new service. Optionally include some options here.
	service, err := gemstone.NewService()
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
	rsp, err := greeter.Hello(context.Background(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}

```