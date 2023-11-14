package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"net"
	"os"
	"strconv"

	pb "DSMutualExclusion/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type node struct {
	id         int
	token      bool
	ownport    int
	holderport int
	pb.UnimplementedMutualServer
}

var (
	ID         = flag.Int("id", 0, "node id")
	hasToken   = flag.Bool("hastoken", false, "node id")
	ownPort    = flag.Int("ownport", 0, "client port number")
	holderPort = flag.Int("holderport", 0, "server port number")
)

func main() {
	flag.Parse()

	node := &node{
		id:         *ID,
		token:      *hasToken,
		ownport:    *ownPort,
		holderport: *holderPort,
	}

	//Requests token from another node (that might have token)
	go waitForTokenRequest(node)

	//Starts own node server
	go NodeServer(node)

	// Keep the server running until it is manually quit
	for {

	}
}

//______________________________________

func waitForTokenRequest(n *node) {
	connection, _ := connect()

	consoleReader := bufio.NewReader(os.Stdin)
	userInput, _ := consoleReader.ReadString('\n')
	if userInput == "request" {

		if n.token {
			log.Printf("Node with ID: %d has token, and is in CS", n.id)
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
			log.Printf("Node with ID: %d requests token", n.id)
			tokenResponse, err := connection.RequestToken(context.Background(), &pb.TokenRequest{
				ID: int64(n.id),
			})
			if err != nil {
				log.Printf(err.Error())
			} else {
				//Prints 1 when true
				log.Printf("Tokenholder %d says the access is %d\n", tokenResponse.HolderID, tokenResponse.Access)
			}
		}
	}
}

func (c *node) RequestToken(ctx context.Context, in *pb.TokenRequest) (*pb.TokenResponse, error) {
	log.Printf("Node with ID %d requests token", in.ID)
	return &pb.TokenResponse{
		HolderID: in.ID,
		Access:   true,
	}, nil
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

func (c *node) ReleaseToken(ctx context.Context, in *pb.ReleaseRequest) (*pb.ReleaseResponse, error) {
	log.Printf("Node with ID %d wants to release token", in.HolderID)
	return &pb.ReleaseResponse{
		Access: false,
	}, nil
}

func NodeServer(node *node) {
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
