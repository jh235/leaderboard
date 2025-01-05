package main

import (
	"fmt"
	"game_leaderboard/internal/leaderboard"
	pb "game_leaderboard/internal/leaderboard"
	"game_leaderboard/internal/leaderboard/leaderboard_impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func runGrpcServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	redisLeaderboard := leaderboard_impl.NewRedisLeaderboard("localhost:6379", "playerRank")
	service := leaderboard.NewLeaderboardService(redisLeaderboard)

	grpcServer := grpc.NewServer()
	pb.RegisterLeaderboardServiceServer(grpcServer, service)

	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {
	go runGrpcServer()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Printf("\n收到信号: %s\n", sig)
		os.Exit(0)
	}()

	select {}

}
