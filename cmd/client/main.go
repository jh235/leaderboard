package main

import (
	"context"
	"log"
	"time"

	"game_leaderboard/internal/leaderboard"
	"google.golang.org/grpc"
)

func main() {
	// 连接到 gRPC 服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := leaderboard.NewLeaderboardServiceClient(conn)

	// 测试 UpdateScore 方法
	_, err = client.UpdateScore(context.Background(), &leaderboard.UpdateScoreRequest{
		PlayerId:  "123",
		Score:     200,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		log.Fatalf("could not update score: %v", err)
	}
	log.Println("Score updated successfully!")

	client.UpdateScore(context.Background(), &leaderboard.UpdateScoreRequest{
		PlayerId:  "1234",
		Score:     100,
		Timestamp: time.Now().Unix(),
	})

	client.UpdateScore(context.Background(), &leaderboard.UpdateScoreRequest{
		PlayerId:  "12345",
		Score:     100,
		Timestamp: time.Now().Unix(),
	})

	client.UpdateScore(context.Background(), &leaderboard.UpdateScoreRequest{
		PlayerId:  "123456",
		Score:     90,
		Timestamp: time.Now().Unix(),
	})

	// 测试 GetPlayerRank 方法
	rankInfo, err := client.GetPlayerRank(context.Background(), &leaderboard.GetPlayerRankRequest{
		PlayerId: "123456",
	})
	if err != nil {
		log.Fatalf("could not get player rank: %v", err)
	}
	log.Printf("Player Rank: %+v", rankInfo)

	// 测试 GetTopN 方法
	topNResponse, err := client.GetTopN(context.Background(), &leaderboard.GetTopNRequest{
		N: 5,
	})
	if err != nil {
		log.Fatalf("could not get top N players: %v", err)
	}
	log.Printf("Top N players: %+v", topNResponse)

	// 测试 GetPlayerRankRange 方法
	rankRangeResponse, err := client.GetPlayerRankRange(context.Background(), &leaderboard.GetPlayerRankRangeRequest{
		PlayerId: "123456",
		Range:    10,
	})
	if err != nil {
		log.Fatalf("could not get player rank range: %v", err)
	}
	log.Printf("Player Rank Range: %+v", rankRangeResponse)
}
