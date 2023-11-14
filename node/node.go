package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"
	"time"

	pb "DSMutualExclusion/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Node struct {
	id         int
	token      bool
	ownport    int
	holderport int
	pb.UnimplementedMutualServer
}

var (
	ownid      = flag.Int("id", 0, "node id")
	hasToken   = flag.Bool("hastoken", false, "has token")
	ownPort    = flag.Int("ownport", 0, "client port number")
	holderPort = flag.Int("holderport", 0, "server port number")
)

func main() {
	flag.Parse()

	node := &Node{
		id:         *ownid,
		token:      *hasToken,
		ownport:    *ownPort,
		holderport: *holderPort,
	}

	//Starts own node server
	go NodeServer(node)

	time.Sleep(1000)
	//Requests token from another node (that might have token)
	go waitForTokenRequest(node)

	// Keep the server running until it is manually quit
	for {

	}
}

//______________________________________

func waitForTokenRequest(n *Node) {
	connection, _ := connect()

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//input := scanner.Text()
	//log.Printf("Node with ID: %d begins requesting with input: %s\n", n.id, input)
	for {
		if n.token {
			time.Sleep(15 * time.Second)
			log.Printf("Node with ID: %d has token, and is in the CRITICAL SECTION", n.id)
			releaseresponse, err := connection.ReleaseToken(context.Background(), &pb.ReleaseRequest{
				HolderID: int64(n.id),
			})
			if err != nil {
				log.Printf(err.Error())
			} else {
				n.token = releaseresponse.Access
				//Prints 1 when true
				log.Printf("Tokenholder with ID: %d releases the token", n.id)
			}
		} else {
			time.Sleep(15 * time.Second)
			log.Printf("Node with ID: %d requests token", n.id)
			tokenResponse, err := connection.RequestToken(context.Background(), &pb.TokenRequest{
				ID: int64(n.id),
			})
			if err != nil {
				log.Printf(err.Error())
			} else {
				//Prints 1 when true
				log.Printf("Tokenholder %d says the access is %d\n", tokenResponse.HolderID, tokenResponse.Access)
				if tokenResponse.Access {
					n.token = tokenResponse.Access
				} else {
					n.token = false
				}
			}
		}
	}
}

func (c *Node) RequestToken(ctx context.Context, in *pb.TokenRequest) (*pb.TokenResponse, error) {
	log.Printf("Node with ID %d requests token", in.ID)
	if !c.token {
		return &pb.TokenResponse{
			HolderID: in.ID,
			Access:   true,
		}, nil
	} else {
		return &pb.TokenResponse{
			HolderID: in.ID,
			Access:   false,
		}, nil
	}
}

func connect() (pb.MutualClient, error) {

	conn, err := grpc.Dial(":"+strconv.Itoa(*holderPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %d", *holderPort)
	} else {
		log.Printf("Connected to the server at port %d", *holderPort)
	}
	return pb.NewMutualClient(conn), nil

}

//____________________________________

func (c *Node) ReleaseToken(ctx context.Context, in *pb.ReleaseRequest) (*pb.ReleaseResponse, error) {
	log.Printf("Node with ID %d wants to release token", in.HolderID)
	return &pb.ReleaseResponse{
		Access: false,
	}, nil
}

func NodeServer(node *Node) {
	// Start the server
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(node.ownport))
	if err != nil {
		log.Fatalf("Could not create the server %v", err)
	}
	log.Printf("Started server at port: %d", node.ownport)

	pb.RegisterMutualServer(s, node)

	serveError := s.Serve(listener)
	if serveError != nil {
		log.Fatalf("Could not serve listener")
	}

}
